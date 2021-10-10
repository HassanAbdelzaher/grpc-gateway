package grpc_invoke

import (
	"github.com/jhump/protoreflect/desc"
)

type MethodInfo struct {
	Name               string
	InputType          string
	OutputType         string
	FullyQualifiedName string
	IsServerStreaming  bool
	IsClientStreaming  bool
	Options            string
	MethodOptions      string
	SourceInfo         string
	Desciptor          *desc.MethodDescriptor
	InputFields        []*desc.FieldDescriptor
	OutputFields       []*desc.FieldDescriptor
	Others             map[string]interface{}
}
