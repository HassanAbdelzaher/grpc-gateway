package grpc_proxy

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"

	"github.com/sirupsen/logrus"
	"github.com/wmdev4/shipswift-gateway/config"
)

func buildBackendTlsOrFail(o *config.BackEndConfig) *tls.Config {
	tlsConfig := &tls.Config{}
	tlsConfig.MinVersion = tls.VersionTLS12
	if o.BackendTlsNoVerify {
		tlsConfig.InsecureSkipVerify = true
	} else {
		if len(o.BackendTlsCa) > 0 {
			tlsConfig.RootCAs = x509.NewCertPool()
			for _, path := range o.BackendTlsCa {
				data, err := ioutil.ReadFile(path)
				if err != nil {
					logrus.Fatalf("failed reading backend CA file %v: %v", path, err)
				}
				if ok := tlsConfig.RootCAs.AppendCertsFromPEM(data); !ok {
					logrus.Fatalf("failed processing backend CA file %v", path)
				}
			}
		}
	}
	if o.BackendTlsClientCert != "" || o.BackendTlsClientKey != "" {
		if o.BackendTlsClientCert == "" {
			logrus.Fatal("flag 'backend_client_tls_cert_file' must be set when 'backend_client_tls_key_file' is set")
		}
		if o.BackendTlsClientKey == "" {
			logrus.Fatal("flag 'backend_client_tls_key_file' must be set when 'backend_client_tls_cert_file' is set")
		}
		cert, err := tls.LoadX509KeyPair(o.BackendTlsClientCert, o.BackendTlsClientKey)
		if err != nil {
			logrus.Fatalf("failed reading TLS client keys: %v", err)
		}
		tlsConfig.Certificates = append(tlsConfig.Certificates, cert)
	}
	return tlsConfig
}
