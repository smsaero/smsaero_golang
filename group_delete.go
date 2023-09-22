package smsaero_golang

import "net/url"

type GroupDeleteMsg struct {
	Data interface{}
	ErrorResponse
}

func (c *Client) GroupDelete(id string) (bool, error) {
	response := new(GroupDeleteMsg)

	data := url.Values{}
	data.Set("id", id)

	if err := c.executeRequest(`group/delete`, response, data); err != nil {
		return false, err
	}

	if response.IsErrorResponse() {
		return false, response.GetError()
	}

	return true, nil
}
