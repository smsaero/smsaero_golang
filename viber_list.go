package smsaero_golang

type ViberListMsg struct {
	Data interface{}
	ErrorResponse
}

func (c *Client) ViberList() (interface{}, error) {
	response := new(ViberListMsg)

	if err := c.executeRequest(`viber/list`, response, nil); err != nil {
		return nil, err
	}

	if response.IsErrorResponse() {
		return nil, response.GetError()
	}

	return response.Data, nil
}
