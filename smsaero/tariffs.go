package smsaero_golang

type TariffsMsg struct {
	Data interface{} `json:"data"`
	ErrorResponse
}

// Tariffs retrieves the tariffs for the user's account.
//
// This method returns pricing information for SMS messages based on different
// signatures and mobile operators. The tariffs show the cost per SMS for
// each operator and signature combination.
//
// Returns:
//   - interface{}: The response data containing tariff information
//   - error: error if the request fails or API returns an error
//
// Example response:
//
//	{
//	  "success": true,
//	  "data": {
//	    "FREE SIGN": {
//	      "MEGAFON": "8.99",
//	      "MTS": "4.99",
//	      "BEELINE": "5.49",
//	      "TELE2": "4.79",
//	      "OTHER": "5.19"
//	    }
//	  }
//	}
//
// Example usage:
//
//	tariffs, err := client.Tariffs()
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Printf("Available tariffs: %+v\n", tariffs)
func (c *Client) Tariffs() (interface{}, error) {
	response := new(TariffsMsg)

	if err := c.executeRequest(`tariffs`, response, nil); err != nil {
		return nil, err
	}

	if response.IsErrorResponse() {
		return nil, response.GetError()
	}

	return response.Data, nil
}
