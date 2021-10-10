package grpc_proxy

import (
	"net/http"
	_ "net/http/pprof" // register in DefaultServerMux

	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"github.com/wmdev4/shipswift-gateway/config"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/mwitkow/grpc-proxy/proxy"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	_ "golang.org/x/net/trace" // register in DefaultServerMux
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type grpcProxy struct {
	grpcHandler http.Handler
}

func NewGrpcProxy() *grpcProxy {
	var conf = config.Config
	logEntry := logrus.NewEntry(logrus.StandardLogger())
	grpcServer := buildGrpcProxyServer(logEntry)
	options := grpc_web_options(conf)
	wrappedGrpc := grpcweb.WrapServer(grpcServer, options...)
	h2s := &http2.Server{} // http2 clear text
	h2Handler := h2c.NewHandler(wrappedGrpc, h2s)
	return &grpcProxy{
		grpcHandler: h2Handler,
	}
}

func (p *grpcProxy) ServeHTTP(wr http.ResponseWriter, req *http.Request) {
	p.grpcHandler.ServeHTTP(wr, req)
}

func buildGrpcProxyServer(logger *logrus.Entry) *grpc.Server {
	var conf = config.Config
	// gRPC-wide changes.
	grpc.EnableTracing = true
	grpc_logrus.ReplaceGrpcLogger(logger)
	// gRPC proxy logic.

	director := func(ctx context.Context, fullMethodName string) (context.Context, *grpc.ClientConn, error) {
		backendConn, err := findService(fullMethodName)
		if err != nil {
			return ctx, nil, err
		}
		md, _ := metadata.FromIncomingContext(ctx)
		outCtx, _ := context.WithCancel(ctx)
		mdCopy := md.Copy()
		delete(mdCopy, "user-agent")
		// If this header is present in the request from the web client,
		// the actual connection to the backend will not be established.
		// https://github.com/improbable-eng/grpc-web/issues/568
		delete(mdCopy, "connection")
		outCtx = metadata.NewOutgoingContext(outCtx, mdCopy)
		return outCtx, backendConn, nil
	}
	// Server with logging and monitoring enabled.
	return grpc.NewServer(
		grpc.CustomCodec(proxy.Codec()), // needed for proxy to function.
		grpc.UnknownServiceHandler(proxy.TransparentHandler(director)),
		grpc.MaxRecvMsgSize(conf.MaxCallRecvMsgSize),
		grpc_middleware.WithUnaryServerChain(
			grpc_logrus.UnaryServerInterceptor(logger),
			grpc_prometheus.UnaryServerInterceptor,
		),
		grpc_middleware.WithStreamServerChain(
			grpc_logrus.StreamServerInterceptor(logger),
			grpc_prometheus.StreamServerInterceptor,
		),
	)
}
