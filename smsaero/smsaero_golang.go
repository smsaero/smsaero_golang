package smsaero_golang

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	badRequestStatusCode = 400
)

// GateUrls contains list of available SMS Aero gateways
var GateUrls = []string{
	"gate.smsaero.ru/v2/",
	"gate.smsaero.org/v2/",
	"gate.smsaero.net/v2/",
}

// ErrorResponse represents base error response structure
type ErrorResponse struct {
	Success bool   `json:"success,omitempty"`
	Message string `json:"message,omitempty"`
	Result  string `json:"result,omitempty"`
	Reason  string `json:"reason,omitempty"`
}

// IsErrorResponse checks if response is an error
func (e ErrorResponse) IsErrorResponse() bool {
	return !e.Success || e.Result == "reject" || e.Result == "no credits"
}

// GetError returns error based on response
func (e ErrorResponse) GetError() error {
	if !e.IsErrorResponse() {
		return nil
	}

	switch e.Result {
	case "no credits":
		return NewNoMoneyError(e.Message, 0.0)
	case "reject":
		return NewAPIError(e.Reason, badRequestStatusCode)
	default:
		if e.Message != "" {
			return NewAPIError(e.Message, badRequestStatusCode)
		}

		return NewAPIError("unknown API error", badRequestStatusCode)
	}
}

// WithErrorResponse interface for structures with error handling
type WithErrorResponse interface {
	IsErrorResponse() bool
	GetError() error
}

func WithHTTPClient(httpClient *http.Client) ClientOption {
	return func(c *Client) {
		c.client = httpClient
	}
}

func WithContext(ctx context.Context) ClientOption {
	return func(c *Client) {
		c.ctx = ctx
	}
}

func WithTimeout(timeout time.Duration) ClientOption {
	return func(c *Client) {
		c.client.Timeout = timeout
	}
}

func WithPhoneValidation(enabled bool) ClientOption {
	return func(c *Client) {
		c.phoneValidation = enabled
	}
}

func WithTest(enabled bool) ClientOption {
	return func(c *Client) {
		c.testMode = enabled
	}
}

func WithSign(sign string) ClientOption {
	return func(c *Client) {
		c.sign = sign
	}
}

type Logger interface {
	Printf(format string, v ...interface{})
}

func WithLogger(logger Logger) ClientOption {
	return func(c *Client) {
		c.logger = logger
	}
}

type ClientOption func(*Client)

// Client represents client for working with SMS Aero API
type Client struct {
	username        string
	password        string
	client          *http.Client
	ctx             context.Context
	phoneValidation bool
	testMode        bool
	logger          Logger
	sign            string
	httpProtocol    string
	validator       *Validator
	gateUrls        []string
}

// NewSmsAeroClient creates new SMS Aero client
func NewSmsAeroClient(username, password string, opts ...ClientOption) (*Client, error) {
	// Validate credentials
	validator := NewValidator(true)
	if err := validator.ValidateCredentials(username, password); err != nil {
		return nil, err
	}

	client := &Client{
		username:        username,
		password:        password,
		client:          http.DefaultClient,
		ctx:             context.Background(),
		phoneValidation: true,
		testMode:        false,
		sign:            "Sms Aero",
		httpProtocol:    "https",
		validator:       validator,
		gateUrls:        make([]string, len(GateUrls)),
	}
	copy(client.gateUrls, GateUrls)

	for _, opt := range opts {
		opt(client)
	}

	// Update validator with client settings
	client.validator = NewValidator(client.phoneValidation)

	return client, nil
}

func (c *Client) logf(format string, v ...interface{}) {
	if c.logger != nil {
		c.logger.Printf(format, v...)
	}
}

// executeRequest executes HTTP request to SMS Aero API
func (c *Client) executeRequest(path string, destination WithErrorResponse, params url.Values) error {
	var lastErr error
	protocol := c.httpProtocol // Start with current protocol

	for gatewayIndex, baseURL := range c.gateUrls {
		// Ensure baseURL ends with slash and path starts with slash
		if !strings.HasSuffix(baseURL, "/") {
			baseURL += "/"
		}
		path = strings.TrimPrefix(path, "/")
		fullURL := protocol + "://" + baseURL + path

		c.logf("Sending request to %s", fullURL)

		if params == nil {
			params = make(url.Values)
		}

		req, err := http.NewRequestWithContext(c.ctx, http.MethodPost, fullURL, strings.NewReader(params.Encode()))
		if err != nil {
			c.logf("Error creating request to %s: %v", fullURL, err)
			_ = NewConnectionError(fmt.Sprintf("error creating request: %v", err))

			continue
		}

		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		req.Header.Add("User-Agent", "SAGoClient/2.0.0")
		req.SetBasicAuth(c.username, c.password)

		resp, err := c.client.Do(req)
		if err != nil {
			c.logf("Error executing request to %s: %v", fullURL, err)

			// Check if this is SSL/TLS error
			if isSSLError(err) && protocol == "https" {
				c.logf("SSL error, switching to HTTP for gateway %s", baseURL)
				protocol = "http"
				// Retry with HTTP for the same gateway
				_ = gatewayIndex // Decrease index to retry the same gateway

				continue
			}

			lastErr = NewConnectionError(fmt.Sprintf("error executing request: %v", err))

			// If this is the last gateway, return error
			if (gatewayIndex + 1) == len(c.gateUrls) {
				return lastErr
			}

			continue
		}
		defer resp.Body.Close()

		c.logf("Received response from %s (status: %d)", fullURL, resp.StatusCode)

		// Check response status code
		if resp.StatusCode >= badRequestStatusCode {
			lastErr = NewAPIError(fmt.Sprintf("HTTP error %d", resp.StatusCode), resp.StatusCode)
			if (gatewayIndex + 1) == len(c.gateUrls) {
				return lastErr
			}

			continue
		}

		if err := json.NewDecoder(resp.Body).Decode(destination); err != nil {
			c.logf("Error decoding response from %s: %v", fullURL, err)
			lastErr = NewAPIError(fmt.Sprintf("error decoding response: %v", err), resp.StatusCode)
			if (gatewayIndex + 1) == len(c.gateUrls) {
				return lastErr
			}

			continue
		}

		// Successful response
		return nil
	}

	// If all gateways are unavailable
	return ErrAllGatesUnavailable
}

// isSSLError checks if error is SSL/TLS error
func isSSLError(err error) bool {
	if err == nil {
		return false
	}

	errStr := err.Error()
	// Check various types of SSL/TLS errors
	return strings.Contains(errStr, "tls:") ||
		strings.Contains(errStr, "x509:") ||
		strings.Contains(errStr, "certificate") ||
		strings.Contains(errStr, "SSL") ||
		strings.Contains(errStr, "TLS") ||
		strings.Contains(errStr, "handshake failure") ||
		strings.Contains(errStr, "certificate verify failed") ||
		strings.Contains(errStr, "server gave HTTP response to HTTPS client")
}

func (c *Client) getAPIPath(normalPath, testPath string) string {
	if c.testMode {
		return testPath
	}

	return normalPath
}

// SetGateUrls sets list of gateway URLs for client
func (c *Client) SetGateUrls(urls []string) {
	c.gateUrls = make([]string, len(urls))
	copy(c.gateUrls, urls)
}

// SetHTTPProtocol sets protocol for HTTP requests
func (c *Client) SetHTTPProtocol(protocol string) {
	c.httpProtocol = protocol
}
