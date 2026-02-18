package smsaero_golang

import "net/url"

type ContactDeleteAllMsg struct {
	Data interface{} `json:"data"`
	ErrorResponse
}

// ContactDeleteAll deletes all contacts from the user's account.
//
// This method removes all contacts from the user's account.
// All contacts will be permanently deleted and cannot be recovered.
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
//	  "message": "Contact delete."
//	}
//
// Example usage:
//
//	success, err := client.ContactDeleteAll()
//	if err != nil {
//	    log.Fatal(err)
//	}
//	if success {
//	    fmt.Println("All contacts deleted successfully")
//	}
func (c *Client) ContactDeleteAll() (bool, error) {
	response := new(ContactDeleteAllMsg)

	data := url.Values{}

	if err := c.executeRequest("contact/delete-all", response, data); err != nil {
		return false, err
	}

	if response.IsErrorResponse() {
		return false, response.GetError()
	}

	return true, nil
}
