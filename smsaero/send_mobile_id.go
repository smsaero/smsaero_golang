package smsaero_golang

import (
	"net/url"
	"strconv"
)

// SendMobileId is the response from sending a Mobile ID request.
type SendMobileId struct {
	ID         int    `json:"id"`
	Number     string `json:"number"`
	AuthType   string `json:"authType"`
	CodeSms    string `json:"codeSms"`
	Status     int    `json:"status"`
	Cost       int    `json:"cost"`
	DateCreate int    `json:"dateCreate"`
	DateSend   int    `json:"dateSend"`
}

// SendMobileIdMsg wraps the API response for Mobile ID send.
type SendMobileIdMsg struct {
	Data SendMobileId `json:"data"`
	ErrorResponse
}

// SendMobileId sends a Mobile ID authorization request to the specified phone number.
//
// Parameters:
//   - number (int): The recipient's phone number in international format (e.g., 79031234567)
//   - sign (string): The sender signature for Mobile ID (e.g., "SMSAero")
//   - callbackUrl (string): URL for receiving authorization status callbacks
//
// Returns:
//   - SendMobileId: The server's response containing request details
//   - error: Any error that occurred during the request
//
// Example usage:
//
//	result, err := client.SendMobileId(79031234567, "SMSAero", "https://example.com/callback")
func (c *Client) SendMobileId(number int, sign string, callbackUrl string) (SendMobileId, error) {
	empty := SendMobileId{}

	// Validate input parameters
	phoneStr := strconv.Itoa(number)
	if err := c.validator.ValidatePhone(phoneStr); err != nil {
		return empty, err
	}

	// Validate signature
	if err := c.validator.ValidateSign(sign); err != nil {
		return empty, err
	}

	// Validate callback URL
	if err := c.validator.ValidateCallbackURL(callbackUrl); err != nil {
		return empty, err
	}

	response := new(SendMobileIdMsg)

	data := url.Values{}
	data.Set("number", phoneStr)
	data.Set("sign", sign)
	data.Set("callbackUrl", callbackUrl)

	if err := c.executeRequest("mobile-id/send", response, data); err != nil {
		return empty, err
	}

	if response.IsErrorResponse() {
		return empty, response.GetError()
	}

	return response.Data, nil
}
