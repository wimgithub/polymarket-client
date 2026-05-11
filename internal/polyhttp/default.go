package polyhttp

import (
	"net"
	"net/http"
	"time"
)

// NewDefaultHTTPClient creates the default HTTP client used by Polymarket API clients.
func NewDefaultHTTPClient() *http.Client {
	return &http.Client{
		Timeout:   15 * time.Second,
		Transport: NewDefaultTransport(),
	}
}

// NewDefaultTransport creates the default HTTP transport used by Polymarket API clients.
func NewDefaultTransport() *http.Transport {
	return &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   5 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
		TLSHandshakeTimeout:   5 * time.Second,
		ResponseHeaderTimeout: 10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
		MaxIdleConns:          100,
		MaxIdleConnsPerHost:   20,
		IdleConnTimeout:       90 * time.Second,
	}
}
