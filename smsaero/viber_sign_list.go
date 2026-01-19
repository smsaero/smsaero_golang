package smsaero_golang

type ViberSignListMsg struct {
	Data interface{} `json:"data"`
	ErrorResponse
}

// ViberSignList retrieves a list of Viber signatures.
//
// This method returns all available Viber signatures that can be used for sending
// Viber messages. Each signature has a status indicating whether it's approved
// and available for use with different operators.
//
// Returns:
//   - interface{}: The response data containing the list of Viber signatures
//   - error: error if the request fails or API returns an error
//
// Example response:
//
//	{
//	  "success": true,
//	  "data": {
//	    "0": {
//	      "id": 12345,
//	      "name": "GOOD SIGN",
//	      "status": 1,
//	      "extendStatus": "active",
//	      "statusOperators": {
//	        "1": {
//	          "operator": 1,
//	          "extendOperator": "MEGAFON",
//	          "status": 0,
//	          "extendStatus": "moderation"
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
//	signs, err := client.ViberSignList()
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Printf("Available Viber signatures: %+v\n", signs)
func (c *Client) ViberSignList() (interface{}, error) {
	response := new(ViberSignListMsg)

	if err := c.executeRequest(`viber/sign/list`, response, nil); err != nil {
		return nil, err
	}

	if response.IsErrorResponse() {
		return nil, response.GetError()
	}

	return response.Data, nil
}
