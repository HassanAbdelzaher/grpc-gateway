package grpc

type ServiceInfo struct {
	SeriveName string
	Methods    []MethodInfo
	Others    map[string]interface{}
}