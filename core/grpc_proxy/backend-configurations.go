package grpc_proxy

import (
	"errors"
	"fmt"
	"sync"

	"github.com/wmdev4/shipswift-gateway/config"
	"google.golang.org/grpc"
)

var lock sync.Mutex

type backEndConfigurtaions map[string]*config.BackEndConfig

var BackEnds map[string]*grpc.ClientConn
var BackEndConfig *backEndConfigurtaions

func init() {
	BackEnds = make(map[string]*grpc.ClientConn)
	BackEndConfig = &backEndConfigurtaions{}
	for id := range config.Config.Services {
		srv := config.Config.Services[id]
		if srv == nil || srv.ServiceName == "" || srv.Backends == nil || len(srv.Backends) == 0 {
			continue
		}
		backend := srv.Backends[0]
		BackEndConfig.Add(srv.ServiceName, backend)
	}
}
func (o *backEndConfigurtaions) Add(name string, op *config.BackEndConfig) {
	if o == nil {
		var _op backEndConfigurtaions = make(map[string]*config.BackEndConfig)
		o = &_op
	}
	var mp map[string]*config.BackEndConfig = *o
	mp[name] = op
}
func (o *backEndConfigurtaions) Get(serviceName string) (*config.BackEndConfig, error) {
	if o == nil {
		return nil, errors.New("missing configuration")
	}
	var mp map[string]*config.BackEndConfig = *o
	op, ok := mp[serviceName]
	if !ok {
		return nil, errors.New(fmt.Sprintf("missing configuration for service:%s", serviceName))
	}
	if op == nil {
		return nil, errors.New(fmt.Sprintf("invalied configuration for service:%s", serviceName))
	}
	return op, nil
}

func DefualtBackEndOptions() *config.BackEndConfig {
	return &config.BackEndConfig{
		BackendHostPort:         "",
		BackendIsUsingTls:       false,
		BackendTlsNoVerify:      false,
		BackendTlsClientCert:    "",
		BackendTlsClientKey:     "",
		MaxCallRecvMsgSize:      1024 * 1024 * 4,
		BackendTlsCa:            []string{},
		BackendDefaultAuthority: "",
		BackendBackoffMaxDelay:  grpc.DefaultBackoffConfig.MaxDelay,
	}
}
