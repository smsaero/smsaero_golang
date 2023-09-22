package smsaero_golang

import "net/url"

type HlrStatusMsg struct {
	Data HlrCheck
	ErrorResponse
}

func (c *Client) HlrStatus(id string) (HlrCheck, error) {
	response := new(HlrStatusMsg)
	empty := HlrCheck{}

	data := url.Values{}
	data.Set("id", id)

	if err := c.executeRequest(`hlr/status`, response, data); err != nil {
		return empty, err
	}

	if response.IsErrorResponse() {
		return empty, response.GetError()
	}

	return response.Data, nil
}
