package core

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	_ "net/http/pprof" // register in DefaultServerMux
	"os"
	"time"

	"github.com/wmdev4/shipswift-gateway/config"
	grpc2 "github.com/wmdev4/shipswift-gateway/grpc"

	"crypto/tls"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"github.com/mwitkow/go-conntrack"
	"github.com/mwitkow/grpc-proxy/proxy"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
	"golang.org/x/net/context"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	_ "golang.org/x/net/trace" // register in DefaultServerMux
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/metadata"
)

var infoProvider = grpc2.ReflectionServiceProvider{}

func HandleListMethods(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in HandleListMethods", r)
		}
	}()

	var ret = make(map[string]interface{})
	//fullName := fmt.Sprintf("%v.%v.%v", pckg, service, method)
	addr := r.URL.Query().Get("address")
	if addr == "" {
		addr = r.URL.Query().Get("url")
	}
	services, err := infoProvider.ListMethods(r.Context(), addr)
	if err != nil {
		ret["error"] = err.Error()
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(services)
}

func Start() {
	var conf = config.Config
	pflag.Parse()
	for _, flag := range pflag.Args() {
		if flag == "true" || flag == "false" {
			logrus.Fatal("Boolean flags should be set using --flag=false, --flag=true or --flag (which is short for --flag=true). You cannot use --flag true or --flag false.")
		}
		logrus.Fatal("Unknown argument: " + flag)
	}

	logrus.SetOutput(os.Stdout)
	logEntry := logrus.NewEntry(logrus.StandardLogger())

	if conf.AllowAllOrigins && len(conf.AllowedOrigins) != 0 {
		logrus.Fatal("Ambiguous --allow_all_origins and --allow_origins configuration. Either set --allow_all_origins=true OR specify one or more origins to whitelist with --allow_origins, not both.")
	}

	grpcServer := buildGrpcProxyServer(logEntry)
	errChan := make(chan error)

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

	wrappedGrpc := grpcweb.WrapServer(grpcServer, options...)

	if !conf.RunHttpServer && !conf.RunTlsServer {
		logrus.Fatalf("Both run_http_server and run_tls_server are set to false. At least one must be enabled for grpcweb proxy to function correctly.")
	}

	if conf.RunHttpServer {
		// Debug server.
		debugServer := buildServer(wrappedGrpc)
		http.Handle("/metrics", promhttp.Handler())
		http.HandleFunc("/info", HandleListMethods)
		debugListener := buildListenerOrFail("http", conf.HttpPort)
		serveServer(debugServer, debugListener, "http", errChan)
	}

	if conf.RunTlsServer {
		// Debug server.
		servingServer := buildServer(wrappedGrpc)
		servingListener := buildListenerOrFail("http", conf.HttpTlsPort)
		tlsConfig := buildServerTlsOrFail(nil)
		if tlsConfig == nil {
			tlsConfig = &tls.Config{
				Certificates: []tls.Certificate{},
				// GetCertificate: getCertificate,
			}
		}
		servingListener = tls.NewListener(servingListener, tlsConfig)
		serveServer(servingServer, servingListener, "http_tls", errChan)
	}

	<-errChan
}

func buildServer(wrappedGrpc *grpcweb.WrappedGrpcServer) *http.Server {
	//build with enabling http2 clear text
	var conf = config.Config
	h2s := &http2.Server{} // http2 clear text
	h2Hnadler := h2c.NewHandler(wrappedGrpc, h2s)
	return &http.Server{
		WriteTimeout: conf.HttpMaxWriteTimeout.Duration,
		ReadTimeout:  conf.HttpMaxReadTimeout.Duration,
		Handler:      h2Hnadler,
	}
}

func serveServer(server *http.Server, listener net.Listener, name string, errChan chan error) {
	go func() {
		logrus.Infof("listening for %s on: %v", name, listener.Addr().String())
		if err := server.Serve(listener); err != nil {
			errChan <- fmt.Errorf("%s server error: %v", name, err)
		}
	}()
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

func buildListenerOrFail(name string, port int) net.Listener {
	var conf = config.Config
	addr := fmt.Sprintf("%s:%d", conf.BindAddr, port)
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed listening for '%v' on %v: %v", name, port, err)
	}
	return conntrack.NewListener(listener,
		conntrack.TrackWithName(name),
		conntrack.TrackWithTcpKeepAlive(20*time.Second),
		conntrack.TrackWithTracing(),
	)
}

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
