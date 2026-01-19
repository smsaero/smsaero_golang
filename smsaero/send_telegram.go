package smsaero_golang

import (
	"net/url"
	"strconv"
)

// SendTelegramOption represents an option for sending Telegram code
type SendTelegramOption func(*sendTelegramOptions)

type sendTelegramOptions struct {
	sign string
	text string
}

// WithSendTelegramSign sets the sender signature for Telegram message
func WithSendTelegramSign(sign string) SendTelegramOption {
	return func(opts *sendTelegramOptions) {
		opts.sign = sign
	}
}

// WithSendTelegramText sets custom text for Telegram message
func WithSendTelegramText(text string) SendTelegramOption {
	return func(opts *sendTelegramOptions) {
		opts.text = text
	}
}

type SendTelegram struct {
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

type SendTelegramMsg struct {
	Data SendTelegram `json:"data"`
	ErrorResponse
}

// SendTelegram sends a Telegram code to the specified phone number.
//
// Parameters:
//   - number (int): The recipient's phone number in international format (e.g., 79031234567)
//   - code (int): The verification code to send (4-8 digits)
//   - opts (...SendTelegramOption): Optional parameters for customizing the message
//
// Options:
//   - WithSendTelegramSign(sign): Set custom sender signature (1-11 characters)
//   - WithSendTelegramText(text): Set custom message text (minimum 10 characters)
//
// Returns:
//   - SendTelegram: The server's response containing message details
//   - error: Any error that occurred during the request
//
// Example response:
//
//	{
//	  "id": 12345,
//	  "number": "79031234567",
//	  "telegramCode": "1234",
//	  "status": 0,
//	  "extendStatus": "queue",
//	  "channel": "FREE SIGN",
//	  "cost": 5.49,
//	  "dateCreate": 1719119523
//	}
//
// Example usage:
//
//	result, err := client.SendTelegram(79031234567, 1234,
//	    WithSendTelegramSign("MyCompany"),
//	    WithSendTelegramText("Your verification code: 1234"))
func (c *Client) SendTelegram(number int, code int, opts ...SendTelegramOption) (SendTelegram, error) {
	empty := SendTelegram{}

	// Validate input parameters
	phoneStr := strconv.Itoa(number)
	if err := c.validator.ValidatePhone(phoneStr); err != nil {
		return empty, err
	}

	if err := c.validator.ValidateTelegramCode(code); err != nil {
		return empty, err
	}

	// Process options
	options := &sendTelegramOptions{}
	for _, opt := range opts {
		opt(options)
	}

	// Validate signature if specified
	if options.sign != "" {
		if err := c.validator.ValidateSign(options.sign); err != nil {
			return empty, err
		}
	}

	// Validate text if specified
	if options.text != "" {
		if err := c.validator.ValidateMessage(options.text); err != nil {
			return empty, err
		}
	}

	response := new(SendTelegramMsg)

	data := url.Values{}
	data.Set("number", phoneStr)
	data.Set("code", strconv.Itoa(code))

	// Optional parameters
	if options.sign != "" {
		data.Set("sign", options.sign)
	}
	if options.text != "" {
		data.Set("text", options.text)
	}

	if err := c.executeRequest("telegram/send", response, data); err != nil {
		return empty, err
	}

	if response.IsErrorResponse() {
		return empty, response.GetError()
	}

	return response.Data, nil
}
