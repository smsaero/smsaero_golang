package smsaero_golang

import (
	"net/url"
	"strconv"
)

type NumberOperator struct {
	Number         string `json:"number"`
	Operator       int    `json:"operator"`
	ExtendOperator string `json:"extendOperator"`
}

type NumberOperatorMsg struct {
	Data NumberOperator `json:"data"`
	ErrorResponse
}

func (c *Client) NumberOperator(number int) (NumberOperator, error) {
	response := new(NumberOperatorMsg)
	empty := NumberOperator{}

	data := url.Values{}
	data.Set("number", strconv.Itoa(number))

	if err := c.executeRequest(`number/operator`, response, data); err != nil {
		return empty, err
	}

	if response.IsErrorResponse() {
		return empty, response.GetError()
	}

	return response.Data, nil
}
