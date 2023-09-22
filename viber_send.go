package smsaero_golang

import "net/url"

type ViberSendMsg struct {
	Data interface{}
	ErrorResponse
}

func (c *Client) ViberSend(sign string, channel string, text string, number string, groupId string, imageSource string, textButton string, linkButton string, dateSend string, signSms string, channelSms string, textSms string, priceSms string) (interface{}, error) {
	response := new(ViberSendMsg)

	data := url.Values{}
	data.Set("sign", sign)
	data.Set("channel", channel)
	data.Set("text", text)

	if number != "" {
		data.Set("number", number)
	}
	if groupId != "" {
		data.Set("groupId", groupId)
	}
	if imageSource != "" {
		data.Set("imageSource", imageSource)
	}
	if textButton != "" {
		data.Set("textButton", textButton)
	}
	if linkButton != "" {
		data.Set("linkButton", linkButton)
	}
	if dateSend != "" {
		data.Set("dateSend", dateSend)
	}
	if signSms != "" {
		data.Set("signSms", signSms)
	}
	if channelSms != "" {
		data.Set("channelSms", channelSms)
	}
	if textSms != "" {
		data.Set("textSms", textSms)
	}
	if priceSms != "" {
		data.Set("priceSms string", priceSms)
	}

	if err := c.executeRequest(`viber/send`, response, data); err != nil {
		return nil, err
	}

	if response.IsErrorResponse() {
		return nil, response.GetError()
	}

	return response.Data, nil
}
