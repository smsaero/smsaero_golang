package smsaero_golang

import (
	"net/url"
	"strconv"
)

type TelegramStatus struct {
	ID           int    `json:"id"`
	Number       string `json:"number"`
	TelegramCode string `json:"telegramCode"`
	SmsText      string `json:"smsText"`
	SmsFrom      string `json:"smsFrom"`
	IDSms        *int   `json:"idSms"`
	Status       int    `json:"status"`
	ExtendStatus string `json:"extendStatus"`
	Cost         string `json:"cost"`
	DateCreate   int    `json:"dateCreate"`
}

type TelegramStatusMsg struct {
	Data TelegramStatus `json:"data"`
	ErrorResponse
}

// TelegramStatus retrieves the status of a previously sent Telegram message.
//
// Parameters:
//   - telegramId (int): The unique identifier of the Telegram message (obtained from SendTelegram response)
//
// Returns:
//   - TelegramStatus: The current status and details of the Telegram message
//   - error: Any error that occurred during the request
//
// Example response:
//
//	{
//	  "id": 12345,
//	  "number": "79031234567",
//	  "telegramCode": "1234",
//	  "smsText": "Your verification code: 1234",
//	  "smsFrom": "MyCompany",
//	  "idSms": 67890,
//	  "status": 1,
//	  "extendStatus": "delivered",
//	  "cost": "5.49",
//	  "dateCreate": 1719119523
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
//	status, err := client.TelegramStatus(12345)
func (c *Client) TelegramStatus(telegramID int) (TelegramStatus, error) {
	response := new(TelegramStatusMsg)
	empty := TelegramStatus{}

	// Validate ID
	if telegramID <= 0 {
		return empty, NewValidationError("telegram message ID must be a positive number", "INVALID_TELEGRAM_ID")
	}

	data := url.Values{}
	data.Set("id", strconv.Itoa(telegramID))

	if err := c.executeRequest(`telegram/status`, response, data); err != nil {
		return empty, err
	}

	if response.IsErrorResponse() {
		return empty, response.GetError()
	}

	return response.Data, nil
}
