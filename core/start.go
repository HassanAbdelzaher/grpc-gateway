package core

import (
	"fmt"
	"log"
	"net"
	"net/http"
	_ "net/http/pprof" // register in DefaultServerMux
	"os"
	"time"

	"github.com/wmdev4/shipswift-gateway/config"
	grpc_proxy "github.com/wmdev4/shipswift-gateway/core/grpc_proxy"

	"crypto/tls"

	"github.com/mwitkow/go-conntrack"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
	_ "golang.org/x/net/trace" // register in DefaultServerMux
)

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

	if conf.AllowAllOrigins && len(conf.AllowedOrigins) != 0 {
		logrus.Fatal("Ambiguous --allow_all_origins and --allow_origins configuration. Either set --allow_all_origins=true OR specify one or more origins to whitelist with --allow_origins, not both.")
	}

	errChan := make(chan error)

	if !conf.RunHttpServer && !conf.RunTlsServer {
		logrus.Fatalf("Both run_http_server and run_tls_server are set to false. At least one must be enabled for grpcweb proxy to function correctly.")
	}

	if conf.RunHttpServer {
		// Debug server.
		debugServer := buildServer()
		http.Handle("/metrics", promhttp.Handler())
		http.HandleFunc("/info", grpc_proxy.HandleListMethods)
		debugListener := buildListenerOrFail("http", conf.HttpPort)
		serveServer(debugServer, debugListener, "http", errChan)
	}

	if conf.RunTlsServer {
		// Debug server.
		servingServer := buildServer()
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

func buildServer() *http.Server {
	//build with enabling http2 clear text
	var conf = config.Config
	handler := NewHttpSplitter()
	return &http.Server{
		WriteTimeout: conf.HttpMaxWriteTimeout.Duration,
		ReadTimeout:  conf.HttpMaxReadTimeout.Duration,
		Handler:      handler,
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

func buildListenerOrFail(name string, port int) net.Listener {
	var conf = config.Config
	addr := fmt.Sprintf("%s:%d", conf.BindAddr, port)
	logrus.Println(addr)
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
