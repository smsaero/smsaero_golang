package smsaero_golang

import "net/url"

type BlackListAdd struct {
	ID     int    `json:"id"`
	Number string `json:"number"`
	Fname  string `json:"fname"`
	Lname  string `json:"lname"`
	Sname  string `json:"sname"`
	Bday   string `json:"bday"`
	Sex    string `json:"sex"`
	Param  string `json:"param"`
	Param2 string `json:"param2"`
	Param3 string `json:"param3"`
}

type BlackListAddMsg struct {
	Data BlackListAdd `json:"data"`
	ErrorResponse
}

// BlackListAddOption represents an option for adding to blacklist
type BlackListAddOption func(*blackListAddOptions)

type blackListAddOptions struct {
	birthday  string
	sex       string
	lastName  string
	firstName string
	surname   string
	param1    string
	param2    string
	param3    string
}

// WithBlackListAddBirthday sets the birthday for the blacklist entry (Unix timestamp)
func WithBlackListAddBirthday(birthday string) BlackListAddOption {
	return func(opts *blackListAddOptions) {
		opts.birthday = birthday
	}
}

// WithBlackListAddSex sets the gender for the blacklist entry (male/female)
func WithBlackListAddSex(sex string) BlackListAddOption {
	return func(opts *blackListAddOptions) {
		opts.sex = sex
	}
}

// WithBlackListAddLastName sets the last name for the blacklist entry
func WithBlackListAddLastName(lastName string) BlackListAddOption {
	return func(opts *blackListAddOptions) {
		opts.lastName = lastName
	}
}

// WithBlackListAddFirstName sets the first name for the blacklist entry
func WithBlackListAddFirstName(firstName string) BlackListAddOption {
	return func(opts *blackListAddOptions) {
		opts.firstName = firstName
	}
}

// WithBlackListAddSurname sets the middle name for the blacklist entry
func WithBlackListAddSurname(surname string) BlackListAddOption {
	return func(opts *blackListAddOptions) {
		opts.surname = surname
	}
}

// WithBlackListAddParam1 sets additional parameter 1 for the blacklist entry
func WithBlackListAddParam1(param1 string) BlackListAddOption {
	return func(opts *blackListAddOptions) {
		opts.param1 = param1
	}
}

// WithBlackListAddParam2 sets additional parameter 2 for the blacklist entry
func WithBlackListAddParam2(param2 string) BlackListAddOption {
	return func(opts *blackListAddOptions) {
		opts.param2 = param2
	}
}

// WithBlackListAddParam3 sets additional parameter 3 for the blacklist entry
func WithBlackListAddParam3(param3 string) BlackListAddOption {
	return func(opts *blackListAddOptions) {
		opts.param3 = param3
	}
}

// BlackListAdd adds a phone number to the blacklist.
//
// This method adds a phone number to the blacklist to prevent sending SMS messages
// to it. The phone number will be blocked from receiving any SMS messages.
//
// Parameters:
//   - number (string): The phone number to add to the blacklist (e.g., "79031234567")
//   - opts (...BlackListAddOption): Optional parameters for blacklist entry details
//
// Options:
//   - WithBlackListAddBirthday(birthday): Set birthday (Unix timestamp)
//   - WithBlackListAddSex(sex): Set gender (male/female)
//   - WithBlackListAddLastName(lname): Set last name
//   - WithBlackListAddFirstName(fname): Set first name
//   - WithBlackListAddSurname(sname): Set middle name
//   - WithBlackListAddParam1-3(value): Set additional custom parameters
//
// Returns:
//   - BlackListAdd: The response data containing the added blacklist entry details
//   - error: error if the request fails or API returns an error
//
// Example response:
//
//	{
//	  "success": true,
//	  "data": {
//	    "id": 12345,
//	    "number": "79031234567",
//	    "fname": "John",
//	    "lname": "Doe",
//	    "sname": "Smith",
//	    "bday": "1990-01-01",
//	    "sex": "male",
//	    "param": "param1",
//	    "param2": "param2",
//	    "param3": "param3"
//	  }
//	}
//
// Usage:
//
//	result, err := client.BlackListAdd("79031234567",
//	    WithBlackListAddFirstName("John"),
//	    WithBlackListAddLastName("Doe"),
//	    WithBlackListAddSex("male"))
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Printf("Added to blacklist with ID: %d\n", result.ID)
func (c *Client) BlackListAdd(number string, opts ...BlackListAddOption) (BlackListAdd, error) {
	response := new(BlackListAddMsg)
	empty := BlackListAdd{}

	// Process options
	options := &blackListAddOptions{}
	for _, opt := range opts {
		opt(options)
	}

	data := url.Values{}
	data.Set("number", number)

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

	if err := c.executeRequest("blacklist/add", response, data); err != nil {
		return empty, err
	}

	if response.IsErrorResponse() {
		return empty, response.GetError()
	}

	return response.Data, nil
}
