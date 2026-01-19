package smsaero_golang

type SignListMsg struct {
	Data interface{} `json:"data"`
	ErrorResponse
}

// SignList retrieves a list of signatures for the user's account.
//
// This method returns all available signatures that can be used for sending
// SMS messages. Each signature has a status indicating whether it's approved
// and available for use with different operators.
//
// Returns:
//   - interface{}: The response data containing the list of signatures
//   - error: error if the request fails or API returns an error
//
// Example response:
//
//	{
//	  "success": true,
//	  "data": {
//	    "totalCount": "1",
//	    "0": {
//	      "id": 12345,
//	      "name": "TestSign",
//	      "status": 1,
//	      "extendStatus": "approve",
//	      "statusOperators": {
//	        "1": {
//	          "operator": 1,
//	          "extendOperator": "MEGAFON",
//	          "status": 1,
//	          "extendStatus": "approve"
//	        },
//	        "4": {
//	          "operator": 4,
//	          "extendOperator": "MTS",
//	          "status": 1,
//	          "extendStatus": "approve"
//	        }
//	      }
//	    }
//	  }
//	}
//
// Example usage:
//
//	signs, err := client.SignList()
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Printf("Available signatures: %+v\n", signs)
func (c *Client) SignList() (interface{}, error) {
	response := new(SignListMsg)

	if err := c.executeRequest(`sign/list`, response, nil); err != nil {
		return nil, err
	}

	if response.IsErrorResponse() {
		return nil, response.GetError()
	}

	return response.Data, nil
}
