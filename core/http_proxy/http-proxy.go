package http_proxy

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/wmdev4/shipswift-gateway/config"
	"github.com/wmdev4/shipswift-gateway/core/balancer"
)

// Hop-by-hop headers. These are removed when sent to the backend.
// http://www.w3.org/Protocols/rfc2616/rfc2616-sec13.html
var hopHeaders = []string{
	"Connection",
	"Keep-Alive",
	"Proxy-Authenticate",
	"Proxy-Authorization",
	"Te", // canonicalized version of "TE"
	"Trailers",
	"Transfer-Encoding",
	"Upgrade",
}

func copyHeader(dst, src http.Header) {
	for k, vv := range src {
		for _, v := range vv {
			dst.Add(k, v)
		}
	}
}

func delHopHeaders(header http.Header) {
	for _, h := range hopHeaders {
		header.Del(h)
	}
}

func appendHostToXForwardHeader(header http.Header, host string) {
	// If we aren't the first proxy retain prior
	// X-Forwarded-For information as a comma+space
	// separated list and fold multiple headers into one.
	if prior, ok := header["X-Forwarded-For"]; ok {
		host = strings.Join(prior, ", ") + ", " + host
	}
	header.Set("X-Forwarded-For", host)
}

type httpProxy struct {
	balancer *balancer.LoadBalancer
}

func NewHttpProxy() *httpProxy {
	balancer := balancer.NewLoadBalancer()
	routes := config.Config.HttpRoutes
	if routes != nil {
		for _, r := range routes {
			if r.Backends == nil {
				continue
			}
			for _, b := range r.Backends {
				if b.BackendHostPort == "" {
					continue
				}
				balancer.AddToBackend(r.Url, b.BackendHostPort)
			}
		}
	}
	return &httpProxy{balancer: balancer}
}
func (p *httpProxy) ServeHTTP(wr http.ResponseWriter, r *http.Request) {
	defer func() {
		if r := recover(); r != nil {
			wr.WriteHeader(500)
			fmt.Fprint(wr, r)
		}
	}()
	path := r.URL.Path
	addr, err := p.balancer.Get(path)
	if err != nil {
		http.Error(wr, "Gateway Error : "+err.Error(), http.StatusInternalServerError)
		return
	}
	url := addr + r.RequestURI
	logrus.Println("fowrward to:", url)
	client := &http.Client{}
	req, err := http.NewRequest(r.Method, url, r.Body)
	if err != nil {
		http.Error(wr, err.Error(), http.StatusInternalServerError)
		return
	}
	for name, value := range r.Header {
		val := strings.Join(value, ",")
		req.Header.Set(name, val)
	}
	delHopHeaders(r.Header)
	if clientIP, _, err := net.SplitHostPort(r.RemoteAddr); err == nil {
		appendHostToXForwardHeader(req.Header, clientIP)
	}
	resp, err := client.Do(req)
	if err != nil {
		http.Error(wr, "Server Error : "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()
	delHopHeaders(resp.Header)
	copyHeader(wr.Header(), resp.Header)
	wr.WriteHeader(resp.StatusCode)
	io.Copy(wr, resp.Body)
}
