package smsaero_golang

type Auth struct {
	Data interface{}
	ErrorResponse
}

// IsAuthorized checks if the client is authorized to use the SMS Aero API.
//
// This method verifies the authentication credentials (API key and email) by making
// a request to the auth endpoint. It returns true if the credentials are valid,
// false if they are invalid, or an error if the request fails.
//
// Returns:
//   - bool: true if authorized, false if not authorized
//   - error: error if the request fails or API returns an error
//
// Example response:
//
//	{
//	  "success": true,
//	  "data": {}
//	}
//
// Usage:
//
//	authorized, err := client.IsAuthorized()
//	if err != nil {
//	    log.Fatal(err)
//	}
//	if authorized {
//	    fmt.Println("Client is authorized")
//	} else {
//	    fmt.Println("Client is not authorized")
//	}
func (c *Client) IsAuthorized() (bool, error) {
	response := new(Auth)

	if err := c.executeRequest(`auth`, response, nil); err != nil {
		return false, err
	}

	if response.IsErrorResponse() {
		return false, response.GetError()
	}

	return response.Success, nil
}
