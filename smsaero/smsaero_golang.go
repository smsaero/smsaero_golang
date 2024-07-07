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

var GateUrls = []string{
	"gate.smsaero.ru/v2/",
	"gate.smsaero.org/v2/",
	"gate.smsaero.net/v2/",
}

type ErrorResponse struct {
	Success bool   `json:"success,omitempty"`
	Message string `json:"message,omitempty"`
}

func (e ErrorResponse) IsErrorResponse() bool {
	return !e.Success
}

func (e ErrorResponse) GetError() error {
	if e.IsErrorResponse() {
		return fmt.Errorf(e.Message)
	}
	return nil
}

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
}

func NewSmsAeroClient(username, password string, opts ...ClientOption) *Client {
	client := &Client{
		username:        username,
		password:        password,
		client:          http.DefaultClient,
		ctx:             context.Background(),
		phoneValidation: true,
		testMode:        false,
		sign:            "Sms Aero",
		httpProtocol:    "https",
	}
	for _, opt := range opts {
		opt(client)
	}
	return client
}

func (c *Client) log(format string, v ...interface{}) {
	if c.logger != nil {
		c.logger.Printf(format, v...)
	}
}

func (c *Client) executeRequest(path string, destination WithErrorResponse, params url.Values) error {
	for i, baseUrl := range GateUrls {
		fullURL := c.httpProtocol + "://" + baseUrl + path

		c.log("Sending request to %s", fullURL)

		if params == nil {
			params = make(url.Values)
		}

		req, err := http.NewRequestWithContext(c.ctx, "POST", fullURL, strings.NewReader(params.Encode()))
		if err != nil {
			c.log("Error sending request to %s: %v", fullURL, err)
			return err
		}

		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		req.Header.Add("User-Agent", "SAGoClient/2.0.0")
		req.SetBasicAuth(c.username, c.password)

		resp, err := c.client.Do(req)
		if err != nil {
			if (i + 1) != len(GateUrls) {
				continue
			}
			return err
		}
		defer resp.Body.Close()

		c.log("Received response from %s", fullURL)
		if err := json.NewDecoder(resp.Body).Decode(destination); err != nil {
			fmt.Println(resp.Body)
			return err
		}

		break
	}
	return nil
}

func (c *Client) getApiPath(normalPath, testPath string) string {
	if c.testMode {
		return testPath
	}
	return normalPath
}
