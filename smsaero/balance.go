package smsaero_golang

type Balance struct {
	Balance float64 `json:"balance,omitempty"`
}

type BalanceMsg struct {
	Data Balance `json:"data"`
	ErrorResponse
}

// Balance retrieves the current account balance.
//
// Returns:
//   - float64: The current account balance in rubles
//   - error: Any error that occurred during the request
//
// Example response:
//
//	{
//	  "balance": 150.75
//	}
//
// Example usage:
//
//	balance, err := client.Balance()
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Printf("Current balance: %.2f rubles\n", balance)
func (c *Client) Balance() (float64, error) {
	response := new(BalanceMsg)

	if err := c.executeRequest("balance", response, nil); err != nil {
		return 0, err
	}

	if response.IsErrorResponse() {
		return 0, response.GetError()
	}

	return response.Data.Balance, nil
}
