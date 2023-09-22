package smsaero_golang

type FlashCallListMsg struct {
	Data interface{}
	ErrorResponse
}

func (c *Client) FlashCallList() (interface{}, error) {
	response := new(FlashCallListMsg)

	if err := c.executeRequest(`flashcall/list`, response, nil); err != nil {
		return nil, err
	}

	if response.IsErrorResponse() {
		return nil, response.GetError()
	}

	return response.Data, nil
}
