package smsaero_golang

type CardsMsg struct {
	Data interface{}
	ErrorResponse
}

func (c *Client) Cards() (interface{}, error) {
	response := new(CardsMsg)

	if err := c.executeRequest(`cards`, response, nil); err != nil {
		return nil, err
	}

	if response.IsErrorResponse() {
		return nil, response.GetError()
	}

	return response.Data, nil
}
