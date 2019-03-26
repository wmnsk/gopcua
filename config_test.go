package opcua

import (
	"crypto/rsa"
	"crypto/tls"
	"encoding/pem"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/gopcua/opcua/ua"
	"github.com/gopcua/opcua/uapolicy"
	"github.com/gopcua/opcua/uasc"

	"github.com/pascaldekloe/goe/verify"
)

// test certificate generated with
// go run ~/sdk/gotip/src/crypto/tls/generate_cert.go -rsa-bits 1024 -host localhost
// expires Jun  5 20:10:13 2020 GMT
var (
	certPEM = []byte(`-----BEGIN CERTIFICATE-----
MIIB9TCCAV6gAwIBAgIRAJkygYaTfLZ9tOwtJvxdP7EwDQYJKoZIhvcNAQELBQAw
EjEQMA4GA1UEChMHQWNtZSBDbzAeFw0xOTA2MDYyMDEwMTNaFw0yMDA2MDUyMDEw
MTNaMBIxEDAOBgNVBAoTB0FjbWUgQ28wgZ8wDQYJKoZIhvcNAQEBBQADgY0AMIGJ
AoGBAJ8cfel4q4jcXrmGxiXmPrJzg0aZxKdWlPE23fr9KwpYkrQZ5ykzs6sGuuXE
OYtqINQNBP/5VXCinnDOZppI4QHlbUrWfKoGgJU2wQZuAQ7+Pz4l96EM5DnIBArb
liSp5s2LZiVLgw6v9tS6yu/Ci5QyfuyMz4JLg25Vt1KHFCD9AgMBAAGjSzBJMA4G
A1UdDwEB/wQEAwIFoDATBgNVHSUEDDAKBggrBgEFBQcDATAMBgNVHRMBAf8EAjAA
MBQGA1UdEQQNMAuCCWxvY2FsaG9zdDANBgkqhkiG9w0BAQsFAAOBgQAPuVO5vkF9
LfQ1JCrXC52CdKm8Gs+bYaDLQa6re4HaNPHAuEJaeMAHJ/4PHSsg6ghZ1MmBj1pc
GY1Q+sfu64IRjFdhnbL97a6GL+MgEVIvT9cl/DDcXtNZIl28Xk4KwAp3/lB1XrgK
cdqKnNkOBU19ulD8SOKzAPch5ydHPFfXCw==
-----END CERTIFICATE-----
`)

	keyPEM = []byte(`-----BEGIN RSA PRIVATE KEY-----
MIICXQIBAAKBgQCfHH3peKuI3F65hsYl5j6yc4NGmcSnVpTxNt36/SsKWJK0Gecp
M7OrBrrlxDmLaiDUDQT/+VVwop5wzmaaSOEB5W1K1nyqBoCVNsEGbgEO/j8+Jfeh
DOQ5yAQK25YkqebNi2YlS4MOr/bUusrvwouUMn7sjM+CS4NuVbdShxQg/QIDAQAB
AoGAfDdmJUtLv4ErgnOxZg0mjXKY3hlV6b4ycU6AZl4Xp/AWN/yw9v2iGrcaWh6j
PjAQiPvRF3W0Okb6ot7AQ1r6LaGWoRRlAhapJ0ZYw/TrlXULgXuvDjmlU6yfiSQQ
VBBdqJkqDr5El6a85MZlprxYz1OPudZyUGIPcr/wwj+a7pECQQDFVDS5G54T+x/m
RHUkGV+M2osHHBBaRKmiKlblg2U0Ep2P2eYmldcnt5xQxxrQedxkDJAuwR7hTQH0
o/Nnqh9fAkEAzmtUtCJ4Yp+6IBKFffGj86PfGHNSh9DARZ3fQPBDVFtRT9qVCwMA
rU9vJ/n1jvtYCYhzCFCuSSFTEIBglPiJIwJAQSdvjRsgU8qcGsTJxNSX5wMV2pAa
miOHuyKttHRxCwOGgMPaqSzacKPAei9znBhQe7xmMvnS/2MU3TjxGm5ikwJBALXl
YNvnsDwAUsymZZoJEJfHJPXv0Z869eOi7bPUxRAV9D4w+LueZr9SSzpoCtp3ZCnq
YqvGJP7ubbsR1YoQxQ8CQQCyCrltDYji5+KdxMOsDt0v7bCQWkQ3+pik09faK51Y
4upIBnmHPbJ80DfFIj/93JXna5JQpnIZGn/hitRixBWU
-----END RSA PRIVATE KEY-----
`)

	certDER = derBytes(certPEM)
	keyDER  = derBytes(keyPEM)
	cert    = x509Cert(certPEM, keyPEM)

	endpoints = []*ua.EndpointDescription{
		// anonymous auth
		&ua.EndpointDescription{
			SecurityPolicyURI: "a", // random value for testing
			SecurityMode:      5,   // random value for testing
			ServerCertificate: certDER,
			UserIdentityTokens: []*ua.UserTokenPolicy{
				&ua.UserTokenPolicy{
					TokenType:         ua.UserTokenTypeAnonymous,
					SecurityPolicyURI: "b", // random value for testing
				},
			},
		},
		// username auth
		&ua.EndpointDescription{
			SecurityPolicyURI: "a", // random value for testing
			SecurityMode:      5,   // random value for testing
			ServerCertificate: certDER,
			UserIdentityTokens: []*ua.UserTokenPolicy{
				&ua.UserTokenPolicy{
					TokenType:         ua.UserTokenTypeUserName,
					SecurityPolicyURI: "b", // random value for testing
				},
			},
		},
		// x509 cert auth
		&ua.EndpointDescription{
			SecurityPolicyURI: "a", // random value for testing
			SecurityMode:      5,   // random value for testing
			ServerCertificate: certDER,
			UserIdentityTokens: []*ua.UserTokenPolicy{
				&ua.UserTokenPolicy{
					TokenType:         ua.UserTokenTypeCertificate,
					SecurityPolicyURI: "b", // random value for testing
				},
			},
		},
		// issued token auth
		&ua.EndpointDescription{
			SecurityPolicyURI: "a", // random value for testing
			SecurityMode:      5,   // random value for testing
			ServerCertificate: certDER,
			UserIdentityTokens: []*ua.UserTokenPolicy{
				&ua.UserTokenPolicy{
					TokenType:         ua.UserTokenTypeIssuedToken,
					SecurityPolicyURI: "b", // random value for testing
				},
			},
		},
	}
)

