package smsaero_golang

import "net/url"

type Send struct {
	Id           int
	From         string
	Number       string
	Text         string
	Status       int
	ExtendStatus string
	Channel      string
	Cost         float64
	DateCreate   int
	DateSend     int
}

type SendMsg struct {
	Data Send
	ErrorResponse
}

func (c *Client) Send(number string, text string, sign string) (Send, error) {
	response := new(SendMsg)
	empty := Send{}

	data := url.Values{}
	data.Set("number", number)
	data.Set("sign", sign)
	data.Set("text", text)

	if err := c.executeRequest(`sms/send`, response, data); err != nil {
		return empty, err
	}

	if response.IsErrorResponse() {
		return empty, response.GetError()
	}

	return response.Data, nil
}
