package smsaero_golang

import "net/url"

type BlackListAdd struct {
	Id     int
	Number string
	Fname  string
	Lname  string
	Sname  string
	Bday   string
	Sex    string
	Param  string
	Param2 string
}

type BlackListAddMsg struct {
	Data BlackListAdd
	ErrorResponse
}

func (c *Client) BlackListAdd(number string) (BlackListAdd, error) {
	response := new(BlackListAddMsg)
	empty := BlackListAdd{}

	data := url.Values{}
	data.Set("number", number)

	if err := c.executeRequest(`blacklist/add`, response, data); err != nil {
		return empty, err
	}

	if response.IsErrorResponse() {
		return empty, response.GetError()
	}

	return response.Data, nil
}
