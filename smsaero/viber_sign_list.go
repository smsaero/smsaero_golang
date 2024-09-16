package smsaero_golang

type ViberSignListMsg struct {
	Data interface{}
	ErrorResponse
}

func (c *Client) ViberSignList() (interface{}, error) {
	response := new(ViberSignListMsg)

	if err := c.executeRequest(`viber/sign/list`, response, nil); err != nil {
		return nil, err
	}

	if response.IsErrorResponse() {
		return nil, response.GetError()
	}

	return response.Data, nil
}
