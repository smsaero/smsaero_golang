package smsaero_golang

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

var GateUrls = []string{"gate.smsaero.ru/v2/", "gate.smsaero.org/v2/", "gate.smsaero.net/v2/", "gate.smsaero.uz/v2/"}

type ErrorResponse struct {
	Success bool   `json:"success,omitempty"`
	Message string `json:"message,omitempty"`
}

func (e ErrorResponse) IsErrorResponse() bool {
	return e.Success == false
}

func (e ErrorResponse) GetError() error {
	return fmt.Errorf(e.Message)
}

type WithErrorResponse interface {
	IsErrorResponse() bool
	GetError() error
}

type Client struct {
	username string
	password string
	client   *http.Client
}

func (c *Client) executeRequest(path string, destination WithErrorResponse, params url.Values) error {
	for i, baseUrl := range GateUrls {
		fullURL := "https://" + baseUrl + path

		if params == nil {
			params = make(url.Values, 0)
		}

		req, err := http.NewRequest("POST", fullURL, strings.NewReader(params.Encode()))
		if err != nil {
			return err
		}

		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

		req.SetBasicAuth(c.username, c.password)

		resp, err := c.client.Do(req)
		if err != nil {
			if (i + 1) != len(GateUrls) {
				continue
			}
			return err
		}
		defer resp.Body.Close()

		if err := json.NewDecoder(resp.Body).Decode(destination); err != nil {
			return err
		}

		break
	}
	return nil
}

func NewSmsAeroClient(username, password string) *Client {
	httpClient := http.DefaultClient
	return &Client{
		username: username,
		password: password,
		client:   httpClient,
	}
}
