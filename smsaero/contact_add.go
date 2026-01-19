package smsaero_golang

import (
	"net/url"
	"strconv"
)

type ContactAddMsg struct {
	Data BlackListAdd `json:"data"`
	ErrorResponse
}

// ContactAddOption represents an option for adding a contact
type ContactAddOption func(*contactAddOptions)

type contactAddOptions struct {
	groupID   int
	birthday  string
	sex       string
	lastName  string
	firstName string
	surname   string
	param1    string
	param2    string
	param3    string
}

// WithContactAddGroupID sets the group ID for the contact
func WithContactAddGroupID(groupID int) ContactAddOption {
	return func(opts *contactAddOptions) {
		opts.groupID = groupID
	}
}

// WithContactAddBirthday sets the birthday for the contact (format: YYYY-MM-DD)
func WithContactAddBirthday(birthday string) ContactAddOption {
	return func(opts *contactAddOptions) {
		opts.birthday = birthday
	}
}

// WithContactAddSex sets the gender for the contact (male/female)
func WithContactAddSex(sex string) ContactAddOption {
	return func(opts *contactAddOptions) {
		opts.sex = sex
	}
}

// WithContactAddLastName sets the last name for the contact
func WithContactAddLastName(lastName string) ContactAddOption {
	return func(opts *contactAddOptions) {
		opts.lastName = lastName
	}
}

// WithContactAddFirstName sets the first name for the contact
func WithContactAddFirstName(firstName string) ContactAddOption {
	return func(opts *contactAddOptions) {
		opts.firstName = firstName
	}
}

// WithContactAddSurname sets the middle name for the contact
func WithContactAddSurname(surname string) ContactAddOption {
	return func(opts *contactAddOptions) {
		opts.surname = surname
	}
}

// WithContactAddParam1 sets additional parameter 1 for the contact
func WithContactAddParam1(param1 string) ContactAddOption {
	return func(opts *contactAddOptions) {
		opts.param1 = param1
	}
}

// WithContactAddParam2 sets additional parameter 2 for the contact
func WithContactAddParam2(param2 string) ContactAddOption {
	return func(opts *contactAddOptions) {
		opts.param2 = param2
	}
}

// WithContactAddParam3 sets additional parameter 3 for the contact
func WithContactAddParam3(param3 string) ContactAddOption {
	return func(opts *contactAddOptions) {
		opts.param3 = param3
	}
}

// ContactAdd adds a new contact to the address book.
//
// Parameters:
//   - number (string): The contact's phone number in international format (e.g., "79031234567")
//   - opts (...ContactAddOption): Optional parameters for contact details
//
// Options:
//   - WithContactAddGroupID(id): Assign contact to a specific group
//   - WithContactAddFirstName(name): Set the contact's first name
//   - WithContactAddLastName(name): Set the contact's last name
//   - WithContactAddSurname(name): Set the contact's middle name
//   - WithContactAddBirthday(date): Set birthday (format: YYYY-MM-DD)
//   - WithContactAddSex(gender): Set gender (male/female)
//   - WithContactAddParam1-3(value): Set additional custom parameters
//
// Returns:
//   - BlackListAdd: The server's response containing contact details
//   - error: Any error that occurred during the request
//
// Example response:
//
//	{
//	  "id": 12345,
//	  "number": "79031234567",
//	  "firstName": "John",
//	  "lastName": "Doe",
//	  "groupID": 1
//	}
//
// Example usage:
//
//	contact, err := client.ContactAdd("79031234567",
//	    WithContactAddFirstName("John"),
//	    WithContactAddLastName("Doe"),
//	    WithContactAddGroupID(1))
func (c *Client) ContactAdd(number string, opts ...ContactAddOption) (BlackListAdd, error) {
	empty := BlackListAdd{}

	// Validate phone number
	if err := c.validator.ValidatePhone(number); err != nil {
		return empty, err
	}

	// Process options
	options := &contactAddOptions{}
	for _, opt := range opts {
		opt(options)
	}

	// Validate group ID if specified
	if options.groupID > 0 {
		if err := c.validator.ValidateGroupID(options.groupID); err != nil {
			return empty, err
		}
	}

	response := new(ContactAddMsg)

	data := url.Values{}
	data.Set("number", number)

	if options.groupID > 0 {
		data.Set("groupId", strconv.Itoa(options.groupID))
	}
	if options.birthday != "" {
		data.Set("birthday", options.birthday)
	}
	if options.sex != "" {
		data.Set("sex", options.sex)
	}
	if options.lastName != "" {
		data.Set("lname", options.lastName)
	}
	if options.firstName != "" {
		data.Set("fname", options.firstName)
	}
	if options.surname != "" {
		data.Set("sname", options.surname)
	}
	if options.param1 != "" {
		data.Set("param1", options.param1)
	}
	if options.param2 != "" {
		data.Set("param2", options.param2)
	}
	if options.param3 != "" {
		data.Set("param3", options.param3)
	}

	if err := c.executeRequest("contact/add", response, data); err != nil {
		return empty, err
	}

	if response.IsErrorResponse() {
		return empty, response.GetError()
	}

	return response.Data, nil
}
