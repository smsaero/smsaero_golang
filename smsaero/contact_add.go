package smsaero_golang

import "net/url"

type ContactAddMsg struct {
	Data BlackListAdd
	ErrorResponse
}

func (c *Client) ContactAdd(number string, groupId string) (BlackListAdd, error) {
	response := new(ContactAddMsg)
	empty := BlackListAdd{}

	data := url.Values{}
	data.Set("number", number)
	data.Set("groupId", groupId)

	if err := c.executeRequest(`contact/add`, response, data); err != nil {
		return empty, err
	}

	if response.IsErrorResponse() {
		return empty, response.GetError()
	}

	return response.Data, nil
}
