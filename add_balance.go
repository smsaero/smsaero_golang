package smsaero_golang

import "net/url"

type AddBalanceMsg struct {
	Data interface{}
	ErrorResponse
}

func (c *Client) AddBalance(sum string, cardId string) (interface{}, error) {
	response := new(AddBalanceMsg)

	data := url.Values{}
	data.Set("sum", sum)
	data.Set("cardId", cardId)

	if err := c.executeRequest(`balance/add`, response, data); err != nil {
		return nil, err
	}

	if response.IsErrorResponse() {
		return nil, response.GetError()
	}

	return response.Data, nil
}
