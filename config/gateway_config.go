package config

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

var Config *GatewayConfig

func init() {
	filename := "config.json"
	cbytes, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	var _Config GatewayConfig
	err = json.Unmarshal(cbytes, &_Config)
	if err != nil {
		str := string(cbytes)
		logrus.Println(str)
		panic(err)
	}
	Config = &_Config
}

// fileExists checks if a file exists and is not a directory before we
// try using it to prevent further errors.
func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

type Duration struct {
	time.Duration
}

func (d Duration) getDuration() time.Duration {
	return d.Duration
}

func (d Duration) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.String())
}

func (d *Duration) UnmarshalJSON(b []byte) error {
	var v interface{}
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	switch value := v.(type) {
	case float64:
		d.Duration = time.Duration(value)
		return nil
	case string:
		var err error
		d.Duration, err = time.ParseDuration(value)
		if err != nil {
			return err
		}
		return nil
	default:
		return errors.New("invalid duration")
	}
}

type BackEndConfig struct {
	BackendHostPort         string        `json:"address"`                //A host:port (IP or hostname) of the gRPC server to forward it to.
	BackendIsUsingTls       bool          `json:"is_using_tls"`           //Whether the gRPC server of the backend is serving in plaintext (false) or over TLS (true).
	BackendTlsNoVerify      bool          `json:"tls_no_verify"`          //Whether to ignore TLS verification checks (cert validity, hostname). *DO NOT USE IN PRODUCTION*.
	BackendTlsClientCert    string        `json:"tls_client_cert"`        //Path to the PEM certificate used when the backend requires client certificates for TLS.
	BackendTlsClientKey     string        `json:"tls_client_key"`         //Path to the PEM key used when the backend requires client certificates for TLS.
	MaxCallRecvMsgSize      int           `json:"max_call_recv_msg_size"` //Maximum receive message size limit. If not specified, the default of 4MB will be used.
	BackendTlsCa            []string      `json:"tls_ca"`                 //"Paths (comma separated) to PEM certificate chains used for verification of backend certificates. If empty, host CA chain will be used."
	BackendDefaultAuthority string        `json:"default_authority"`      //Default value to use for the HTTP/2 :authority header commonly used for routing gRPC calls through a backend gateway.
	BackendBackoffMaxDelay  time.Duration `json:"backoff_max_delay"`      //Maximum delay when backing off after failed connection attempts to the backend.
}

type MicroService struct {
	ServiceName string           `json:"service_name"`
	Backends    []*BackEndConfig `json:"backends"`
}

type HttpRoute struct {
	Url      string           `json:"url"`
	Backends []*BackEndConfig `json:"backends"`
}

type GatewayConfig struct {
	BindAddr              string          `json:"host"`
	HttpPort              int             `json:"http_port"`
	HttpTlsPort           int             `json:"tls_port"`
	AllowAllOrigins       bool            `json:"allow_all_origin"`
	AllowedOrigins        []string        `json:"allowed_origins"`
	AllowedHeaders        []string        `json:"allowed_headers"`
	RunHttpServer         bool            `json:"run_http_server"`
	RunTlsServer          bool            `json:"run_tls_server"`
	UseWebsockets         bool            `json:"use_websockets"`
	MaxCallRecvMsgSize    int             `json:"max_call_recv_msg_size"`
	WebsocketPingInterval Duration        `json:"websocket_ping_interval"`
	HttpMaxWriteTimeout   Duration        `json:"http_max_write_timeout"`
	HttpMaxReadTimeout    Duration        `json:"http_max_read_timeout"`
	Services              []*MicroService `json:"grpc_services"`
	HttpRoutes            []*HttpRoute    `json:"http_routes"`
}
