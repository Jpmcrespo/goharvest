package utlsclient

import (
	"context"
	"crypto/tls"
	"net"
	"net/http"
	"time"

	utls "github.com/refraction-networking/utls"
	"golang.org/x/net/http2"
)

// NewSpoofedHTTPClient returns a *http.Client using uTLS + HTTP/2
func NewSpoofedHTTPClient(opts RequestOptions) (*http.Client, error) {

	ja3Id, err := parseJA3(opts.JA3)
	if err != nil {
		return nil, err
	}

	dialTLS := func(ctx context.Context, network, addr string, cfg *tls.Config) (net.Conn, error) {
		rawConn, err := (&net.Dialer{Timeout: 10 * time.Second}).DialContext(ctx, network, addr)
		if err != nil {
			return nil, err
		}

		uTlsConfig := &utls.Config{
			ServerName:         cfg.ServerName,
			InsecureSkipVerify: cfg.InsecureSkipVerify,
			NextProtos:         []string{"h2"},
		}

		uConn := utls.UClient(rawConn, uTlsConfig, ja3Id)
		if err := uConn.Handshake(); err != nil {
			return nil, err
		}

		return uConn, nil
	}

	transport := &http2.Transport{
		DialTLSContext: dialTLS,
	}

	client := &http.Client{
		Transport: transport,
		Timeout:   time.Duration(opts.Timeout) * time.Second,
	}

	return client, nil
}
