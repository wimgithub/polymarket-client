package polyhttp

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type AuthLevel int

const (
	AuthNone AuthLevel = iota
	AuthL1
	AuthL2
)

type HeaderFunc func(ctx context.Context, method, path string, body []byte, level AuthLevel, nonce *int64) (map[string]string, error)

type Client struct {
	BaseURL    string
	HTTPClient *http.Client
	UserAgent  string
	Headers    HeaderFunc
}

type APIError struct {
	StatusCode  int
	Message     string
	Body        []byte
	RequestBody []byte
}

func (e *APIError) Error() string {
	if e.Message != "" {
		return fmt.Sprintf("polymarket API error: status %d: %s", e.StatusCode, e.Message)
	}
	return fmt.Sprintf("polymarket API error: status %d", e.StatusCode)
}

func (e *APIError) HTTPStatus() int { return e.StatusCode }

func (c *Client) GetJSON(ctx context.Context, path string, query url.Values, auth AuthLevel, out any) error {
	return c.DoJSON(ctx, http.MethodGet, path, query, nil, auth, nil, out)
}

func (c *Client) PostJSON(ctx context.Context, path string, body any, auth AuthLevel, out any) error {
	return c.DoJSON(ctx, http.MethodPost, path, nil, body, auth, nil, out)
}

func (c *Client) DeleteJSON(ctx context.Context, path string, body any, auth AuthLevel, out any) error {
	return c.DoJSON(ctx, http.MethodDelete, path, nil, body, auth, nil, out)
}

func (c *Client) DoJSON(ctx context.Context, method, path string, query url.Values, body any, auth AuthLevel, nonce *int64, out any) error {
	requestBody, err := marshalBody(body)
	if err != nil {
		return err
	}
	fullURL := strings.TrimRight(c.BaseURL, "/") + path
	if len(query) > 0 {
		fullURL += "?" + query.Encode()
	}
	var bodyReader io.Reader
	if len(requestBody) > 0 {
		bodyReader = bytes.NewReader(requestBody)
	}
	req, err := http.NewRequestWithContext(ctx, method, fullURL, bodyReader)
	if err != nil {
		return fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("User-Agent", c.UserAgent)
	if len(requestBody) > 0 {
		req.Header.Set("Content-Type", "application/json")
	}
	if c.Headers != nil {
		headers, err := c.Headers(ctx, method, path, requestBody, auth, nonce)
		if err != nil {
			return err
		}
		for key, value := range headers {
			req.Header.Set(key, value)
		}
	}
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return fmt.Errorf("perform request: %w", err)
	}
	defer resp.Body.Close()
	payload, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("read response: %w", err)
	}
	if resp.StatusCode >= http.StatusBadRequest {
		return newAPIError(resp, payload, requestBody)
	}
	if out == nil || len(bytes.TrimSpace(payload)) == 0 {
		return nil
	}
	switch target := out.(type) {
	case *int64:
		parsed, err := strconv.ParseInt(strings.TrimSpace(string(payload)), 10, 64)
		if err != nil {
			return fmt.Errorf("decode integer response: %w", err)
		}
		*target = parsed
		return nil
	case *string:
		var decoded string
		if err := json.Unmarshal(payload, &decoded); err == nil {
			*target = decoded
		} else {
			*target = strings.TrimSpace(string(payload))
		}
		return nil
	default:
		if err := json.Unmarshal(payload, out); err != nil {
			return fmt.Errorf("decode response: %w", err)
		}
		return nil
	}
}

func marshalBody(body any) ([]byte, error) {
	if body == nil {
		return nil, nil
	}
	switch typed := body.(type) {
	case []byte:
		return typed, nil
	case string:
		return []byte(typed), nil
	default:
		payload, err := json.Marshal(typed)
		if err != nil {
			return nil, fmt.Errorf("marshal request body: %w", err)
		}
		return payload, nil
	}
}

func newAPIError(resp *http.Response, body, requestBody []byte) *APIError {
	err := &APIError{StatusCode: resp.StatusCode, Body: bytes.Clone(body), RequestBody: bytes.Clone(requestBody)}
	var payload struct {
		Error any `json:"error"`
	}
	if json.Unmarshal(body, &payload) == nil && payload.Error != nil {
		err.Message = fmt.Sprint(payload.Error)
		return err
	}
	err.Message = strings.TrimSpace(string(body))
	return err
}
