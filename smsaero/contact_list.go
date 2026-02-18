package smsaero_golang

type ContactListMsg struct {
	Data interface{} `json:"data"`
	ErrorResponse
}

// ContactList retrieves a list of contacts from the user's account.
//
// This method returns all contacts associated with the user's account.
// The response includes contact details such as phone numbers, names, and other information.
//
// Returns:
//   - interface{}: The response data containing the list of contacts
//   - error: error if the request fails or API returns an error
//
// Example response:
//
//	{
//	  "success": true,
//	  "data": {
//	    "0": {
//	      "id": 12345,
//	      "number": 79031234567,
//	      "sex": "male",
//	      "lname": "Doe",
//	      "fname": "John",
//	      "sname": "Smith",
//	      "param1": "",
//	      "param2": "",
//	      "param3": "",
//	      "operator": 5,
//	      "extendOperator": "BEELINE",
//	      "groups": [
//	        {
//	          "id": 12345,
//	          "name": "test_group"
//	        }
//	      ],
//	      "hlrStatus": 1,
//	      "extendHlrStatus": "available"
//	    },
//	    "links": {
//	      "self": "/v2/contact/list?page=1",
//	      "first": "/v2/contact/list?page=1",
//	      "last": "/v2/contact/list?page=1"
//	    },
//	    "totalCount": "5"
//	  }
//	}
//
// Example usage:
//
//	contacts, err := client.ContactList()
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Printf("Total contacts: %+v\n", contacts)
func (c *Client) ContactList() (interface{}, error) {
	response := new(ContactListMsg)

	if err := c.executeRequest(`contact/list`, response, nil); err != nil {
		return nil, err
	}

	if response.IsErrorResponse() {
		return nil, response.GetError()
	}

	return response.Data, nil
}
