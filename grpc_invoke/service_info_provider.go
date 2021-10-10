package grpc_invoke

import "context"

type ServiceInfoProvider interface {
	ListMethods(ctx context.Context, address string) ([]*ServiceInfo, error)
}
