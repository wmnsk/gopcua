package opcua_test

import (
	"context"
	"crypto/rsa"
	"crypto/tls"
	"fmt"
	"testing"
	"time"

	"github.com/gopcua/opcua"
	"github.com/gopcua/opcua/ua"
	"github.com/gopcua/opcua/uatest"
)

const (
	clientCertFile string = "test/client_cert.pem"
	clientKeyFile  string = "test/client_key.pem"

	authUser string = "test"
	authPass string = "test"
)

func TestClientRead(t *testing.T) {
	t.Parallel()

	s := uatest.NewGoServer(t)
	defer s.Close()

	endpoints, err := opcua.GetEndpoints(s.URL())
	if err != nil {
		t.Fatal(err)
	}

	for _, e := range endpoints {
		for _, a := range e.UserIdentityTokens {
			t.Run(fmt.Sprintf("%s/%s/%s", e.SecurityPolicyURI, e.SecurityMode, a.TokenType), func(t *testing.T) {
				t.Log(t.Name())
				var opts []opcua.Option

				cert, err := tls.LoadX509KeyPair(clientCertFile, clientKeyFile)
				if err != nil {
					t.Errorf("Failed to load certificate: %s", err)
				} else {
					pk, ok := cert.PrivateKey.(*rsa.PrivateKey)
					if !ok {
						t.Errorf("Invalid private key")
					}
					opts = append(opts, opcua.PrivateKey(pk), opcua.Certificate(cert.Certificate[0]))
				}

				opts = append(opts, opcua.SecurityFromEndpoint(e, a.TokenType))

				switch a.TokenType {
				case ua.UserTokenTypeUserName:
					opts = append(opts, opcua.AuthUsername(authUser, authPass))
				case ua.UserTokenTypeCertificate:
					opts = append(opts, opcua.Certificate(cert.Certificate[0]))
				}

				c := opcua.NewClient(s.URL(), opts...)
				if err := c.Connect(context.Background()); err != nil {
					t.Fatal(err)
				}
				defer c.Close()

				v, err := c.Node(ua.NewNumericNodeID(0, 2258)).Value()
				if err != nil {
					t.Error(err)
				}

				t.Logf("Time: %v", v.Value())

				// Connections get dropped if there's no delay between tries for some external servers
				// Suspect this is an anti-DDOS technique
				time.Sleep(500 * time.Millisecond)
			})
		}
	}

}

func Test100Clients(t *testing.T) {
	t.Parallel()

	s := uatest.NewGoServer(t)
	defer s.Close()

	var opts []opcua.Option

	var clients []*opcua.Client
	for i := 0; i < 100; i++ {
		c := opcua.NewClient(s.URL(), opts...)
		if err := c.Connect(context.Background()); err != nil {
			t.Fatal(err)
		}
		defer c.Close()

		clients = append(clients, c)
	}

	if len(clients) != 100 {
		t.Errorf("didn't create all clients, only %d of 100 started", len(clients))
	}

	for _, c := range clients {
		_, err := c.Node(ua.NewNumericNodeID(0, 2258)).Value()
		if err != nil {
			t.Error(err)
		}
	}

}

func TestDelayedRead(t *testing.T) {
	t.Parallel()

	s := uatest.NewGoServer(t)
	defer s.Close()

	var opts []opcua.Option

	c := opcua.NewClient(s.URL(), opts...)
	if err := c.Connect(context.Background()); err != nil {
		t.Fatal(err)
	}
	defer c.Close()

	// Default timeout is 10s: delay for longer than that
	time.Sleep(20 * time.Second)

	v, err := c.Node(ua.NewNumericNodeID(0, 2258)).Value()
	if err != nil {
		t.Error(err)
	}

	t.Logf("Time: %v", v.Value())
}
