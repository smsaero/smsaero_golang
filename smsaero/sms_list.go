package smsaero_golang

type SmsListMsg struct {
	Data interface{} `json:"data"`
	ErrorResponse
}

// SmsList retrieves a list of SMS messages.
//
// This method returns all SMS messages sent from the user's account.
// The response includes message details such as status, delivery information,
// and timestamps. In test mode, it returns test messages.
//
// Returns:
//   - interface{}: The response data containing the list of SMS messages
//   - error: error if the request fails or API returns an error
//
// Example response:
//
//	{
//	  "success": true,
//	  "data": {
//	    "0": {
//	      "id": 12345,
//	      "from": "Sms Aero",
//	      "number": "79031234567",
//	      "text": "Hello, World!",
//	      "status": 1,
//	      "extendStatus": "delivery",
//	      "channel": "FREE SIGN",
//	      "cost": "5.49",
//	      "dateCreate": 1697533302,
//	      "dateSend": 1697533302,
//	      "dateAnswer": 1697533306
//	    },
//	    "links": {
//	      "self": "/v2/sms/list?page=1",
//	      "first": "/v2/sms/list?page=1",
//	      "last": "/v2/sms/list?page=3",
//	      "next": "/v2/sms/list?page=2"
//	    },
//	    "totalCount": "138"
//	  }
//	}
//
// Example usage:
//
//	messages, err := client.SmsList()
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Printf("Total messages: %+v\n", messages)
func (c *Client) SmsList() (interface{}, error) {
	response := new(SmsListMsg)

	path := c.getAPIPath("sms/list", "sms/testlist")

	if err := c.executeRequest(path, response, nil); err != nil {
		return nil, err
	}

	if response.IsErrorResponse() {
		return nil, response.GetError()
	}

	return response.Data, nil
}
