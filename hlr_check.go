package smsaero_golang

import "net/url"

type HlrCheck struct {
	Id              int
	Number          string
	HlrStatus       int
	ExtendHlrStatus string
}

type HlrCheckMsg struct {
	Data HlrCheck
	ErrorResponse
}

func (c *Client) HlrCheck(number string) (HlrCheck, error) {
	response := new(HlrCheckMsg)
	empty := HlrCheck{}

	data := url.Values{}
	data.Set("number", number)

	if err := c.executeRequest(`hlr/check`, response, data); err != nil {
		return empty, err
	}

	if response.IsErrorResponse() {
		return empty, response.GetError()
	}

	return response.Data, nil
}
