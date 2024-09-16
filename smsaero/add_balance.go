package smsaero_golang

import (
	"fmt"
	"net/url"
)

type AddBalanceMsg struct {
	Data interface{}
	ErrorResponse
}

func (c *Client) AddBalance(sum float64, cardId int) (interface{}, error) {
	response := new(AddBalanceMsg)

	data := url.Values{}
	data.Set("sum", fmt.Sprintf("%.2f", sum))
	data.Set("cardId", fmt.Sprintf("%d", cardId))

	if err := c.executeRequest(`balance/add`, response, data); err != nil {
		return nil, err
	}

	if response.IsErrorResponse() {
		return nil, response.GetError()
	}

	return response.Data, nil
}
