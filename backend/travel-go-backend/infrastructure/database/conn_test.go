package pg

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseConnectionString(t *testing.T) {
	tc := []struct {
		Connection Connection
		Expected   string
	}{
		{
			Connection: Connection{
				Host:     "localhost",
				Port:     "123456",
				User:     "admin",
				Password: "admin",
				Database: "test",
				SSLMode:  "disable",
			},
			Expected: "postgresql://admin:admin@localhost:123456/test?sslmode=disable",
		},
		{
			Connection: Connection{
				Host:                        "localhost",
				Port:                        "123456",
				User:                        "admin",
				Password:                    "admin",
				Database:                    "test",
				SSLMode:                     "require",
				SSLCertAuthorityCertificate: "ca.crt",
				SSLPublicCertificate:        "public.crt",
				SSLPrivateKey:               "private.key",
			},
			Expected: "postgresql://admin:admin@localhost:123456/test?sslmode=require&sslrootcert=ca.crt&sslcert=public.crt&sslkey=private.key",
		},
		{
			Connection: Connection{
				Host:                        "localhost",
				Port:                        "123456",
				User:                        "admin",
				Password:                    "admin",
				Database:                    "test",
				SSLMode:                     "require",
				SSLCertAuthorityCertificate: "ca.crt",
				SSLPublicCertificate:        "public.crt",
				SSLPrivateKey:               "private.key",
				FallbackConnections: []FallbackConnection{
					{Host: "localhost", Port: "12345"},
				},
			},
			Expected: "postgresql://admin:admin@localhost:123456,localhost:12345/test?sslmode=require&sslrootcert=ca.crt&sslcert=public.crt&sslkey=private.key",
		},
		{
			Connection: Connection{
				Host:                        "localhost",
				Port:                        "123456",
				User:                        "admin",
				Password:                    "admin",
				Database:                    "test",
				SSLMode:                     "require",
				SSLCertAuthorityCertificate: "ca.crt",
				SSLPublicCertificate:        "public.crt",
				SSLPrivateKey:               "private.key",
				FallbackConnections: []FallbackConnection{
					{Host: "localhost", Port: "12345"},
				},
			},
			Expected: "postgresql://admin:admin@localhost:123456,localhost:12345/test?sslmode=require&sslrootcert=ca.crt&sslcert=public.crt&sslkey=private.key",
		},
	}

	for _, v := range tc {
		require.Equal(t, v.Expected, v.Connection.ToConnectionString())
	}
}
