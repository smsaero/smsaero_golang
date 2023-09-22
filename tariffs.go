package smsaero_golang

type TariffsMsg struct {
	Data interface{}
	ErrorResponse
}

func (c *Client) Tariffs() (interface{}, error) {
	response := new(TariffsMsg)

	if err := c.executeRequest(`tariffs`, response, nil); err != nil {
		return nil, err
	}

	if response.IsErrorResponse() {
		return nil, response.GetError()
	}

	return response.Data, nil
}
