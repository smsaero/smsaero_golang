package smsaero_golang

type GroupListMsg struct {
	Data interface{} `json:"data"`
	ErrorResponse
}

// GroupList retrieves a list of groups from the SmsAero service.
//
// This method returns all contact groups associated with the user's account.
// The response includes group details such as ID, name, and contact count.
//
// Returns:
//   - interface{}: The response data containing the list of groups
//   - error: error if the request fails or API returns an error
//
// Example response:
//
//	{
//	  "success": true,
//	  "data": {
//	    "0": {
//	      "id": "12345",
//	      "name": "test_group",
//	      "countContacts": "0"
//	    },
//	    "links": {
//	      "self": "/v2/group/list?page=1",
//	      "first": "/v2/group/list?page=1",
//	      "last": "/v2/group/list?page=1"
//	    },
//	    "totalCount": "4"
//	  }
//	}
//
// Example usage:
//
//	groups, err := client.GroupList()
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Printf("Total groups: %+v\n", groups)
func (c *Client) GroupList() (interface{}, error) {
	response := new(GroupListMsg)

	if err := c.executeRequest(`group/list`, response, nil); err != nil {
		return nil, err
	}

	if response.IsErrorResponse() {
		return nil, response.GetError()
	}

	return response.Data, nil
}
