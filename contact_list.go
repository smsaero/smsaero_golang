package smsaero_golang

type ContactListMsg struct {
	Data interface{}
	ErrorResponse
}

func (c *Client) ContactList() (interface{}, error) {
	response := new(ContactListMsg)

	if err := c.executeRequest(`contact/list`, response, nil); err != nil {
		return nil, err
	}

	if response.IsErrorResponse() {
		return nil, response.GetError()
	}

	return response.Data, nil
}
