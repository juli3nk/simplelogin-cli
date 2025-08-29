package simplelogin

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
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
	logger     *log.Logger
}

// NewClient creates a new SimpleLogin API client with a custom base URL
// This is useful for testing or when using a different SimpleLogin instance
// baseURL: The custom base URL for the API
// apiKey: The API key for authentication
// Returns a configured client or an error if validation fails
func NewClient(baseURL *string, apiKey string) (*Client, error) {
	url := BaseURL

	if apiKey == "" {
		return nil, &ValidationError{Field: "apiKey", Message: "API key is required"}
	}
	if baseURL != nil {
		url = *baseURL
	}

	client := &Client{
		baseURL: url,
		apiKey:  apiKey,
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
		logger: log.New(io.Discard, "", 0), // Default to no logging
	}

	return client, nil
}

// SetLogger sets a custom logger for the client
// This allows for custom logging configuration and output
// logger: The logger instance to use for client logging
func (c *Client) SetLogger(logger *log.Logger) {
	c.logger = logger
}

// doRequest performs an HTTP request with the API key
func (c *Client) doRequest(method, endpoint string, body io.Reader) (*http.Response, error) {
	return c.doRequestWithContext(context.Background(), method, endpoint, body)
}

// doRequestWithContext performs an HTTP request with context support
func (c *Client) doRequestWithContext(ctx context.Context, method, endpoint string, body io.Reader) (*http.Response, error) {
	url := fmt.Sprintf("%s%s", c.baseURL, endpoint)

	req, err := http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

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

// doRequestWithRetry performs an HTTP request with retry logic for rate limiting
func (c *Client) doRequestWithRetry(method, endpoint string, body io.Reader) (*http.Response, error) {
	return c.doRequestWithRetryAndContext(context.Background(), method, endpoint, body)
}

// doRequestWithRetryAndContext performs an HTTP request with retry logic and context support
func (c *Client) doRequestWithRetryAndContext(ctx context.Context, method, endpoint string, body io.Reader) (*http.Response, error) {
	maxRetries := 3

	for attempt := 0; attempt < maxRetries; attempt++ {
		resp, err := c.doRequestWithContext(ctx, method, endpoint, body)
		if err != nil {
			return nil, err
		}

		// If not rate limited, return immediately
		if resp.StatusCode != http.StatusTooManyRequests {
			return resp, nil
		}

		// Handle rate limiting
		retryAfter := 1
		if retryAfterStr := resp.Header.Get("Retry-After"); retryAfterStr != "" {
			if retry, err := strconv.Atoi(retryAfterStr); err == nil {
				retryAfter = retry
			}
		}

		resp.Body.Close()

		// Wait before retrying
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-time.After(time.Duration(retryAfter) * time.Second):
			continue
		}
	}

	return nil, fmt.Errorf("max retries exceeded")
}

// handleResponse handles the HTTP response and unmarshals JSON if needed
func (c *Client) handleResponse(resp *http.Response, v interface{}) error {
	defer resp.Body.Close()

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
