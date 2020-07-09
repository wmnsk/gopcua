package uatest

import (
	"context"
	"crypto/rsa"
	"crypto/tls"
	"os"
	"testing"

	"github.com/gopcua/opcua/server"
	"github.com/gopcua/opcua/ua"
)

const (
	serverCertFile string = "test/server_cert.pem"
	serverKeyFile  string = "test/server_key.pem"

	authUser string = "test"
	authPass string = "test"
)

type GoServer struct {
	ext    bool
	extURL string

	// closeFunc is a teardown function called by Close()
	// Using a closure allows for more control of the inner workings of the function
	closeFunc func()

	s *server.Server
}

// Close gracefully shuts down the server
func (s *GoServer) Close() {
	if s.closeFunc != nil {
		s.closeFunc()
	}
}

func (s *GoServer) URL() string {
	if s.ext {
		return s.extURL
	}

	return s.s.URL()
}

func NewGoServer(t *testing.T) *GoServer {
	extServer := os.Getenv("OPC_EXTERNALSERVER")
	if extServer != "" {

		return &GoServer{
			ext:       true,
			extURL:    extServer,
			closeFunc: func() {},
		}
	}

	var opts []server.Option

	opts = append(opts,
		server.EnableSecurity("None", ua.MessageSecurityModeNone),
		server.EnableSecurity("Basic128Rsa15", ua.MessageSecurityModeSign),
		server.EnableSecurity("Basic128Rsa15", ua.MessageSecurityModeSignAndEncrypt),
		server.EnableSecurity("Basic256", ua.MessageSecurityModeSign),
		server.EnableSecurity("Basic256", ua.MessageSecurityModeSignAndEncrypt),
		server.EnableSecurity("Basic256Sha256", ua.MessageSecurityModeSignAndEncrypt),
		server.EnableSecurity("Basic256Sha256", ua.MessageSecurityModeSign),
		server.EnableSecurity("Aes128_Sha256_RsaOaep", ua.MessageSecurityModeSign),
		server.EnableSecurity("Aes128_Sha256_RsaOaep", ua.MessageSecurityModeSignAndEncrypt),
		server.EnableSecurity("Aes256_Sha256_RsaPss", ua.MessageSecurityModeSign),
		server.EnableSecurity("Aes256_Sha256_RsaPss", ua.MessageSecurityModeSignAndEncrypt),
	)

	opts = append(opts,
		server.EnableAuthMode(ua.UserTokenTypeAnonymous),
		server.EnableAuthMode(ua.UserTokenTypeUserName),
		server.EnableAuthMode(ua.UserTokenTypeCertificate),
	)

	c, err := tls.LoadX509KeyPair(serverCertFile, serverKeyFile)
	if err != nil {
		t.Errorf("Failed to load certificate: %s", err)
	} else {
		pk, ok := c.PrivateKey.(*rsa.PrivateKey)
		if !ok {
			t.Errorf("Invalid private key")
		}
		cert := c.Certificate[0]
		opts = append(opts, server.PrivateKey(pk), server.Certificate(cert))
	}

	s := server.New("", opts...)

	ctx, cancel := context.WithCancel(context.Background())
	if err := s.Start(ctx); err != nil {
		t.Fatalf("Error starting server, exiting: %s", err)
	}

	teardown := func() {
		cancel()
		err := s.Close()
		if err != nil {
			t.Errorf("Error shutting down server: %s", err)
		}
	}

	return &GoServer{s: s, closeFunc: teardown}
}
