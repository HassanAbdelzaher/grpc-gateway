package grpc_proxy

import (
	"net/http"

	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"github.com/wmdev4/shipswift-gateway/config"
	"google.golang.org/grpc/grpclog"
)

func makeHttpOriginFunc(allowedOrigins *allowedOrigins) func(origin string) bool {
	var conf = config.Config
	if conf.AllowAllOrigins {
		return func(origin string) bool {
			return true
		}
	}
	return allowedOrigins.IsAllowed
}

func makeWebsocketOriginFunc(allowedOrigins *allowedOrigins) func(req *http.Request) bool {
	var conf = config.Config
	if conf.AllowAllOrigins {
		return func(req *http.Request) bool {
			return true
		}
	} else {
		return func(req *http.Request) bool {
			origin, err := grpcweb.WebsocketRequestOrigin(req)
			if err != nil {
				grpclog.Warning(err)
				return false
			}
			return allowedOrigins.IsAllowed(origin)
		}
	}
}

func makeAllowedOrigins(origins []string) *allowedOrigins {
	o := map[string]struct{}{}
	for _, allowedOrigin := range origins {
		o[allowedOrigin] = struct{}{}
	}
	return &allowedOrigins{
		origins: o,
	}
}

type allowedOrigins struct {
	origins map[string]struct{}
}

func (a *allowedOrigins) IsAllowed(origin string) bool {
	_, ok := a.origins[origin]
	return ok
}
