package smsaero_golang

type BlackListMsg struct {
	Data interface{} `json:"data"`
	ErrorResponse
}

// BlackList retrieves the list of phone numbers in the blacklist.
//
// This method returns all phone numbers that are currently in the blacklist,
// preventing them from receiving SMS messages.
//
// Returns:
//   - interface{}: The response data containing the list of blacklisted numbers
//   - error: error if the request fails or API returns an error
//
// Example response:
//
//	{
//	  "success": true,
//	  "data": [
//	    {
//	      "id": 12345,
//	      "number": "79031234567",
//	      "fname": "John",
//	      "lname": "Doe"
//	    },
//	    {
//	      "id": 12346,
//	      "number": "79031234568",
//	      "fname": "Jane",
//	      "lname": "Smith"
//	    }
//	  ]
//	}
//
// Usage:
//
//	result, err := client.BlackList()
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Printf("Blacklist contains %d entries\n", len(result.([]interface{})))
func (c *Client) BlackList() (interface{}, error) {
	response := new(BlackListMsg)

	if err := c.executeRequest(`blacklist/list`, response, nil); err != nil {
		return nil, err
	}

	if response.IsErrorResponse() {
		return nil, response.GetError()
	}

	return response.Data, nil
}
