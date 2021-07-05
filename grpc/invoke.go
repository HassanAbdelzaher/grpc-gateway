package grpc

import (
	"bytes"
	"context"
	"fmt"
	curl "github.com/fullstorydev/grpcurl"
	"github.com/jhump/protoreflect/grpcreflect"
	"google.golang.org/grpc/codes"
	rpb "google.golang.org/grpc/reflection/grpc_reflection_v1alpha"
	"io"
)

//GrpcError grpc error struct
type GrpcError struct {
	Code    codes.Code
	Message string
	Details string
}

func (er *GrpcError) Error() string {
	if er == nil {
		return ""
	}
	return er.Message + "  " + er.Details
}

//NewError create new error instance
func NewError(message string) *GrpcError {
	err := &GrpcError{Message: message, Code: codes.Internal}
	return err
}

//NewErrorWithCode NewErrorWithCode
func NewErrorWithCode(message string, code codes.Code) *GrpcError {
	err := &GrpcError{Message: message, Code: code}
	return err
}

//NewErrorWithCode NewErrorWithCode
func NewErrorDetails(message string, code codes.Code, details string) *GrpcError {
	err := &GrpcError{Message: message, Code: code, Details: details}
	return err
}

//FromError create grpc error struct from error
func FromError(_err error) *GrpcError {
	var message = ""
	if _err != nil {
		message = _err.Error()
	}
	err := &GrpcError{Message: message, Code: codes.Internal}
	return err
}


// InvokeRPCRequest InvokeRPCRequest
func InvokegRPCRequest(ctx context.Context,addr string,fullMethodName string, message []byte, headers []string, writer io.Writer) (grpcError *GrpcError) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in InvokeRPCRequest", r)
			var message = ""
			str, ok := r.(string)
			if ok {
				message = str
			}
			err, ok := r.(error)
			if ok && err != nil {
				message = err.Error()
			}
			grpcError = &GrpcError{Code: codes.Internal, Message: "InvokeRPCRequest panic " + message}
		}
	}()
	bcon, err := curl.BlockingDial(ctx, "tcp", addr, nil)
	if err != nil {
		return FromError(err)
	}
	defer bcon.Close()
	stub := rpb.NewServerReflectionClient(bcon)
	refClient := grpcreflect.NewClient(ctx, stub)
	source := curl.DescriptorSourceFromServer(ctx, refClient)
	in := bytes.NewReader(message)
	//reqParser, _, err := curl.RequestParserAndFormatterFor(grpcurl.Format("json"), descriptorSrc, true, false, in)
	//emitDefaults:=true
	//resolver := curl.AnyResolverFromDescriptorSource(source)
	//res2 := curl.anyResolverWithFallback{AnyResolver: resolver}
	/*marshaler := jsonpb.Marshaler{
		EmitDefaults: true,
		Indent:       "  ",
		AnyResolver:  resolver,
		OrigName:     true, //this make json fields match proto fileds
	}
	formater := marshaler.MarshalToString //convert proto message to json string*/
	parser := NewProtobuffRequestParser(in)
	h := NewDefaultEventHandler(writer, source, false)

	err = curl.InvokeRPC(ctx, source, bcon, fullMethodName, headers, h,parser.Next)
	if err != nil {
		return FromError(err)
	}
	var status = h.Status
	if status.Code() != codes.OK {
		var code = ""
		if status.Code() == codes.Canceled {
			code = "Canceled:"
		}
		if status.Code() == codes.DeadlineExceeded {
			code = "DeadlineExceeded :"
		}
		if status.Code() == codes.PermissionDenied {
			code = "PermissionDenied :"
		}
		if status.Code() == codes.Unauthenticated {
			code = "Unauthenticated :"
		}
		if status.Code() == codes.Unimplemented {
			code = "Unimplemented :"
		}
		if status.Code() == codes.OutOfRange {
			code = "OutOfRange :"
		}
		if status.Code() == codes.Aborted {
			code = "Aborted :"
		}
		var details = status.Message()
		if err != nil {
			details = details + " " + err.Error()
		}

		return NewErrorDetails(code, status.Code(), details)
	}
	//str := writer.Buf.String()
	return nil
}

