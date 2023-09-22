package smsaero_golang

import "net/url"

type BlackListDeleteMsg struct {
	Data interface{}
	ErrorResponse
}

func (c *Client) BlackListDelete(id string) (bool, error) {
	response := new(BlackListDeleteMsg)

	data := url.Values{}
	data.Set("id", id)

	if err := c.executeRequest(`blacklist/delete`, response, data); err != nil {
		return false, err
	}

	if response.IsErrorResponse() {
		return false, response.GetError()
	}

	return true, nil
}
