package grpc

import (
	"context"
	"fmt"
	"github.com/jhump/protoreflect/desc"
	curl "github.com/fullstorydev/grpcurl"
	"github.com/jhump/protoreflect/grpcreflect"
	rpb "google.golang.org/grpc/reflection/grpc_reflection_v1alpha"
)

type ReflectionServiceProvider struct {

}

func (p *ReflectionServiceProvider) ListMethods(ctx context.Context,addr string)([]*ServiceInfo,error){
	bcon, err := curl.BlockingDial(ctx, "tcp", addr, nil)
	if err != nil {
		return nil, err
	}
	defer bcon.Close()
	//var args interface{}
	stub := rpb.NewServerReflectionClient(bcon)
	refClient := grpcreflect.NewClient(ctx, stub)
	source := curl.DescriptorSourceFromServer(ctx, refClient)
	return ListMethodsFromSource(source)
}

func ListMethodsFromSource(source curl.DescriptorSource) ([]*ServiceInfo, error) {
	services, err := source.ListServices()
	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("failed to query for service descriptor %v", err)
	}
	ret := []*ServiceInfo{}
	for _, svc := range services {
		dsc, err := source.FindSymbol(svc)
		if err != nil {
			fmt.Println(err)
			return nil, fmt.Errorf("failed to query for service descriptor %v", err)
		}
		sd, ok := dsc.(*desc.ServiceDescriptor)
		if !ok {
			return nil, fmt.Errorf("target server does not expose service parsing error")
		}

		methods := sd.GetMethods()
		service := ServiceInfo{
			SeriveName: svc,
			Methods:    make([]MethodInfo, 0),
		}
		for _, m := range methods {
			method := MethodInfo{
				Name:               m.GetName(),
				InputType:          m.GetInputType().GetName(),
				OutputType:         m.GetOutputType().GetName(),
				FullyQualifiedName: m.GetFullyQualifiedName(),
				IsClientStreaming:  m.IsClientStreaming(),
				IsServerStreaming:  m.IsServerStreaming(),
				Desciptor:          m,
				Others:            make(map[string]interface{}),
			}
			method.InputFields = m.GetInputType().GetFields()
			method.OutputFields = m.GetOutputType().GetFields()
			method.Options = m.GetOptions().String()
			method.MethodOptions = m.GetMethodOptions().String()
			method.SourceInfo = m.GetSourceInfo().String()
			service.Methods = append(service.Methods, method)
		}
		ret = append(ret, &service)

	}

	return ret, nil
}
