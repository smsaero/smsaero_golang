package smsaero_golang

import (
	"net/url"
	"strconv"
)

type HlrCheck struct {
	ID              int    `json:"id"`
	Number          string `json:"number"`
	HlrStatus       int    `json:"hlrStatus"`
	ExtendHlrStatus string `json:"extendHlrStatus"`
}

type HlrCheckMsg struct {
	Data HlrCheck `json:"data"`
	ErrorResponse
}

// HlrCheck checks the Home Location Register (HLR) for a phone number.
//
// This method performs an HLR lookup to determine the status and availability
// of a phone number. HLR checks are used to verify if a number is active
// and can receive SMS messages before sending.
//
// Parameters:
//   - number (int): The phone number to check in international format (e.g., 79031234567)
//
// Returns:
//   - HlrCheck: The response data containing HLR check details
//   - error: error if the request fails or API returns an error
//
// Example response:
//
//	{
//	  "success": true,
//	  "data": {
//	    "id": 12345,
//	    "number": "79031234567",
//	    "hlrStatus": 4,
//	    "extendHlrStatus": "in work"
//	  }
//	}
//
// Example usage:
//
//	result, err := client.HlrCheck(79031234567)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Printf("HLR Status: %s\n", result.ExtendHlrStatus)
func (c *Client) HlrCheck(number int) (HlrCheck, error) {
	empty := HlrCheck{}

	// Validate phone number
	phoneStr := strconv.Itoa(number)
	if err := c.validator.ValidatePhone(phoneStr); err != nil {
		return empty, err
	}

	response := new(HlrCheckMsg)

	data := url.Values{}
	data.Set("number", phoneStr)

	if err := c.executeRequest("hlr/check", response, data); err != nil {
		return empty, err
	}

	if response.IsErrorResponse() {
		return empty, response.GetError()
	}

	return response.Data, nil
}
