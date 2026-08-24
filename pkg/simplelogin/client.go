package simplelogin

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

const (
	BaseURL = "https://app.simplelogin.io/api"
)

// Client represents a SimpleLogin API client
// It provides methods to interact with the SimpleLogin API
type Client struct {
	baseURL    string
	apiKey     string
	httpClient *http.Client
	ctx        context.Context
}

// NewClient creates a new SimpleLogin API client with a custom base URL
// This is useful for testing or when using a different SimpleLogin instance
// baseURL: The custom base URL for the API
// apiKey: The API key for authentication
// Returns a configured client or an error if validation fails
func NewClient(baseURL *string, apiKey string) (*Client, error) {
	base := BaseURL

	if apiKey == "" {
		return nil, &ValidationError{Field: "apiKey", Message: "API key is required"}
	}
	if baseURL != nil {
		base = *baseURL
	}

	if base == "" {
		return nil, &ValidationError{Field: "baseURL", Message: "base URL is required"}
	}
	if _, err := url.Parse(base); err != nil {
		return nil, &ValidationError{Field: "baseURL", Message: fmt.Sprintf("invalid base URL: %s", err)}
	}

	client := &Client{
		baseURL: base,
		apiKey:  apiKey,
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
		ctx: context.Background(),
	}

	return client, nil
}

// WithContext returns a copy of the client configured to use the provided
// context for all subsequent API requests. This allows callers (e.g. CLI
// commands or tests) to control cancellation and timeouts.
func (c *Client) WithContext(ctx context.Context) APIClient {
	if ctx == nil {
		ctx = context.Background()
	}
	return &Client{
		baseURL:    c.baseURL,
		apiKey:     c.apiKey,
		httpClient: c.httpClient,
		ctx:        ctx,
	}
}

// doRequest performs an HTTP request with the API key
func (c *Client) doRequest(method, endpoint string, body io.Reader) (*http.Response, error) {
	return c.doRequestWithContext(c.ctx, method, endpoint, body)
}

// doRequestWithContext performs an HTTP request with context support
func (c *Client) doRequestWithContext(ctx context.Context, method, endpoint string, body io.Reader) (*http.Response, error) {
	url := fmt.Sprintf("%s%s", c.baseURL, endpoint)

	req, err := http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// SimpleLogin uses the non-standard "Authentication" header instead of the
	// more common "Authorization" header to pass the API key. Do not change this
	// unless the SimpleLogin API documentation explicitly requires it.
	req.Header.Set("Authentication", c.apiKey)
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to execute request: %w", err)
	}

	return resp, nil
}

// handleResponse handles the HTTP response and unmarshals JSON if needed
func (c *Client) handleResponse(resp *http.Response, v interface{}) error {
	defer func() {
		_ = resp.Body.Close()
	}()

	if resp.StatusCode >= 400 {
		body, _ := io.ReadAll(resp.Body)

		// Try to parse error response as JSON
		var apiError struct {
			Error string `json:"error"`
		}
		if err := json.Unmarshal(body, &apiError); err == nil && apiError.Error != "" {
			return &APIError{
				StatusCode: resp.StatusCode,
				Message:    apiError.Error,
				Body:       string(body),
			}
		}

		// Handle specific status codes
		switch resp.StatusCode {
		case http.StatusUnauthorized:
			return &AuthenticationError{Message: "Invalid API key"}
		case http.StatusTooManyRequests:
			retryAfter := 0
			if retryAfterStr := resp.Header.Get("Retry-After"); retryAfterStr != "" {
				if retry, err := strconv.Atoi(retryAfterStr); err == nil {
					retryAfter = retry
				}
			}
			return &RateLimitError{
				RetryAfter: retryAfter,
				Message:    "Rate limit exceeded",
			}
		default:
			return &APIError{
				StatusCode: resp.StatusCode,
				Body:       string(body),
			}
		}
	}

	if v != nil {
		if err := json.NewDecoder(resp.Body).Decode(v); err != nil {
			return fmt.Errorf("failed to decode response: %w", err)
		}
	}

	return nil
}
