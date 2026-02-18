package smsaero_golang

import (
	"net/url"
	"strconv"
	"time"
)

// SendSmsOption represents an option for sending SMS
type SendSmsOption func(*sendSmsOptions)

type sendSmsOptions struct {
	sign           string
	dateToSend     *time.Time
	callbackURL    string
	callbackFormat string
	shortLink      bool
}

// WithSendSmsSign sets the sender signature for SMS
func WithSendSmsSign(sign string) SendSmsOption {
	return func(opts *sendSmsOptions) {
		opts.sign = sign
	}
}

// WithSendSmsDateToSend sets the scheduled date for delayed SMS sending
func WithSendSmsDateToSend(dateToSend time.Time) SendSmsOption {
	return func(opts *sendSmsOptions) {
		opts.dateToSend = &dateToSend
	}
}

// WithSendSmsCallbackURL sets the callback URL for status notifications
func WithSendSmsCallbackURL(callbackURL string) SendSmsOption {
	return func(opts *sendSmsOptions) {
		opts.callbackURL = callbackURL
	}
}

// WithSendSmsCallbackFormat sets the callback format (JSON or x-www-form-urlencoded)
// When set to "JSON", status updates will be sent in JSON format
func WithSendSmsCallbackFormat(format string) SendSmsOption {
	return func(opts *sendSmsOptions) {
		opts.callbackFormat = format
	}
}

// WithSendSmsShortLink enables automatic URL shortening in the message
// When enabled (true), all URLs in the message will be automatically shortened
func WithSendSmsShortLink(enabled bool) SendSmsOption {
	return func(opts *sendSmsOptions) {
		opts.shortLink = enabled
	}
}

type SendSms struct {
	ID           int     `json:"id"`
	From         string  `json:"from"`
	Number       string  `json:"number"`
	Text         string  `json:"text"`
	Status       int     `json:"status"`
	ExtendStatus string  `json:"extendStatus"`
	Channel      string  `json:"channel"`
	Cost         float64 `json:"cost"`
	DateCreate   int     `json:"dateCreate"`
	DateSend     int     `json:"dateSend"`
}

type SendMsg struct {
	Data SendSms `json:"data"`
	ErrorResponse
}

// SendSms sends an SMS message to the specified phone number.
//
// Parameters:
//   - number (int): The recipient's phone number in international format (e.g., 79031234567)
//   - text (string): The text content of the message (1-800 characters)
//   - opts (...SendSmsOption): Optional parameters for customizing the SMS
//
// Options:
//   - WithSendSmsSign(sign): Set custom sender signature (1-11 characters)
//   - WithSendSmsDateToSend(date): Schedule message for future delivery
//   - WithSendSmsCallbackURL(url): Set callback URL for status notifications
//   - WithSendSmsCallbackFormat(format): Set callback format ("JSON" for JSON format)
//   - WithSendSmsShortLink(enabled): Enable automatic URL shortening
//
// Returns:
//   - SendSms: The server's response containing message details
//   - error: Any error that occurred during the request
//
// Example response:
//
//	{
//	  "id": 12345,
//	  "from": "SMS Aero",
//	  "number": "79031234567",
//	  "text": "Hello, World!",
//	  "status": 0,
//	  "extendStatus": "queue",
//	  "channel": "FREE SIGN",
//	  "cost": 5.49,
//	  "dateCreate": 1719119523,
//	  "dateSend": 1719119523
//	}
//
// Example usage:
//
//	result, err := client.SendSms(79031234567, "Hello!",
//	    WithSendSmsSign("MyCompany"),
//	    WithSendSmsCallbackURL("https://example.com/callback"))
func (c *Client) SendSms(number int, text string, opts ...SendSmsOption) (SendSms, error) {
	empty := SendSms{}

	// Validate input parameters
	phoneStr := strconv.Itoa(number)
	if err := c.validator.ValidatePhone(phoneStr); err != nil {
		return empty, err
	}

	if err := c.validator.ValidateMessage(text); err != nil {
		return empty, err
	}

	// Process options
	options := &sendSmsOptions{
		sign: c.sign,
	}

	for _, opt := range opts {
		opt(options)
	}

	// Validate signature
	if err := c.validator.ValidateSign(options.sign); err != nil {
		return empty, err
	}

	// Validate callback URL
	if err := c.validator.ValidateCallbackURL(options.callbackURL); err != nil {
		return empty, err
	}

	// Validate send date
	if err := c.validator.ValidateDateToSend(options.dateToSend); err != nil {
		return empty, err
	}

	response := new(SendMsg)

	data := url.Values{}
	data.Set("number", phoneStr)
	data.Set("sign", options.sign)
	data.Set("text", text)

	// Add callback URL if specified
	if options.callbackURL != "" {
		data.Set("callbackUrl", options.callbackURL)
	}

	// Add callback format if specified
	if options.callbackFormat != "" {
		data.Set("callbackFormat", options.callbackFormat)
	}

	// Add short link option if enabled
	if options.shortLink {
		data.Set("shortLink", "1")
	}

	// Add send date if specified
	if options.dateToSend != nil {
		data.Set("dateSend", strconv.FormatInt(options.dateToSend.Unix(), 10))
	}

	path := c.getAPIPath("sms/send", "sms/testsend")

	if err := c.executeRequest(path, response, data); err != nil {
		return empty, err
	}

	if response.IsErrorResponse() {
		return empty, response.GetError()
	}

	return response.Data, nil
}
