package smsaero_golang

import "net/url"

const (
	minGroupNameLength = 2
	maxGroupNameLength = 64
)

type GroupAdd struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type GroupAddMsg struct {
	Data GroupAdd `json:"data"`
	ErrorResponse
}

// GroupAdd adds a new group to the SmsAero service.
//
// This method creates a new contact group that can be used to organize contacts.
// The group name must be between 2 and 64 characters long.
//
// Parameters:
//   - name (string): The name of the group to be added (2-64 characters)
//
// Returns:
//   - GroupAdd: The response data containing the created group details
//   - error: error if the request fails or API returns an error
//
// Example response:
//
//	{
//	  "success": true,
//	  "data": {
//	    "id": 12345,
//	    "name": "TestGroup"
//	  }
//	}
//
// Example usage:
//
//	group, err := client.GroupAdd("My Contacts")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Printf("Created group with ID: %d\n", group.Id)
func (c *Client) GroupAdd(name string) (GroupAdd, error) {
	empty := GroupAdd{}

	// Validate group name
	if name == "" {
		return empty, NewValidationError("name", "group name cannot be empty")
	}

	if len(name) < minGroupNameLength {
		return empty, NewValidationError("name", "group name is too short (minimum 2 characters)")
	}

	if len(name) > maxGroupNameLength {
		return empty, NewValidationError("name", "group name is too long (maximum 64 characters)")
	}

	response := new(GroupAddMsg)

	data := url.Values{}
	data.Set("name", name)

	if err := c.executeRequest("group/add", response, data); err != nil {
		return empty, err
	}

	if response.IsErrorResponse() {
		return empty, response.GetError()
	}

	return response.Data, nil
}
