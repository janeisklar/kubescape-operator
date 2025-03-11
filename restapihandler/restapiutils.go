package restapihandler

import (
	"crypto/tls"
	"flag"
	"fmt"
)

func (handler *HTTPHandler) loadTLSKey() error {
	certFile := ""
	keyFile := ""

	flag.StringVar(&certFile, "tlsCertFile", "", "File containing the x509 Certificate for HTTPS.")
	flag.StringVar(&keyFile, "tlsKeyFile", "", "File containing the x509 private key to --tlsCertFile.")
	flag.Parse()

	if keyFile == "" || certFile == "" {
		return nil
	}

	pair, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		return fmt.Errorf("failed to load key pair: %w", err)
	}
	handler.keyPair = &pair
	return nil
}
