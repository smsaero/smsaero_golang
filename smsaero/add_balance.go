package smsaero_golang

import (
	"fmt"
	"net/url"
	"strconv"
)

type AddBalanceMsg struct {
	Data interface{} `json:"data"`
	ErrorResponse
}

// AddBalance adds funds to the account balance using a registered card.
//
// Parameters:
//   - amount (float64): The amount to add to the balance (minimum 1.00 rubles)
//   - cardID (int): The ID of the registered payment card
//
// Returns:
//   - interface{}: The server's response containing transaction details
//   - error: Any error that occurred during the request
//
// Example response:
//
//	{
//	  "success": true,
//	  "transactionId": "tx_123456789",
//	  "amount": 100.00,
//	  "newBalance": 250.75
//	}
//
// Example usage:
//
//	result, err := client.AddBalance(100.00, 12345)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Printf("Balance added successfully: %v\n", result)
func (c *Client) AddBalance(amount float64, cardID int) (interface{}, error) {
	if err := c.validator.ValidateAmount(amount); err != nil {
		return nil, err
	}

	if err := c.validator.ValidateCardID(cardID); err != nil {
		return nil, err
	}

	response := new(AddBalanceMsg)

	data := url.Values{}
	data.Set("sum", fmt.Sprintf("%.2f", amount))
	data.Set("cardId", strconv.Itoa(cardID))

	if err := c.executeRequest("balance/add", response, data); err != nil {
		return nil, err
	}

	if response.IsErrorResponse() {
		return nil, response.GetError()
	}

	return response.Data, nil
}
