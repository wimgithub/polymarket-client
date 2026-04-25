package ws

import (
	"net/http"
	"strings"
	"time"
)

const DefaultHost = "wss://ws-subscriptions-clob.polymarket.com"

type Client struct {
	host       string
	httpClient *http.Client
}

// Config configures a CLOB WebSocket client.
type Config struct {
	Host       string
	HTTPClient *http.Client
}

// New creates a CLOB WebSocket client.
func New(config Config) *Client {
	if config.Host == "" {
		config.Host = DefaultHost
	}
	if config.HTTPClient == nil {
		config.HTTPClient = &http.Client{Timeout: 15 * time.Second}
	}
	return &Client{host: strings.TrimRight(config.Host, "/"), httpClient: config.HTTPClient}
}

// Host returns the configured WebSocket host.
func (c *Client) Host() string { return c.host }

// MarketURL returns the market-channel WebSocket URL.
func (c *Client) MarketURL() string { return c.host + "/ws/market" }

// UserURL returns the user-channel WebSocket URL.
func (c *Client) UserURL() string { return c.host + "/ws/user" }

// SportsURL returns the public sports-channel WebSocket URL.
func (c *Client) SportsURL() string { return c.host + "/ws" }
