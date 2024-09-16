package smsaero_golang

import (
	"fmt"
	"net/url"
)

type SmsStatus struct {
	Id           int
	From         string
	Number       int
	Text         string
	Status       int
	ExtendStatus string
	Channel      string
	Cost         string
	DateCreate   int
	DateSend     int
	DateAnswer   int
}

type SmsStatusMsg struct {
	Data SmsStatus
	ErrorResponse
}

func (c *Client) SmsStatus(id int) (SmsStatus, error) {
	response := new(SmsStatusMsg)
	empty := SmsStatus{}

	data := url.Values{}
	data.Set("id", fmt.Sprintf("%d", id))

	path := c.getApiPath("sms/status", "sms/teststatus")

	if err := c.executeRequest(path, response, data); err != nil {
		return empty, err
	}

	if response.IsErrorResponse() {
		return empty, response.GetError()
	}

	return response.Data, nil
}
