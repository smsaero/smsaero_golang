package smsaero_golang

import (
	"net/url"
	"strconv"
)

type HlrStatusMsg struct {
	Data HlrCheck `json:"data"`
	ErrorResponse
}

// HlrStatus retrieves the status of a specific HLR check.
//
// This method returns the current status of an HLR check that was previously
// initiated using the HlrCheck method. It provides detailed information about
// the phone number's availability and status.
//
// Parameters:
//   - id (int): The ID of the HLR check (returned by HlrCheck method)
//
// Returns:
//   - HlrCheck: The response data containing HLR status details
//   - error: error if the request fails or API returns an error
//
// Example response:
//
//	{
//	  "success": true,
//	  "data": {
//	    "id": 12345,
//	    "number": "79031234567",
//	    "hlrStatus": 1,
//	    "extendHlrStatus": "available"
//	  }
//	}
//
// Example usage:
//
//	status, err := client.HlrStatus(12345)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Printf("HLR Status: %s\n", status.ExtendHlrStatus)
func (c *Client) HlrStatus(hlrStatusID int) (HlrCheck, error) {
	empty := HlrCheck{}

	// Validate HLR check ID
	if err := c.validator.ValidateHlrID(hlrStatusID); err != nil {
		return empty, err
	}

	response := new(HlrStatusMsg)

	data := url.Values{}
	data.Set("id", strconv.Itoa(hlrStatusID))

	if err := c.executeRequest("hlr/status", response, data); err != nil {
		return empty, err
	}

	if response.IsErrorResponse() {
		return empty, response.GetError()
	}

	return response.Data, nil
}
