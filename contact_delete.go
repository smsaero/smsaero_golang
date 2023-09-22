package smsaero_golang

import "net/url"

type ContactDeleteMsg struct {
	Data interface{}
	ErrorResponse
}

func (c *Client) ContactDelete(id string) (bool, error) {
	response := new(ContactDeleteMsg)

	data := url.Values{}
	data.Set("id", id)

	if err := c.executeRequest(`contact/delete`, response, data); err != nil {
		return false, err
	}

	if response.IsErrorResponse() {
		return false, response.GetError()
	}

	return true, nil
}
