package smsaero_golang

type BlackListMsg struct {
	Data interface{}
	ErrorResponse
}

func (c *Client) BlackList() (interface{}, error) {
	response := new(BlackListMsg)

	if err := c.executeRequest(`blacklist/list`, response, nil); err != nil {
		return nil, err
	}

	if response.IsErrorResponse() {
		return nil, response.GetError()
	}

	return response.Data, nil
}
