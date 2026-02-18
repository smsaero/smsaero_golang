package smsaero_golang

import (
	"net/url"
	"strconv"
)

type BlackListDeleteMsg struct {
	Data interface{} `json:"data"`
	ErrorResponse
}

// BlackListDelete removes a phone number from the blacklist.
//
// This method removes a phone number from the blacklist by its ID, allowing
// SMS messages to be sent to that number again.
//
// Parameters:
//   - id (int): The ID of the blacklist entry to remove
//
// Returns:
//   - bool: true if successfully removed, false if not found or error occurred
//   - error: error if the request fails or API returns an error
//
// Example response:
//
//	{
//	  "success": true,
//	  "data": {}
//	}
//
// Usage:
//
//	success, err := client.BlackListDelete(12345)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	if success {
//	    fmt.Println("Successfully removed from blacklist")
//	} else {
//	    fmt.Println("Failed to remove from blacklist")
//	}
func (c *Client) BlackListDelete(id int) (bool, error) {
	response := new(BlackListDeleteMsg)

	data := url.Values{}
	data.Set("id", strconv.Itoa(id))

	if err := c.executeRequest(`blacklist/delete`, response, data); err != nil {
		return false, err
	}

	if response.IsErrorResponse() {
		return false, response.GetError()
	}

	return true, nil
}
