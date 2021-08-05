package core

import (
"crypto/tls"

"github.com/mwitkow/go-conntrack/connhelpers"
logrus "github.com/sirupsen/logrus"
"crypto/x509"
"io/ioutil"
)

type TlsServerOptions struct{
	TlsServerCert string //Path to the PEM certificate for server use
	TlsServerKey string //Path to the PEM key for the certificate for the server use
	TlsServerClientCertVerification string //Controls whether a client certificate is on. Values: none, verify_if_given, require
	TlsServerClientCAFiles []string //Paths (comma separated) to PEM certificate chains used for client-side verification. If empty, host CA chain will be used.
}

func buildServerTlsOrFail(o *TlsServerOptions) *tls.Config {

	if o==nil || o.TlsServerCert == "" || o.TlsServerKey == "" {
		logrus.Println("flags server_tls_cert_file and server_tls_key_file must be set")
		return nil
	}
	tlsConfig, err := connhelpers.TlsConfigForServerCerts(o.TlsServerCert, o.TlsServerKey)
	if err != nil {
		logrus.Fatalf("failed reading TLS server keys: %v", err)
	}
	tlsConfig.MinVersion = tls.VersionTLS12
	switch o.TlsServerClientCertVerification {
	case "none":
		tlsConfig.ClientAuth = tls.NoClientCert
	case "verify_if_given":
		tlsConfig.ClientAuth = tls.VerifyClientCertIfGiven
	case "require":
		tlsConfig.ClientAuth = tls.RequireAndVerifyClientCert
	default:
		logrus.Fatalf("Uknown value '%v' for server_tls_client_cert_verification", o.TlsServerClientCertVerification)
	}
	if tlsConfig.ClientAuth != tls.NoClientCert {
		if len(o.TlsServerClientCAFiles) > 0 {
			tlsConfig.ClientCAs = x509.NewCertPool()
			for _, path := range o.TlsServerClientCAFiles {
				data, err := ioutil.ReadFile(path)
				if err != nil {
					logrus.Fatalf("failed reading client CA file %v: %v", path, err)
				}
				if ok := tlsConfig.ClientCAs.AppendCertsFromPEM(data); !ok {
					logrus.Fatalf("failed processing client CA file %v", path)
				}
			}
		} else {
			var err error
			tlsConfig.ClientCAs, err = x509.SystemCertPool()
			if err != nil {
				logrus.Fatalf("no client CA files specified, fallback to system CA chain failed: %v", err)
			}
		}

	}
	tlsConfig, err = connhelpers.TlsConfigWithHttp2Enabled(tlsConfig)
	if err != nil {
		logrus.Fatalf("can't configure h2 handling: %v", err)
	}
	return tlsConfig
}
