package smsaero_golang

import "net/url"

type GroupAdd struct {
	Id   int
	Name string
}

type GroupAddMsg struct {
	Data GroupAdd
	ErrorResponse
}

func (c *Client) GroupAdd(name string) (GroupAdd, error) {
	response := new(GroupAddMsg)
	empty := GroupAdd{}

	data := url.Values{}
	data.Set("name", name)

	if err := c.executeRequest(`group/add`, response, data); err != nil {
		return empty, err
	}

	if response.IsErrorResponse() {
		return empty, response.GetError()
	}

	return response.Data, nil
}
