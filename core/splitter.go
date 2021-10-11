package core

import (
	"net/http"
	"strings"

	"github.com/sirupsen/logrus"
	admin_panel "github.com/wmdev4/shipswift-gateway"
	grpc_proxy "github.com/wmdev4/shipswift-gateway/core/grpc_proxy"
	http_proxy "github.com/wmdev4/shipswift-gateway/core/http_proxy"
)

var staticContentTypes = []string{"text/html", "text/css", "text/javascript", "application/javascript", "text/plain"}

type httpSplitter struct {
	grpcHandler        http.Handler
	httpHandler        http.Handler
	adminPandelHandler http.Handler
}

func NewHttpSplitter() *httpSplitter {
	grpcH := grpc_proxy.NewGrpcProxy()
	httpH := http_proxy.NewHttpProxy()
	adminH := admin_panel.NewAdminPanelHandler()
	return &httpSplitter{
		grpcHandler:        grpcH,
		httpHandler:        httpH,
		adminPandelHandler: adminH,
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
	cType = strings.ToLower(cType)
	cType = strings.TrimSpace(cType)
	if strings.Contains(cType, "application/grpc") {
		p.grpcHandler.ServeHTTP(wr, req)
		return
	}
	path := req.URL.Path
	path = strings.TrimSpace(strings.ToLower(path))
	if isStaticContent(req) {
		logrus.Println("admin panel request")
		p.adminPandelHandler.ServeHTTP(wr, req)
		return
	}
	p.httpHandler.ServeHTTP(wr, req)
}

func isStaticContent(r *http.Request) bool {
	if r.URL.Path == "" || r.URL.Path == "/" || r.URL.Path == "/index.html" || strings.Contains(r.URL.Path, "/static/") {
		return true
	}
	cType := r.Header.Get("Content-Type")
	if cType == "" {
		cType = r.Header.Get("Content-type")
	}
	if cType == "" {
		cType = r.Header.Get("content-type")
	}
	cType = strings.TrimSpace(cType)
	logrus.Println("Content-Type:", cType)
	if cType == "" {
		return false
	}
	for _, c := range staticContentTypes {
		if strings.Index(c, cType) >= 0 {
			return true
		}
	}
	return false
}
