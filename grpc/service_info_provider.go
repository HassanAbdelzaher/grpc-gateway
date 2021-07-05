package grpc

import "context"

type ServiceInfoProvider interface {

	ListMethods(ctx context.Context,address string)([]*ServiceInfo,error)
}
