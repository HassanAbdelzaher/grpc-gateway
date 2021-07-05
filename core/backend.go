package core

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"github.com/mwitkow/grpc-proxy/proxy"
	"github.com/sirupsen/logrus"
	"github.com/wmdev4/shipswift-gateway/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
	"google.golang.org/grpc/credentials"
	"io/ioutil"
	"strings"
	"sync"
)

var lock sync.Mutex
type backEndConfigurtaions map[string]*config.BackEndConfig
var BackEnds map[string]*grpc.ClientConn
var  BackEndConfig *backEndConfigurtaions
func init(){
	BackEnds=make(map[string]*grpc.ClientConn)
	BackEndConfig=&backEndConfigurtaions{}
	for id:=range config.Config.Services{
		srv:=config.Config.Services[id]
		if srv==nil || srv.ServiceName=="" || srv.Backends==nil || len(srv.Backends)==0 {
			continue
		}
		backend:=srv.Backends[0]
		BackEndConfig.Add(srv.ServiceName,backend)
	}
}
func (o *backEndConfigurtaions) Add(name string,op *config.BackEndConfig){
	if o==nil{
		var _op backEndConfigurtaions=make(map[string]*config.BackEndConfig)
		o=&_op
	}
	var mp map[string]*config.BackEndConfig=*o
	mp[name]=op
}
func (o *backEndConfigurtaions) Get(serviceName string) (*config.BackEndConfig,error){
	if o==nil{
		return nil,errors.New("missing configuration")
	}
	var mp map[string]*config.BackEndConfig=*o
	op,ok:=mp[serviceName]
	if !ok{
		return nil,errors.New(fmt.Sprintf("missing configuration for service:%s",serviceName))
	}
	if op==nil{
		return nil,errors.New(fmt.Sprintf("invalied configuration for service:%s",serviceName))
	}
	return op,nil
}
func findService(fullMethodName string)  (*grpc.ClientConn,error){
	lock.Lock()
	defer lock.Unlock()
	fullMethodName=strings.TrimLeft(fullMethodName,"/")
	parts:=strings.Split(fullMethodName,"/")
	serviceName:=parts[0]
	backEnd,ok:=BackEnds[serviceName]
	if ok && backEnd!=nil{
		if backEnd.GetState()==connectivity.Ready{
			return backEnd,nil
		}
	}
	bConfig,err:=BackEndConfig.Get(serviceName)
	if err!=nil{
		return nil,err
	}
	bconn:=dialBackendOrFail(bConfig)
	BackEnds[serviceName]=bconn
	return bconn,nil
}
func DefualtBackEndOptions() *config.BackEndConfig{
	return &config.BackEndConfig{
		BackendHostPort:             "",
		BackendIsUsingTls:           false,
		BackendTlsNoVerify:          false,
		BackendTlsClientCert:        "",
		BackendTlsClientKey:     "",
		MaxCallRecvMsgSize:      1024*1024*4,
		BackendTlsCa:            []string{},
		BackendDefaultAuthority: "",
		BackendBackoffMaxDelay:  grpc.DefaultBackoffConfig.MaxDelay,
	}
}
func dialBackendOrFail(o *config.BackEndConfig) *grpc.ClientConn {
	if o==nil{
		o=DefualtBackEndOptions()
	}
	if o.BackendHostPort == "" {
		logrus.Fatalf("flag 'backend_addr' must be set")
	}
	opt := []grpc.DialOption{}
	opt = append(opt, grpc.WithCodec(proxy.Codec()))

	if o.BackendDefaultAuthority != "" {
		opt = append(opt, grpc.WithAuthority(o.BackendDefaultAuthority))
	}

	if o.BackendIsUsingTls {
		opt = append(opt, grpc.WithTransportCredentials(credentials.NewTLS(buildBackendTlsOrFail(o))))
	} else {
		opt = append(opt, grpc.WithInsecure())
	}

	opt = append(opt,
		grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(o.MaxCallRecvMsgSize)),
		grpc.WithBackoffMaxDelay(o.BackendBackoffMaxDelay),
	)

	cc, err := grpc.Dial(o.BackendHostPort, opt...)
	if err != nil {
		logrus.Fatalf("failed dialing backend: %v", err)
	}
	return cc
}
func buildBackendTlsOrFail(o *config.BackEndConfig) *tls.Config {
	tlsConfig := &tls.Config{}
	tlsConfig.MinVersion = tls.VersionTLS12
	if o.BackendTlsNoVerify {
		tlsConfig.InsecureSkipVerify = true
	} else {
		if len(o.BackendTlsCa) > 0 {
			tlsConfig.RootCAs = x509.NewCertPool()
			for _, path := range o.BackendTlsCa {
				data, err := ioutil.ReadFile(path)
				if err != nil {
					logrus.Fatalf("failed reading backend CA file %v: %v", path, err)
				}
				if ok := tlsConfig.RootCAs.AppendCertsFromPEM(data); !ok {
					logrus.Fatalf("failed processing backend CA file %v", path)
				}
			}
		}
	}
	if o.BackendTlsClientCert != "" || o.BackendTlsClientKey != "" {
		if o.BackendTlsClientCert == "" {
			logrus.Fatal("flag 'backend_client_tls_cert_file' must be set when 'backend_client_tls_key_file' is set")
		}
		if o.BackendTlsClientKey == "" {
			logrus.Fatal("flag 'backend_client_tls_key_file' must be set when 'backend_client_tls_cert_file' is set")
		}
		cert, err := tls.LoadX509KeyPair(o.BackendTlsClientCert, o.BackendTlsClientKey)
		if err != nil {
			logrus.Fatalf("failed reading TLS client keys: %v", err)
		}
		tlsConfig.Certificates = append(tlsConfig.Certificates, cert)
	}
	return tlsConfig
}
