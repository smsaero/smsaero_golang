package smsaero_golang

import (
	"net/url"
	"strconv"
)

type GroupDeleteMsg struct {
	Data interface{} `json:"data"`
	ErrorResponse
}

// GroupDelete deletes a group from the SmsAero service.
//
// This method removes a contact group by its ID. The group will be permanently
// deleted and cannot be recovered. All contacts in the group will remain but
// will no longer be associated with the deleted group.
//
// Parameters:
//   - id (int): The ID of the group to be deleted
//
// Returns:
//   - bool: true if deletion was successful, false if not found or error occurred
//   - error: error if the request fails or API returns an error
//
// Example response:
//
//	{
//	  "success": true,
//	  "data": {}
//	}
//
// Example usage:
//
//	success, err := client.GroupDelete(12345)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	if success {
//	    fmt.Println("Group deleted successfully")
//	} else {
//	    fmt.Println("Failed to delete group")
//	}
func (c *Client) GroupDelete(id int) (bool, error) {
	response := new(GroupDeleteMsg)

	data := url.Values{}
	data.Set("id", strconv.Itoa(id))

	if err := c.executeRequest(`group/delete`, response, data); err != nil {
		return false, err
	}

	if response.IsErrorResponse() {
		return false, response.GetError()
	}

	return true, nil
}
