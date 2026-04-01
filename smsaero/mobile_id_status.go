package smsaero_golang

import (
	"net/url"
	"strconv"
)

// MobileIdStatus holds the status of a previously sent Mobile ID request.
type MobileIdStatus struct {
	ID         int    `json:"id"`
	Number     string `json:"number"`
	AuthType   string `json:"authType"`
	CodeSms    string `json:"codeSms"`
	Status     int    `json:"status"`
	Cost       int    `json:"cost"`
	DateCreate int    `json:"dateCreate"`
	DateSend   int    `json:"dateSend"`
}

// MobileIdStatusMsg wraps the API response for Mobile ID status.
type MobileIdStatusMsg struct {
	Data MobileIdStatus `json:"data"`
	ErrorResponse
}

// MobileIdStatus retrieves the status of a Mobile ID request.
//
// Parameters:
//   - reqID (int): The unique identifier of the Mobile ID request (obtained from SendMobileId response)
//
// Returns:
//   - MobileIdStatus: The current status and details of the Mobile ID request
//   - error: Any error that occurred during the request
//
// Example usage:
//
//	status, err := client.MobileIdStatus(273)
func (c *Client) MobileIdStatus(reqID int) (MobileIdStatus, error) {
	empty := MobileIdStatus{}

	if reqID <= 0 {
		return empty, NewValidationError("mobile_id_request_id", "Mobile ID request ID must be a positive number")
	}

	response := new(MobileIdStatusMsg)

	data := url.Values{}
	data.Set("id", strconv.Itoa(reqID))

	if err := c.executeRequest("mobile-id/status", response, data); err != nil {
		return empty, err
	}

	if response.IsErrorResponse() {
		return empty, response.GetError()
	}

	return response.Data, nil
}
