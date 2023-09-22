package smsaero_golang

type SignListMsg struct {
	Data interface{}
	ErrorResponse
}

func (c *Client) SignList() (interface{}, error) {
	response := new(SignListMsg)

	if err := c.executeRequest(`sign/list`, response, nil); err != nil {
		return nil, err
	}

	if response.IsErrorResponse() {
		return nil, response.GetError()
	}

	return response.Data, nil
}
