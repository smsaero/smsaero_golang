package smsaero_golang

type GroupListMsg struct {
	Data interface{}
	ErrorResponse
}

func (c *Client) GroupList() (interface{}, error) {
	response := new(GroupListMsg)

	if err := c.executeRequest(`group/list`, response, nil); err != nil {
		return nil, err
	}

	if response.IsErrorResponse() {
		return nil, response.GetError()
	}

	return response.Data, nil
}
