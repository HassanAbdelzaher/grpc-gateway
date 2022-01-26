package grpc_proxy

import (
	"strings"

	"github.com/mwitkow/grpc-proxy/proxy"
	"github.com/sirupsen/logrus"
	"github.com/wmdev4/shipswift-gateway/config"
	"github.com/wmdev4/shipswift-gateway/core/balancer"
	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
	"google.golang.org/grpc/credentials"
)

var grpcBalancer = &balancer.LoadBalancer{}

func findService(fullMethodName string) (*grpc.ClientConn, error) {
	lock.Lock()
	defer lock.Unlock()
	fullMethodName = strings.TrimLeft(fullMethodName, "/")
	parts := strings.Split(fullMethodName, "/")
	serviceName := parts[0]
	backEnd, ok := BackEnds[serviceName]
	if ok && backEnd != nil {
		if backEnd.GetState() == connectivity.Ready {
			return backEnd, nil
		}
	}
	bConfig, err := BackEndConfig.Get(serviceName)
	if err != nil {
		return nil, err
	}
	bconn := dialBackendOrFail(bConfig)
	BackEnds[serviceName] = bconn
	return bconn, nil
}
func dialBackendOrFail(o *config.BackEndConfig) *grpc.ClientConn {
	if o == nil {
		o = DefualtBackEndOptions()
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
	maxSiz := o.MaxCallRecvMsgSize
	if maxSiz == 0 {
		maxSiz = 16 * 1024 * 1024
	}
	opt = append(opt,
		grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(maxSiz)),
		grpc.WithBackoffMaxDelay(o.BackendBackoffMaxDelay),
	)

	cc, err := grpc.Dial(o.BackendHostPort, opt...)
	if err != nil {
		logrus.Fatalf("failed dialing backend: %v", err)
	}
	return cc
}
