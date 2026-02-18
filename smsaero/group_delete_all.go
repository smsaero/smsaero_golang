package smsaero_golang

import "net/url"

type GroupDeleteAllMsg struct {
	Data interface{} `json:"data"`
	ErrorResponse
}

// GroupDeleteAll deletes all groups from the SmsAero service.
//
// This method removes all contact groups. All groups will be permanently
// deleted and cannot be recovered. All contacts will remain but
// will no longer be associated with any groups.
//
// Returns:
//   - bool: true if deletion was successful
//   - error: error if the request fails or API returns an error
//
// Example response:
//
//	{
//	  "success": true,
//	  "data": null,
//	  "message": "All groups deleted."
//	}
//
// Example usage:
//
//	success, err := client.GroupDeleteAll()
//	if err != nil {
//	    log.Fatal(err)
//	}
//	if success {
//	    fmt.Println("All groups deleted successfully")
//	}
func (c *Client) GroupDeleteAll() (bool, error) {
	response := new(GroupDeleteAllMsg)

	data := url.Values{}

	if err := c.executeRequest("group/delete-all", response, data); err != nil {
		return false, err
	}

	if response.IsErrorResponse() {
		return false, response.GetError()
	}

	return true, nil
}
