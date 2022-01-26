package core

import (
	"log"
	"net/http"
	"strings"

	"github.com/sirupsen/logrus"
	grpc_proxy "github.com/wmdev4/shipswift-gateway/core/grpc_proxy"
	http_proxy "github.com/wmdev4/shipswift-gateway/core/http_proxy"
)

var staticContentTypes = []string{"text/html", "text/css", "text/javascript", "application/javascript", "text/plain"}

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

func setupResponse(w http.ResponseWriter) {
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(w).Header().Set("Access-Control-Allow-Headers", "*")
}

func (p *httpSplitter) ServeHTTP(wr http.ResponseWriter, req *http.Request) {
	setupResponse(wr)
	if strings.ToUpper(req.Method) == "OPTIONS" {
		wr.WriteHeader(http.StatusOK)
		return
	}
	cType := req.Header.Get("Content-Type")
	if cType == "" {
		cType = req.Header.Get("Content-type")
	}
	if cType == "" {
		cType = req.Header.Get("content-type")
	}
	for k, h := range req.Header {
		log.Println("item:,", "=", k, h)
	}
	accept := req.Header.Get("Accept")
	cType = strings.ToLower(cType)
	cType = strings.TrimSpace(cType)
	logrus.Println("content type :" + cType)
	if strings.Contains(cType, "application/grpc") || len(req.Header) == 0 || strings.Contains(accept, "application/grpc") {
		logrus.Println("grpc request")
		p.grpcHandler.ServeHTTP(wr, req)
		return
	}
	logrus.Println("http request")
	path := req.URL.Path
	path = strings.TrimSpace(strings.ToLower(path))
	p.httpHandler.ServeHTTP(wr, req)
}
