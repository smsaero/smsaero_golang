package smsaero_golang

import "net/url"

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

func (c *Client) SmsStatus(id string) (SmsStatus, error) {
	response := new(SmsStatusMsg)
	empty := SmsStatus{}

	data := url.Values{}
	data.Set("id", id)

	if err := c.executeRequest(`sms/status`, response, data); err != nil {
		return empty, err
	}

	if response.IsErrorResponse() {
		return empty, response.GetError()
	}

	return response.Data, nil
}
