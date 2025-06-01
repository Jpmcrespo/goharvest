package utlsclient

import (
	"context"
	"net/http"
)

// Options for spoofed requests
type RequestOptions struct {
	URL     string
	Headers map[string]string
	Timeout int    // in seconds
	JA3     string // e.g. "chrome_112"
}

// FetchURL performs a spoofed GET request using uTLS and HTTP/2
func FetchURL(ctx context.Context, opts RequestOptions) (*http.Response, error) {
	client, err := NewSpoofedHTTPClient(opts)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, "GET", opts.URL, nil)
	if err != nil {
		return nil, err
	}

	for k, v := range opts.Headers {
		req.Header.Set(k, v)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
