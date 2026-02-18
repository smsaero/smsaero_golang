package smsaero_golang

import (
	"net/url"
	"strconv"
)

type ContactDeleteMsg struct {
	Data interface{} `json:"data"`
	ErrorResponse
}

// ContactDelete deletes a contact from the user's account.
//
// This method removes a contact from the user's account by its ID.
// The contact will be permanently deleted and cannot be recovered.
//
// Parameters:
//   - id (int): The ID of the contact to be deleted
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
//	success, err := client.ContactDelete(12345)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	if success {
//	    fmt.Println("Contact deleted successfully")
//	} else {
//	    fmt.Println("Failed to delete contact")
//	}
func (c *Client) ContactDelete(id int) (bool, error) {
	response := new(ContactDeleteMsg)

	data := url.Values{}
	data.Set("id", strconv.Itoa(id))

	if err := c.executeRequest(`contact/delete`, response, data); err != nil {
		return false, err
	}

	if response.IsErrorResponse() {
		return false, response.GetError()
	}

	return true, nil
}
