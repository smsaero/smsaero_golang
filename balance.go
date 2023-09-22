package smsaero_golang

type Balance struct {
	Balance float64 `json:"balance,float64,omitempty"`
}

type BalanceMsg struct {
	Data Balance
	ErrorResponse
}

func (c *Client) Balance() (float64, error) {
	response := new(BalanceMsg)

	if err := c.executeRequest(`balance`, response, nil); err != nil {
		return 0, err
	}

	if response.IsErrorResponse() {
		return 0, response.GetError()
	}

	return response.Data.Balance, nil
}
