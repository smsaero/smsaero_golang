package smsaero_golang

import (
	"fmt"
	"net/url"

	"github.com/nyaruka/phonenumbers"
)

type SendSmsOption func(*sendSmsOptions)

type sendSmsOptions struct {
	sign string
}

func WithSendSmsSign(sign string) SendSmsOption {
	return func(opts *sendSmsOptions) {
		opts.sign = sign
	}
}

type SendSms struct {
	Id           int
	From         string
	Number       string
	Text         string
	Status       int
	ExtendStatus string
	Channel      string
	Cost         float64
	DateCreate   int
	DateSend     int
}

type SendMsg struct {
	Data SendSms
	ErrorResponse
}

func (c *Client) SendSms(number int, text string, opts ...SendSmsOption) (SendSms, error) {
	options := &sendSmsOptions{
		sign: c.sign,
	}

	for _, opt := range opts {
		opt(options)
	}

	response := new(SendMsg)
	empty := SendSms{}

	if !c.phoneIsValid(fmt.Sprintf("%d", number)) {
		return empty, fmt.Errorf("invalid phone number")
	}

	data := url.Values{}
	data.Set("number", fmt.Sprintf("%d", number))
	data.Set("sign", c.sign)
	data.Set("text", text)

	path := c.getApiPath("sms/send", "sms/testsend")

	if err := c.executeRequest(path, response, data); err != nil {
		return empty, err
	}

	if response.IsErrorResponse() {
		return empty, response.GetError()
	}

	return response.Data, nil
}

func (c *Client) phoneIsValid(phone string) bool {
	if !c.phoneValidation {
		return true
	}
	metadata, err := phonenumbers.Parse("+"+phone, "RU")
	if err != nil {
		return false
	}

	if !phonenumbers.IsPossibleNumber(metadata) {
		return false
	}

	if !phonenumbers.IsValidNumber(metadata) {
		return false
	}
	return true
}
