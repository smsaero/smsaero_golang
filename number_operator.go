package smsaero_golang

import "net/url"

type NumberOperator struct {
	Number         string
	Operator       int
	ExtendOperator string
}

type NumberOperatorMsg struct {
	Data NumberOperator
	ErrorResponse
}

func (c *Client) NumberOperator(number string) (NumberOperator, error) {
	response := new(NumberOperatorMsg)
	empty := NumberOperator{}

	data := url.Values{}
	data.Set("number", number)

	if err := c.executeRequest(`number/operator`, response, data); err != nil {
		return empty, err
	}

	if response.IsErrorResponse() {
		return empty, response.GetError()
	}

	return response.Data, nil
}