func derBytes(b []byte) []byte {
	block, _ := pem.Decode(b)
	return block.Bytes
}

func x509Cert(c, k []byte) tls.Certificate {
	cert, err := tls.X509KeyPair(c, k)
	if err != nil {
		panic(err)
	}
	return cert
}

func TestOptions(t *testing.T) {
	d, err := ioutil.TempDir("", "gopcua")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(d)

	var (
		certDERFile = filepath.Join(d, "cert.der")
		certPEMFile = filepath.Join(d, "cert.pem")
		keyDERFile  = filepath.Join(d, "key.der")
		keyPEMFile  = filepath.Join(d, "key.pem")
	)

	if err := ioutil.WriteFile(certDERFile, certDER, 0644); err != nil {
		t.Fatal(err)
	}
	if err := ioutil.WriteFile(certPEMFile, certPEM, 0644); err != nil {
		t.Fatal(err)
	}
	if err := ioutil.WriteFile(keyDERFile, keyDER, 0644); err != nil {
		t.Fatal(err)
	}
	if err := ioutil.WriteFile(keyPEMFile, keyPEM, 0644); err != nil {
		t.Fatal(err)
	}
	defer os.Remove(keyPEMFile)

	tests := []struct {
		name string
		opt  Option
		c    *uasc.Config
		sc   *uasc.SessionConfig
	}{
		{
			name: `ApplicationName("a")`,
			opt:  ApplicationName("a"),
			sc: func() *uasc.SessionConfig {
				sc := DefaultSessionConfig()
				sc.ClientDescription.ApplicationName = &ua.LocalizedText{Text: "a"}
				return sc
			}(),
		},
		{
			name: `ApplicationURI("a")`,
			opt:  ApplicationURI("a"),
			sc: func() *uasc.SessionConfig {
				sc := DefaultSessionConfig()
				sc.ClientDescription.ApplicationURI = "a"
				return sc
			}(),
		},
		{
			name: `AuthAnonymous()`,
			opt:  AuthAnonymous(),
			sc: func() *uasc.SessionConfig {
				sc := DefaultSessionConfig()
				sc.UserIdentityToken = &ua.AnonymousIdentityToken{}
				return sc
			}(),
		},
		{
			name: `AuthCertificate()`,
			opt:  AuthCertificate(certDER),
			sc: func() *uasc.SessionConfig {
				sc := DefaultSessionConfig()
				sc.UserIdentityToken = &ua.X509IdentityToken{
					CertificateData: certDER,
				}
				return sc
			}(),
		},
		{
			name: `AuthIssuedToken()`,
			opt:  AuthIssuedToken([]byte("a")),
			sc: func() *uasc.SessionConfig {
				sc := DefaultSessionConfig()
				sc.UserIdentityToken = &ua.IssuedIdentityToken{
					TokenData: []byte("a"),
				}
				return sc
			}(),
		},
		{
			name: `AuthUsername()`,
			opt:  AuthUsername("user", "pass"),
			sc: func() *uasc.SessionConfig {
				sc := DefaultSessionConfig()
				sc.UserIdentityToken = &ua.UserNameIdentityToken{
					UserName: "user",
				}
				sc.AuthPassword = "pass"
				return sc
			}(),
		},
		{
			name: `Certificate`,
			opt:  Certificate(certDER),
			c: func() *uasc.Config {
				c := DefaultClientConfig()
				c.Certificate = certDER
				// todo(fs): test with cert that has a URI
				// sc.ClientDescription.ApplicationURI = ...
				return c
			}(),
		},
		{
			name: `CertificateFile("cert.der")`,
			opt:  CertificateFile(certDERFile),
			c: func() *uasc.Config {
				c := DefaultClientConfig()
				c.Certificate = certDER
				// todo(fs): test with cert that has a URI
				// sc.ClientDescription.ApplicationURI = ...
				return c
			}(),
		},
		{
			name: `CertificateFile("cert.pem")`,
			opt:  CertificateFile(certPEMFile),
			c: func() *uasc.Config {
				c := DefaultClientConfig()
				c.Certificate = certDER
				// todo(fs): test with cert that has a URI
				// sc.ClientDescription.ApplicationURI = ...
				return c
			}(),
		},
		{
			name: `Lifetime(10ms)`,
			opt:  Lifetime(10 * time.Millisecond),
			c: func() *uasc.Config {
				c := DefaultClientConfig()
				c.Lifetime = 10
				return c
			}(),
		},
		{
			name: `Locales("en-us", "de-de")`,
			opt:  Locales("en-us", "de-de"),
			sc: func() *uasc.SessionConfig {
				sc := DefaultSessionConfig()
				sc.LocaleIDs = []string{"en-us", "de-de"}
				return sc
			}(),
		},
		{
			name: `PrivateKey()`,
			opt:  PrivateKey(cert.PrivateKey.(*rsa.PrivateKey)),
			c: func() *uasc.Config {
				c := DefaultClientConfig()
				c.LocalKey = cert.PrivateKey.(*rsa.PrivateKey)
				return c
			}(),
		},
		{
			name: `PrivateKeyFile("key.der")`,
			opt:  PrivateKeyFile(keyDERFile),
			c: func() *uasc.Config {
				c := DefaultClientConfig()
				c.LocalKey = cert.PrivateKey.(*rsa.PrivateKey)
				return c
			}(),
		},
		{
			name: `PrivateKeyFile("key.pem")`,
			opt:  PrivateKeyFile(keyPEMFile),
			c: func() *uasc.Config {
				c := DefaultClientConfig()
				c.LocalKey = cert.PrivateKey.(*rsa.PrivateKey)
				return c
			}(),
		},
		{
			name: `ProductURI("a")`,
			opt:  ProductURI("a"),
			sc: func() *uasc.SessionConfig {
				sc := DefaultSessionConfig()
				sc.ClientDescription.ProductURI = "a"
				return sc
			}(),
		},
		{
			name: `RemoteCertificate`,
			opt:  RemoteCertificate(certDER),
			c: func() *uasc.Config {
				c := DefaultClientConfig()
				c.RemoteCertificate = certDER
				return c
			}(),
		},
		{
			name: `RemoteCertificateFile("cert.der")`,
			opt:  RemoteCertificateFile(certDERFile),
			c: func() *uasc.Config {
				c := DefaultClientConfig()
				c.RemoteCertificate = certDER
				return c
			}(),
		},
		{
			name: `RemoteCertificateFile("cert.pem")`,
			opt:  RemoteCertificateFile(certPEMFile),
			c: func() *uasc.Config {
				c := DefaultClientConfig()
				c.RemoteCertificate = certDER
				return c
			}(),
		},
		{
			name: `RequestTimeout(5s)`,
			opt:  RequestTimeout(5 * time.Second),
			c: func() *uasc.Config {
				c := DefaultClientConfig()
				c.RequestTimeout = 5 * time.Second
				return c
			}(),
		},
		{
			name: `SecurityFromEndpoint(no-match)`,
			opt: SecurityFromEndpoint(&ua.EndpointDescription{
				SecurityPolicyURI: "a",
				SecurityMode:      5,
				ServerCertificate: certDER,
			}, 0),
			c: func() *uasc.Config {
				c := DefaultClientConfig()
				c.SecurityPolicyURI = "a"
				c.SecurityMode = 5
				c.RemoteCertificate = certDER
				c.Thumbprint = uapolicy.Thumbprint(certDER)
				return c
			}(),
			sc: func() *uasc.SessionConfig {
				sc := DefaultSessionConfig()
				sc.UserIdentityToken = &ua.AnonymousIdentityToken{
					PolicyID: defaultAnonymousPolicyID,
				}
				sc.AuthPolicyURI = ua.SecurityPolicyURINone
				return sc
			}(),
		},
		{
			name: `SecurityFromEndpoint(anonymous)`,
			opt:  SecurityFromEndpoint(endpoints[0], ua.UserTokenTypeAnonymous),
			c: func() *uasc.Config {
				c := DefaultClientConfig()
				c.SecurityPolicyURI = "a"
				c.SecurityMode = 5
				c.RemoteCertificate = certDER
				c.Thumbprint = uapolicy.Thumbprint(certDER)
				return c
			}(),
			sc: func() *uasc.SessionConfig {
				sc := DefaultSessionConfig()
				sc.UserIdentityToken = &ua.AnonymousIdentityToken{}
				sc.AuthPolicyURI = "b"
				return sc
			}(),
		},
		{
			name: `SecurityFromEndpoint(username)`,
			opt:  SecurityFromEndpoint(endpoints[1], ua.UserTokenTypeUserName),
			c: func() *uasc.Config {
				c := DefaultClientConfig()
				c.SecurityPolicyURI = "a"
				c.SecurityMode = 5
				c.RemoteCertificate = certDER
				c.Thumbprint = uapolicy.Thumbprint(certDER)
				return c
			}(),
			sc: func() *uasc.SessionConfig {
				sc := DefaultSessionConfig()
				sc.UserIdentityToken = &ua.UserNameIdentityToken{}
				sc.AuthPolicyURI = "b"
				return sc
			}(),
		},
		{
			name: `SecurityFromEndpoint(certificate)`,
			opt:  SecurityFromEndpoint(endpoints[2], ua.UserTokenTypeCertificate),
			c: func() *uasc.Config {
				c := DefaultClientConfig()
				c.SecurityPolicyURI = "a"
				c.SecurityMode = 5
				c.RemoteCertificate = certDER
				c.Thumbprint = uapolicy.Thumbprint(certDER)
				return c
			}(),
			sc: func() *uasc.SessionConfig {
				sc := DefaultSessionConfig()
				sc.UserIdentityToken = &ua.X509IdentityToken{}
				sc.AuthPolicyURI = "b"
				return sc
			}(),
		},
		{
			name: `SecurityFromEndpoint(token)`,
			opt:  SecurityFromEndpoint(endpoints[3], ua.UserTokenTypeIssuedToken),
			c: func() *uasc.Config {
				c := DefaultClientConfig()
				c.SecurityPolicyURI = "a"
				c.SecurityMode = 5
				c.RemoteCertificate = certDER
				c.Thumbprint = uapolicy.Thumbprint(certDER)
				return c
			}(),
			sc: func() *uasc.SessionConfig {
				sc := DefaultSessionConfig()
				sc.UserIdentityToken = &ua.IssuedIdentityToken{}
				sc.AuthPolicyURI = "b"
				return sc
			}(),
		},
		{
			name: `SecurityPolicy("None")`,
			opt:  SecurityPolicy("None"),
			c: func() *uasc.Config {
				c := DefaultClientConfig()
				c.SecurityPolicyURI = ua.SecurityPolicyURINone
				return c
			}(),
		},
		{
			name: `SecurityMode(Sign)`,
			opt:  SecurityMode(ua.MessageSecurityModeSign),
			c: func() *uasc.Config {
				c := DefaultClientConfig()
				c.SecurityMode = ua.MessageSecurityModeSign
				return c
			}(),
		},
		{
			name: `SecurityModeString("bad")`,
			opt:  SecurityModeString("bad"),
			c: func() *uasc.Config {
				c := DefaultClientConfig()
				c.SecurityMode = ua.MessageSecurityModeInvalid
				return c
			}(),
		},
		{
			name: `SecurityModeString("None")`,
			opt:  SecurityModeString("None"),
			c: func() *uasc.Config {
				c := DefaultClientConfig()
				c.SecurityMode = ua.MessageSecurityModeNone
				return c
			}(),
		},
		{
			name: `SecurityModeString("Sign")`,
			opt:  SecurityModeString("Sign"),
			c: func() *uasc.Config {
				c := DefaultClientConfig()
				c.SecurityMode = ua.MessageSecurityModeSign
				return c
			}(),
		},
		{
			name: `SecurityModeString("SignAndEncrypt")`,
			opt:  SecurityModeString("SignAndEncrypt"),
			c: func() *uasc.Config {
				c := DefaultClientConfig()
				c.SecurityMode = ua.MessageSecurityModeSignAndEncrypt
				return c
			}(),
		},
		{
			name: `SecurityPolicy("Basic128Rsa15")`,
			opt:  SecurityPolicy("Basic128Rsa15"),
			c: func() *uasc.Config {
				c := DefaultClientConfig()
				c.SecurityPolicyURI = ua.SecurityPolicyURIBasic128Rsa15
				return c
			}(),
		},
		{
			name: `SecurityPolicy("Basic256")`,
			opt:  SecurityPolicy("Basic256"),
			c: func() *uasc.Config {
				c := DefaultClientConfig()
				c.SecurityPolicyURI = ua.SecurityPolicyURIBasic256
				return c
			}(),
		},
		{
			name: `SecurityPolicy("Basic256Sha256")`,
			opt:  SecurityPolicy("Basic256Sha256"),
			c: func() *uasc.Config {
				c := DefaultClientConfig()
				c.SecurityPolicyURI = ua.SecurityPolicyURIBasic256Sha256
				return c
			}(),
		},
		{
			name: `SecurityPolicy("Aes128_Sha256_RsaOaep")`,
			opt:  SecurityPolicy("Aes128_Sha256_RsaOaep"),
			c: func() *uasc.Config {
				c := DefaultClientConfig()
				c.SecurityPolicyURI = ua.SecurityPolicyURIAes128Sha256RsaOaep
				return c
			}(),
		},
		{
			name: `SecurityPolicy("Aes256_Sha256_RsaPss")`,
			opt:  SecurityPolicy("Aes256_Sha256_RsaPss"),
			c: func() *uasc.Config {
				c := DefaultClientConfig()
				c.SecurityPolicyURI = ua.SecurityPolicyURIAes256Sha256RsaPss
				return c
			}(),
		},
		{
			name: `SessionTimeout(5s)`,
			opt:  SessionTimeout(5 * time.Second),
			sc: func() *uasc.SessionConfig {
				sc := DefaultSessionConfig()
				sc.SessionTimeout = 5 * time.Second
				return sc
			}(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ttc := tt.c
			if ttc == nil {
				ttc = DefaultClientConfig()
			}
			ttsc := tt.sc
			if ttsc == nil {
				ttsc = DefaultSessionConfig()
			}

			c, sc := ApplyConfig(DefaultClientConfig(), DefaultSessionConfig(), tt.opt)
			verify.Values(t, "", c, ttc)
			verify.Values(t, "", sc, ttsc)
		})
	}
}
