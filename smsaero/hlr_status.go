package smsaero_golang

import (
	"fmt"
	"net/url"
)

type HlrStatusMsg struct {
	Data HlrCheck
	ErrorResponse
}

func (c *Client) HlrStatus(id int) (HlrCheck, error) {
	response := new(HlrStatusMsg)
	empty := HlrCheck{}

	data := url.Values{}
	data.Set("id", fmt.Sprintf("%d", id))

	if err := c.executeRequest(`hlr/status`, response, data); err != nil {
		return empty, err
	}

	if response.IsErrorResponse() {
		return empty, response.GetError()
	}

	return response.Data, nil
}
