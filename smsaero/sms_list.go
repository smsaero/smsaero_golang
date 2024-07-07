package smsaero_golang

type SmsListMsg struct {
	Data interface{}
	ErrorResponse
}

func (c *Client) SmsList() (interface{}, error) {
	response := new(SmsListMsg)

	path := c.getApiPath("sms/list", "sms/testlist")

	if err := c.executeRequest(path, response, nil); err != nil {
		return nil, err
	}

	if response.IsErrorResponse() {
		return nil, response.GetError()
	}

	return response.Data, nil
}
