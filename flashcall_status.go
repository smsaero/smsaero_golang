package smsaero_golang

import "net/url"

type FlashCallStatus struct {
	Id         int
	Status     int
	Code       string
	Phone      string
	Cost       string
	TimeCreate int
	TimeUpdate int
	TimeSend   int
}

type FlashCallStatusMsg struct {
	Data FlashCallStatus
	ErrorResponse
}

func (c *Client) FlashCallStatus(id string) (FlashCallStatus, error) {
	response := new(FlashCallStatusMsg)
	empty := FlashCallStatus{}

	data := url.Values{}
	data.Set("id", id)

	if err := c.executeRequest(`flashcall/status`, response, data); err != nil {
		return empty, err
	}

	if response.IsErrorResponse() {
		return empty, response.GetError()
	}

	return response.Data, nil
}
