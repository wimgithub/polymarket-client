package polyhttp

import (
	"errors"
	"net/http"
	"testing"
	"time"
)

type timeoutError struct{}

func (timeoutError) Error() string   { return "timeout" }
func (timeoutError) Timeout() bool   { return true }
func (timeoutError) Temporary() bool { return false }

func TestNewDefaultHTTPClient(t *testing.T) {
	client := NewDefaultHTTPClient()
	if client.Timeout != 15*time.Second {
		t.Fatalf("Timeout = %s, want %s", client.Timeout, 15*time.Second)
	}
	transport, ok := client.Transport.(*http.Transport)
	if !ok {
		t.Fatalf("Transport = %T, want *http.Transport", client.Transport)
	}
	if transport.Proxy == nil {
		t.Fatal("Proxy is nil, want environment proxy support")
	}
	if transport.DialContext == nil {
		t.Fatal("DialContext is nil")
	}
	if transport.TLSHandshakeTimeout != 5*time.Second {
		t.Fatalf("TLSHandshakeTimeout = %s, want %s", transport.TLSHandshakeTimeout, 5*time.Second)
	}
	if transport.ResponseHeaderTimeout != 10*time.Second {
		t.Fatalf("ResponseHeaderTimeout = %s, want %s", transport.ResponseHeaderTimeout, 10*time.Second)
	}
	if transport.ExpectContinueTimeout != time.Second {
		t.Fatalf("ExpectContinueTimeout = %s, want %s", transport.ExpectContinueTimeout, time.Second)
	}
	if transport.MaxIdleConns != 100 {
		t.Fatalf("MaxIdleConns = %d, want 100", transport.MaxIdleConns)
	}
	if transport.MaxIdleConnsPerHost != 20 {
		t.Fatalf("MaxIdleConnsPerHost = %d, want 20", transport.MaxIdleConnsPerHost)
	}
	if transport.IdleConnTimeout != 90*time.Second {
		t.Fatalf("IdleConnTimeout = %s, want %s", transport.IdleConnTimeout, 90*time.Second)
	}
}

func TestNewDefaultHTTPClientReturnsIndependentClients(t *testing.T) {
	first := NewDefaultHTTPClient()
	second := NewDefaultHTTPClient()
	if first == second {
		t.Fatal("NewDefaultHTTPClient returned the same client instance")
	}
	if first.Transport == second.Transport {
		t.Fatal("NewDefaultHTTPClient returned clients sharing the same transport")
	}
}

func TestIsTimeout(t *testing.T) {
	if !IsTimeout(errors.Join(errors.New("request failed"), timeoutError{})) {
		t.Fatal("IsTimeout returned false for wrapped net.Error timeout")
	}
	if IsTimeout(errors.New("request failed")) {
		t.Fatal("IsTimeout returned true for non-timeout error")
	}
}
