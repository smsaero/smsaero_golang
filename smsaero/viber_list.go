package smsaero_golang

type ViberListMsg struct {
	Data interface{} `json:"data"`
	ErrorResponse
}

// ViberList retrieves a list of Viber messages.
//
// This method returns all Viber messages sent from the user's account.
// The response includes message details such as status, delivery statistics,
// and timestamps.
//
// Returns:
//   - interface{}: The response data containing the list of Viber messages
//   - error: error if the request fails or API returns an error
//
// Example response:
//
//	{
//	  "success": true,
//	  "data": {
//	    "0": {
//	      "id": 1,
//	      "number": "79031234567",
//	      "count": 1,
//	      "sign": "Viber",
//	      "channel": "OFFICIAL",
//	      "text": "Hello, World!",
//	      "cost": 2.25,
//	      "status": 1,
//	      "extendStatus": "moderation",
//	      "dateCreate": 1511153253,
//	      "dateSend": 1511153253,
//	      "countSend": 0,
//	      "countDelivered": 0,
//	      "countWrite": 0,
//	      "countUndelivered": 0,
//	      "countError": 0
//	    },
//	    "links": {
//	      "self": "/v2/viber/list?page=1",
//	      "next": "/v2/viber/list?page=2",
//	      "last": "/v2/viber/list?page=3"
//	    }
//	  }
//	}
//
// Example usage:
//
//	messages, err := client.ViberList()
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Printf("Total Viber messages: %+v\n", messages)
func (c *Client) ViberList() (interface{}, error) {
	response := new(ViberListMsg)

	if err := c.executeRequest(`viber/list`, response, nil); err != nil {
		return nil, err
	}

	if response.IsErrorResponse() {
		return nil, response.GetError()
	}

	return response.Data, nil
}
