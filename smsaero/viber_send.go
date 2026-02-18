package smsaero_golang

import (
	"net/url"
	"strconv"
)

type ViberSendMsg struct {
	Data interface{} `json:"data"`
	ErrorResponse
}

// ViberSendOption represents option for sending Viber message
type ViberSendOption func(*viberSendOptions)

type viberSendOptions struct {
	number      string
	groupID     int
	imageSource string
	textButton  string
	linkButton  string
	dateSend    string
	signSms     string
	channelSms  string
	textSms     string
	priceSms    int
	timeout     float64
}

// WithViberSendNumber sets phone number
func WithViberSendNumber(number string) ViberSendOption {
	return func(opts *viberSendOptions) {
		opts.number = number
	}
}

// WithViberSendGroupID sets group ID
func WithViberSendGroupID(groupID int) ViberSendOption {
	return func(opts *viberSendOptions) {
		opts.groupID = groupID
	}
}

// WithViberSendImageSource sets image source
func WithViberSendImageSource(imageSource string) ViberSendOption {
	return func(opts *viberSendOptions) {
		opts.imageSource = imageSource
	}
}

// WithViberSendTextButton sets button text
func WithViberSendTextButton(textButton string) ViberSendOption {
	return func(opts *viberSendOptions) {
		opts.textButton = textButton
	}
}

// WithViberSendLinkButton sets button link
func WithViberSendLinkButton(linkButton string) ViberSendOption {
	return func(opts *viberSendOptions) {
		opts.linkButton = linkButton
	}
}

// WithViberSendDateSend sets send date
func WithViberSendDateSend(dateSend string) ViberSendOption {
	return func(opts *viberSendOptions) {
		opts.dateSend = dateSend
	}
}

// WithViberSendSignSms sets SMS signature
func WithViberSendSignSms(signSms string) ViberSendOption {
	return func(opts *viberSendOptions) {
		opts.signSms = signSms
	}
}

// WithViberSendChannelSms sets SMS channel
func WithViberSendChannelSms(channelSms string) ViberSendOption {
	return func(opts *viberSendOptions) {
		opts.channelSms = channelSms
	}
}

// WithViberSendTextSms sets SMS text
func WithViberSendTextSms(textSms string) ViberSendOption {
	return func(opts *viberSendOptions) {
		opts.textSms = textSms
	}
}

// WithViberSendPriceSms sets SMS price
func WithViberSendPriceSms(priceSms int) ViberSendOption {
	return func(opts *viberSendOptions) {
		opts.priceSms = priceSms
	}
}

// WithViberSendTimeout sets SMS sending timeout for CASCADE channel.
// Possible values: 0.25 (15 min), 0.5 (30 min), 1 (1 hour), 3 (3 hours),
// 6 (6 hours), 12 (12 hours), 24 (24 hours)
func WithViberSendTimeout(timeout float64) ViberSendOption {
	return func(opts *viberSendOptions) {
		opts.timeout = timeout
	}
}

// ViberSend sends a Viber message.
//
// This method sends a Viber message to specified phone numbers or groups.
// Viber messages can include text, images, and buttons for interactive content.
// If the Viber message cannot be delivered, an SMS fallback can be configured.
//
// Parameters:
//   - sign (string): The signature of the message (2-64 characters)
//   - channel (string): The channel of the message (e.g., "OFFICIAL")
//   - text (string): The text content of the message (2-640 characters)
//   - opts (...ViberSendOption): Optional parameters for customizing the message
//
// Options:
//   - WithViberSendNumber(number): Set phone number to send to
//   - WithViberSendGroupID(groupID): Set group ID to send to
//   - WithViberSendImageSource(imageSource): Set image source URL
//   - WithViberSendTextButton(textButton): Set button text
//   - WithViberSendLinkButton(linkButton): Set button link
//   - WithViberSendDateSend(dateSend): Set send date
//   - WithViberSendSignSms(signSms): Set SMS fallback signature
//   - WithViberSendChannelSms(channelSms): Set SMS fallback channel
//   - WithViberSendTextSms(textSms): Set SMS fallback text
//   - WithViberSendPriceSms(priceSms): Set SMS fallback price
//   - WithViberSendTimeout(timeout): Set SMS timeout (0.25, 0.5, 1, 3, 6, 12, 24 hours)
//
// Returns:
//   - interface{}: The response data containing message details
//   - error: error if the request fails or API returns an error
//
// Example response:
//
//	{
//	  "success": true,
//	  "data": {
//	    "id": 12345,
//	    "number": "79031234567",
//	    "count": 1,
//	    "sign": "Viber",
//	    "channel": "OFFICIAL",
//	    "text": "Hello, World!",
//	    "cost": 2.25,
//	    "status": 1,
//	    "extendStatus": "moderation",
//	    "dateCreate": 1511153253,
//	    "dateSend": 1511153253,
//	    "countSend": 0,
//	    "countDelivered": 0,
//	    "countWrite": 0,
//	    "countUndelivered": 0,
//	    "countError": 0
//	  }
//	}
//
// Example usage:
//
//	result, err := client.ViberSend("MyCompany", "OFFICIAL", "Hello, World!",
//	    WithViberSendNumber("79031234567"),
//	    WithViberSendImageSource("https://example.com/image.jpg"),
//	    WithViberSendTextButton("Click here"),
//	    WithViberSendLinkButton("https://example.com"))
func (c *Client) ViberSend(sign string, channel string, text string, opts ...ViberSendOption) (interface{}, error) {
	// Validate required parameters
	if err := c.validator.ValidateSign(sign); err != nil {
		return nil, err
	}

	if channel == "" {
		return nil, NewValidationError("channel", "channel cannot be empty")
	}

	if err := c.validator.ValidateMessage(text); err != nil {
		return nil, err
	}

	// Process options
	options := &viberSendOptions{}
	for _, opt := range opts {
		opt(options)
	}

	// Validate phone number if specified
	if options.number != "" {
		if err := c.validator.ValidatePhone(options.number); err != nil {
			return nil, err
		}
	}

	// Validate group ID if specified
	if options.groupID > 0 {
		if err := c.validator.ValidateGroupID(options.groupID); err != nil {
			return nil, err
		}
	}

	response := new(ViberSendMsg)

	data := url.Values{}
	data.Set("sign", sign)
	data.Set("channel", channel)
	data.Set("text", text)

	if options.number != "" {
		data.Set("number", options.number)
	}
	if options.groupID > 0 {
		data.Set("groupId", strconv.Itoa(options.groupID))
	}
	if options.imageSource != "" {
		data.Set("imageSource", options.imageSource)
	}
	if options.textButton != "" {
		data.Set("textButton", options.textButton)
	}
	if options.linkButton != "" {
		data.Set("linkButton", options.linkButton)
	}
	if options.dateSend != "" {
		data.Set("dateSend", options.dateSend)
	}
	if options.signSms != "" {
		data.Set("signSms", options.signSms)
	}
	if options.channelSms != "" {
		data.Set("channelSms", options.channelSms)
	}
	if options.textSms != "" {
		data.Set("textSms", options.textSms)
	}
	if options.priceSms > 0 {
		data.Set("priceSms", strconv.Itoa(options.priceSms))
	}
	if options.timeout > 0 {
		data.Set("timeout", strconv.FormatFloat(options.timeout, 'f', -1, 64))
	}

	if err := c.executeRequest("viber/send", response, data); err != nil {
		return nil, err
	}

	if response.IsErrorResponse() {
		return nil, response.GetError()
	}

	return response.Data, nil
}
