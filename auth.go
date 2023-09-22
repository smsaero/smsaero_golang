package smsaero_golang

type Auth struct {
	Data interface{}
	ErrorResponse
}

func (c *Client) Auth() (bool, error) {
	response := new(Auth)

	if err := c.executeRequest(`auth`, response, nil); err != nil {
		return false, err
	}

	if response.IsErrorResponse() {
		return false, response.GetError()
	}

	return response.Success, nil
}
