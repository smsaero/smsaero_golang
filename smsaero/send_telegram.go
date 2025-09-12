package smsaero_golang

import (
	"fmt"
	"net/url"
)

type SendTelegramOption func(*sendTelegramOptions)

type sendTelegramOptions struct {
	sign string
	text string
}

func WithSendTelegramSign(sign string) SendTelegramOption {
	return func(opts *sendTelegramOptions) {
		opts.sign = sign
	}
}

func WithSendTelegramText(text string) SendTelegramOption {
	return func(opts *sendTelegramOptions) {
		opts.text = text
	}
}

type SendTelegram struct {
	Id           int    `json:"id"`
	Number       string `json:"number"`
	TelegramCode string `json:"telegramCode"`
	SmsText      string `json:"smsText"`
	SmsFrom      string `json:"smsFrom"`
	IdSms        *int   `json:"idSms"`
	Status       int    `json:"status"`
	ExtendStatus string `json:"extendStatus"`
	Cost         string `json:"cost"`
	DateCreate   int    `json:"dateCreate"`
}

type SendTelegramMsg struct {
	Data SendTelegram
	ErrorResponse
}

func (c *Client) SendTelegram(number int, code int, opts ...SendTelegramOption) (SendTelegram, error) {
	options := &sendTelegramOptions{}

	for _, opt := range opts {
		opt(options)
	}

	response := new(SendTelegramMsg)
	empty := SendTelegram{}

	if !c.phoneIsValid(fmt.Sprintf("%d", number)) {
		return empty, fmt.Errorf("неверный номер телефона")
	}

	if code < 1000 || code > 99999999 {
		return empty, fmt.Errorf("код должен содержать от 4 до 8 цифр")
	}

	data := url.Values{}
	data.Set("number", fmt.Sprintf("%d", number))
	data.Set("code", fmt.Sprintf("%d", code))

	// опциональные параметры
	if options.sign != "" {
		data.Set("sign", options.sign)
	}
	if options.text != "" {
		data.Set("text", options.text)
	}

	if err := c.executeRequest(`telegram/send`, response, data); err != nil {
		return empty, err
	}

	if response.IsErrorResponse() {
		return empty, response.GetError()
	}

	return response.Data, nil
}
