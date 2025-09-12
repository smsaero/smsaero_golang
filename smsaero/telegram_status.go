package smsaero_golang

import (
	"fmt"
	"net/url"
)

type TelegramStatus struct {
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

type TelegramStatusMsg struct {
	Data TelegramStatus
	ErrorResponse
}

func (c *Client) TelegramStatus(telegramId int) (TelegramStatus, error) {
	response := new(TelegramStatusMsg)
	empty := TelegramStatus{}

	data := url.Values{}
	data.Set("id", fmt.Sprintf("%d", telegramId))

	if err := c.executeRequest(`telegram/status`, response, data); err != nil {
		return empty, err
	}

	if response.IsErrorResponse() {
		return empty, response.GetError()
	}

	return response.Data, nil
}
