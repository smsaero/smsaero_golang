package smsaero_golang

type CardsMsg struct {
	Data interface{} `json:"data"`
	ErrorResponse
}

// Cards retrieves the list of available payment cards.
//
// This method returns information about payment cards that can be used
// for adding balance to the account.
//
// Returns:
//   - interface{}: The response data containing the list of available cards
//   - error: error if the request fails or API returns an error
//
// Example response:
//
//	{
//	  "success": true,
//	  "data": [
//	    {
//	      "id": 1,
//	      "name": "Visa",
//	      "min_amount": 100,
//	      "max_amount": 50000
//	    },
//	    {
//	      "id": 2,
//	      "name": "MasterCard",
//	      "min_amount": 100,
//	      "max_amount": 50000
//	    }
//	  ]
//	}
//
// Usage:
//
//	result, err := client.Cards()
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Printf("Available cards: %+v\n", result)
func (c *Client) Cards() (interface{}, error) {
	response := new(CardsMsg)

	if err := c.executeRequest(`cards`, response, nil); err != nil {
		return nil, err
	}

	if response.IsErrorResponse() {
		return nil, response.GetError()
	}

	return response.Data, nil
}
