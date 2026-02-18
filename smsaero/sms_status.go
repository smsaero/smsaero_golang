package smsaero_golang

import (
	"net/url"
	"strconv"
)

type SmsStatus struct {
	ID           int    `json:"id"`
	From         string `json:"from"`
	Number       int    `json:"number"`
	Text         string `json:"text"`
	Status       int    `json:"status"`
	ExtendStatus string `json:"extendStatus"`
	Channel      string `json:"channel"`
	Cost         string `json:"cost"`
	DateCreate   int    `json:"dateCreate"`
	DateSend     int    `json:"dateSend"`
	DateAnswer   int    `json:"dateAnswer"`
}

type SmsStatusMsg struct {
	Data SmsStatus `json:"data"`
	ErrorResponse
}

// SmsStatus retrieves the status of a previously sent SMS message.
//
// Parameters:
//   - id (int): The unique identifier of the SMS message (obtained from SendSms response)
//
// Returns:
//   - SmsStatus: The current status and details of the SMS message
//   - error: Any error that occurred during the request
//
// Example response:
//
//	{
//	  "id": 12345,
//	  "from": "SMS Aero",
//	  "number": 79031234567,
//	  "text": "Hello, World!",
//	  "status": 1,
//	  "extendStatus": "delivered",
//	  "channel": "FREE SIGN",
//	  "cost": "5.49",
//	  "dateCreate": 1719119523,
//	  "dateSend": 1719119523,
//	  "dateAnswer": 1719119525
//	}
//
// Status codes:
//   - 0: Message in queue
//   - 1: Message delivered
//   - 2: Message failed
//   - 3: Message rejected
//
// Example usage:
//
//	status, err := client.SmsStatus(12345)
func (c *Client) SmsStatus(smsStatusID int) (SmsStatus, error) {
	response := new(SmsStatusMsg)
	empty := SmsStatus{}

	// Validate ID
	if err := c.validator.ValidateSmsID(smsStatusID); err != nil {
		return empty, err
	}

	data := url.Values{}
	data.Set("id", strconv.Itoa(smsStatusID))

	path := c.getAPIPath("sms/status", "sms/teststatus")

	if err := c.executeRequest(path, response, data); err != nil {
		return empty, err
	}

	if response.IsErrorResponse() {
		return empty, response.GetError()
	}

	return response.Data, nil
}
