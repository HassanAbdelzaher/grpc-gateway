package core

import (
	"net/http"
	"strings"

	grpc_proxy "github.com/wmdev4/shipswift-gateway/core/grpc_proxy"
	http_proxy "github.com/wmdev4/shipswift-gateway/core/http_proxy"
)

type httpSplitter struct {
	grpcHandler http.Handler
	httpHandler http.Handler
}

func NewHttpSplitter() *httpSplitter {
	grpcH := grpc_proxy.NewGrpcProxy()
	httpH := http_proxy.NewHttpProxy()
	return &httpSplitter{
		grpcHandler: grpcH,
		httpHandler: httpH,
	}
}

func (p *httpSplitter) ServeHTTP(wr http.ResponseWriter, req *http.Request) {
	cType := req.Header.Get("Content-Type")
	if cType == "" {
		cType = req.Header.Get("Content-type")
	}
	if cType == "" {
		cType = req.Header.Get("content-type")
	}
	if strings.Contains(cType, "application/grpc") {
		p.grpcHandler.ServeHTTP(wr, req)
		return
	}
	p.httpHandler.ServeHTTP(wr, req)
}
