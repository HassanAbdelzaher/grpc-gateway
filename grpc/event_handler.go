package grpc

import (
	"fmt"
	"github.com/fullstorydev/grpcurl"
	"github.com/golang/protobuf/proto"
	"github.com/jhump/protoreflect/desc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"io"
)

// DefaultEventHandler logs events to a writer. This is not thread-safe, but is
// safe for use with InvokeRPC as long as NumResponses and Status are not read
// until the call to InvokeRPC completes.
type DefaultEventHandler struct {
	Out       io.Writer
	//Formatter Formatter
	// 0 = default
	// 1 = verbose
	// 2 = very verbose
	VerbosityLevel int

	// NumResponses is the number of responses that have been received.
	NumResponses int
	// Status is the status that was received at the end of an RPC. It is
	// nil if the RPC is still in progress.
	Status *status.Status
}


func NewDefaultEventHandler(out io.Writer, descSource grpcurl.DescriptorSource, verbose bool) *DefaultEventHandler {
	verbosityLevel := 0
	if verbose {
		verbosityLevel = 1
	}
	return &DefaultEventHandler{
		Out:            out,
		VerbosityLevel: verbosityLevel,
	}
}

func (h *DefaultEventHandler) OnResolveMethod(md *desc.MethodDescriptor) {
	if h.VerbosityLevel > 0 {
		txt, err := grpcurl.GetDescriptorText(md, nil)
		if err == nil {
			fmt.Fprintf(h.Out, "\nResolved method descriptor:\n%s\n", txt)
		}
	}
}

func (h *DefaultEventHandler) OnSendHeaders(md metadata.MD) {
	if h.VerbosityLevel > 0 {
		fmt.Fprintf(h.Out, "\nRequest metadata to send:\n%s\n", grpcurl.MetadataToString(md))
	}
}

func (h *DefaultEventHandler) OnReceiveHeaders(md metadata.MD) {
	if h.VerbosityLevel > 0 {
		fmt.Fprintf(h.Out, "\nResponse headers received:\n%s\n", grpcurl.MetadataToString(md))
	}
}

func (h *DefaultEventHandler) OnReceiveResponse(resp proto.Message) {
	h.NumResponses++
	if h.VerbosityLevel > 1 {
		fmt.Fprintf(h.Out, "\nEstimated response size: %d bytes\n", proto.Size(resp))
	}
	if h.VerbosityLevel > 0 {
		fmt.Fprint(h.Out, "\nResponse contents:\n")
	}
	fmt.Fprintln(h.Out, resp.String())
}

func (h *DefaultEventHandler) OnReceiveTrailers(stat *status.Status, md metadata.MD) {
	h.Status = stat
	if h.VerbosityLevel > 0 {
		fmt.Fprintf(h.Out, "\nResponse trailers received:\n%s\n", grpcurl.MetadataToString(md))
	}
}