package smsaero_golang

import (
	"net/url"
	"strconv"
)

// ViberStatisticItem represents a single item in Viber statistic response
type ViberStatisticItem struct {
	Number       string `json:"number"`
	Status       int    `json:"status"`
	ExtendStatus string `json:"extendStatus"`
	DateSend     int    `json:"dateSend"`
}

type ViberStatisticMsg struct {
	Data interface{} `json:"data"`
	ErrorResponse
}

// ViberStatistic retrieves statistics for a specific Viber campaign.
//
// This method returns detailed statistics about a Viber campaign including
// delivery status for each recipient.
//
// Parameters:
//   - sendingID (int): The ID of the Viber campaign to get statistics for
//
// Returns:
//   - interface{}: The response data containing campaign statistics
//   - error: error if the request fails or API returns an error
//
// Example response:
//
//	{
//	  "success": true,
//	  "data": {
//	    "0": {
//	      "number": "79990000000",
//	      "status": 0,
//	      "extendStatus": "send",
//	      "dateSend": 1511153341
//	    },
//	    "1": {
//	      "number": "79990000001",
//	      "status": 2,
//	      "extendStatus": "write",
//	      "dateSend": 1511153341
//	    }
//	  }
//	}
//
// Status values:
//   - 0: sent
//   - 1: delivered
//   - 2: read
//   - 3: not delivered
//   - 4: error
//
// Example usage:
//
//	stats, err := client.ViberStatistic(12345)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Printf("Campaign statistics: %+v\n", stats)
func (c *Client) ViberStatistic(sendingID int) (interface{}, error) {
	if sendingID <= 0 {
		return nil, NewValidationError("sendingId", "sendingId must be a positive integer")
	}

	response := new(ViberStatisticMsg)

	data := url.Values{}
	data.Set("sendingId", strconv.Itoa(sendingID))

	if err := c.executeRequest("viber/statistic", response, data); err != nil {
		return nil, err
	}

	if response.IsErrorResponse() {
		return nil, response.GetError()
	}

	return response.Data, nil
}
