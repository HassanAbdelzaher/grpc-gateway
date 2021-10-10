package grpc_proxy

import (
	"time"

	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"github.com/sirupsen/logrus"
	"github.com/wmdev4/shipswift-gateway/config"
)

func grpc_web_options(conf *config.GatewayConfig) []grpcweb.Option {
	allowedOrigins := makeAllowedOrigins(conf.AllowedOrigins)

	options := []grpcweb.Option{
		grpcweb.WithCorsForRegisteredEndpointsOnly(false),
		grpcweb.WithOriginFunc(makeHttpOriginFunc(allowedOrigins)),
	}

	if conf.UseWebsockets {
		logrus.Println("using websockets")
		options = append(
			options,
			grpcweb.WithWebsockets(true),
			grpcweb.WithWebsocketOriginFunc(makeWebsocketOriginFunc(allowedOrigins)),
		)
		if conf.WebsocketPingInterval.Duration >= time.Second {
			logrus.Infof("websocket keepalive pinging enabled, the timeout interval is %s", conf.WebsocketPingInterval.String())
		}
		options = append(
			options,
			grpcweb.WithWebsocketPingInterval(conf.WebsocketPingInterval.Duration),
		)
	}

	if len(conf.AllowedHeaders) > 0 {
		options = append(
			options,
			grpcweb.WithAllowedRequestHeaders(conf.AllowedHeaders),
		)
	}
	return options
}
