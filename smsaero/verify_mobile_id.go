package smsaero_golang

import (
	"net/url"
	"strconv"
)

// VerifyMobileId is the response from verifying a Mobile ID request.
type VerifyMobileId struct {
	ID         int    `json:"id"`
	Number     string `json:"number"`
	AuthType   string `json:"authType"`
	CodeSms    string `json:"codeSms"`
	Status     int    `json:"status"`
	Cost       int    `json:"cost"`
	DateCreate int    `json:"dateCreate"`
	DateSend   int    `json:"dateSend"`
}

// VerifyMobileIdMsg wraps the API response for Mobile ID verification.
type VerifyMobileIdMsg struct {
	Data VerifyMobileId `json:"data"`
	ErrorResponse
}

// VerifyMobileId verifies a Mobile ID request with the provided code.
//
// Parameters:
//   - reqID (int): The unique identifier of the Mobile ID request (obtained from SendMobileId response)
//   - code (string): The verification code received by the user
//   - sign (string): The sender signature used in the original SendMobileId request
//
// Returns:
//   - VerifyMobileId: The verification result with status
//   - error: Any error that occurred during the request
//
// Example usage:
//
//	result, err := client.VerifyMobileId(273, "1234", "SMSAero")
func (c *Client) VerifyMobileId(reqID int, code string, sign string) (VerifyMobileId, error) {
	empty := VerifyMobileId{}

	// Validate input parameters
	if reqID <= 0 {
		return empty, NewValidationError("mobile_id_request_id", "Mobile ID request ID must be a positive number")
	}

	if code == "" {
		return empty, NewValidationError("code", "verification code cannot be empty")
	}

	// Validate signature
	if err := c.validator.ValidateSign(sign); err != nil {
		return empty, err
	}

	response := new(VerifyMobileIdMsg)

	data := url.Values{}
	data.Set("id", strconv.Itoa(reqID))
	data.Set("code", code)
	data.Set("sign", sign)

	if err := c.executeRequest("mobile-id/verify", response, data); err != nil {
		return empty, err
	}

	if response.IsErrorResponse() {
		return empty, response.GetError()
	}

	return response.Data, nil
}
