package smsaero_golang

import "net/url"

type FlashCall struct {
	Id         int
	Status     int
	Code       string
	Phone      string
	Cost       string
	TimeCreate int
	TimeUpdate int
}

type FlashCallMsg struct {
	Data FlashCall
	ErrorResponse
}

func (c *Client) FlashCall(phone string, code string) (FlashCall, error) {
	response := new(FlashCallMsg)
	empty := FlashCall{}

	data := url.Values{}
	data.Set("phone", phone)
	data.Set("code", code)

	if err := c.executeRequest(`flashcall/send`, response, data); err != nil {
		return empty, err
	}

	if response.IsErrorResponse() {
		return empty, response.GetError()
	}

	return response.Data, nil
}
