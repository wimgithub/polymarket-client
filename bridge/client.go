package bridge

import (
	"context"
	"net/http"
	"time"

	"github.com/bububa/polymarket-client/internal/polyhttp"
)

const DefaultHost = "https://bridge.polymarket.com"

type Client struct {
	host string
	http *polyhttp.Client
}

// Config configures a Bridge API client.
type Config struct {
	Host       string
	HTTPClient *http.Client
	UserAgent  string
}

// New creates a Bridge API client.
func New(config Config) *Client {
	if config.Host == "" {
		config.Host = DefaultHost
	}
	if config.HTTPClient == nil {
		config.HTTPClient = &http.Client{Timeout: 15 * time.Second}
	}
	if config.UserAgent == "" {
		config.UserAgent = "polymarket-client-go/bridge"
	}
	return &Client{
		host: config.Host,
		http: &polyhttp.Client{BaseURL: config.Host, HTTPClient: config.HTTPClient, UserAgent: config.UserAgent},
	}
}

// Host returns the configured Bridge API host.
func (c *Client) Host() string { return c.host }

func (c *Client) GetSupportedAssets(ctx context.Context) (*SupportedAssetsResponse, error) {
	var out SupportedAssetsResponse
	return &out, c.http.GetJSON(ctx, "/supported-assets", nil, polyhttp.AuthNone, &out)
}

func (c *Client) CreateDepositAddress(ctx context.Context, address string) (*DepositResponse, error) {
	var out DepositResponse
	return &out, c.http.PostJSON(ctx, "/deposit", DepositRequest{Address: address}, polyhttp.AuthNone, &out)
}

func (c *Client) GetStatus(ctx context.Context, address string) (*StatusResponse, error) {
	var out StatusResponse
	return &out, c.http.GetJSON(ctx, "/status/"+address, nil, polyhttp.AuthNone, &out)
}

func (c *Client) GetQuote(ctx context.Context, req QuoteRequest) (*QuoteResponse, error) {
	var out QuoteResponse
	return &out, c.http.PostJSON(ctx, "/quote", req, polyhttp.AuthNone, &out)
}

func (c *Client) Withdraw(ctx context.Context, req WithdrawRequest) (*WithdrawResponse, error) {
	var out WithdrawResponse
	return &out, c.http.PostJSON(ctx, "/withdraw", req, polyhttp.AuthNone, &out)
}
