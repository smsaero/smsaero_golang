package smsaero_golang

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/nyaruka/phonenumbers"
)

const (
	maxMessageLength = 640
	minMessageLength = 2
	minSignLength    = 2
	maxSignLength    = 64
	maxDaysInFuture  = 7
)

// Validator provides methods for input data validation
type Validator struct {
	phoneValidation bool
}

// NewValidator creates a new validator instance
func NewValidator(phoneValidation bool) *Validator {
	return &Validator{
		phoneValidation: phoneValidation,
	}
}

// ValidatePhone validates phone number correctness
func (v *Validator) ValidatePhone(phone string) error {
	if phone == "" {
		return ErrInvalidPhoneNumber
	}

	// Remove all non-digit characters except +
	cleaned := regexp.MustCompile(`[^\d+]`).ReplaceAllString(phone, "")

	// If number doesn't start with +, add it
	if !strings.HasPrefix(cleaned, "+") {
		cleaned = "+" + cleaned
	}

	if v.phoneValidation {
		metadata, err := phonenumbers.Parse(cleaned, "RU")
		if err != nil {
			return NewValidationError("phone", fmt.Sprintf("failed to parse number: %v", err))
		}

		if !phonenumbers.IsPossibleNumber(metadata) {
			return NewValidationError("phone", "phone number is not possible")
		}

		if !phonenumbers.IsValidNumber(metadata) {
			return NewValidationError("phone", "phone number is invalid")
		}
	}

	return nil
}

// ValidateMessage validates message correctness
func (v *Validator) ValidateMessage(message string) error {
	if message == "" {
		return ErrEmptyMessage
	}

	if len(message) > maxMessageLength {
		return ErrMessageTooLong
	}

	if len(message) < minMessageLength {
		return NewValidationError("message", "message is too short (minimum 2 characters)")
	}

	return nil
}

// ValidateSign validates sender signature correctness
func (v *Validator) ValidateSign(sign string) error {
	if sign == "" {
		return ErrInvalidSign
	}

	if len(sign) < minSignLength {
		return NewValidationError("sign", "signature is too short (minimum 2 characters)")
	}

	if len(sign) > maxSignLength {
		return NewValidationError("sign", "signature is too long (maximum 64 characters)")
	}

	return nil
}

// ValidateTelegramCode validates Telegram code correctness
func (v *Validator) ValidateTelegramCode(code int) error {
	codeStr := strconv.Itoa(code)
	if len(codeStr) < 4 || len(codeStr) > 8 {
		return ErrInvalidCode
	}

	return nil
}

// ValidateCredentials validates credentials correctness
func (v *Validator) ValidateCredentials(username, password string) error {
	if username == "" {
		return NewValidationError("username", "username cannot be empty")
	}

	if password == "" {
		return NewValidationError("password", "password cannot be empty")
	}

	if len(password) < 16 || len(password) > 32 {
		return NewValidationError("password", "password length must be between 16 and 32 characters")
	}

	return nil
}

// ValidateCallbackURL validates callback URL correctness
func (v *Validator) ValidateCallbackURL(url string) error {
	if url == "" {
		return nil // URL is optional
	}

	// Simple URL validation
	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
		return NewValidationError("callback_url", "URL must start with http:// or https://")
	}

	return nil
}

// ValidateDateToSend validates send date correctness
func (v *Validator) ValidateDateToSend(dateToSend *time.Time) error {
	if dateToSend == nil {
		return nil // Date is optional
	}

	now := time.Now()
	if dateToSend.Before(now) {
		return NewValidationError("date_to_send", "send date cannot be in the past")
	}

	// Maximum 7 days ahead
	maxDate := now.AddDate(0, 0, maxDaysInFuture)
	if dateToSend.After(maxDate) {
		return NewValidationError("date_to_send", "send date cannot be more than 7 days in the future")
	}

	return nil
}

// ValidatePage validates page number correctness
func (v *Validator) ValidatePage(page int) error {
	if page < 1 {
		return NewValidationError("page", "page number must be greater than 0")
	}

	return nil
}

// ValidateGroupID validates group ID correctness
func (v *Validator) ValidateGroupID(groupID int) error {
	if groupID < 1 {
		return NewValidationError("group_id", "group ID must be greater than 0")
	}

	return nil
}

// ValidateContactID validates contact ID correctness
func (v *Validator) ValidateContactID(contactID int) error {
	if contactID < 1 {
		return NewValidationError("contact_id", "contact ID must be greater than 0")
	}

	return nil
}

// ValidateBlacklistID validates blacklist entry ID correctness
func (v *Validator) ValidateBlacklistID(blacklistID int) error {
	if blacklistID < 1 {
		return NewValidationError("blacklist_id", "blacklist entry ID must be greater than 0")
	}

	return nil
}

// ValidateSmsID validates SMS ID correctness
func (v *Validator) ValidateSmsID(smsID int) error {
	if smsID < 1 {
		return NewValidationError("sms_id", "SMS ID must be greater than 0")
	}

	return nil
}

// ValidateTelegramID validates Telegram message ID correctness
func (v *Validator) ValidateTelegramID(telegramID int) error {
	if telegramID < 1 {
		return NewValidationError("telegram_id", "Telegram message ID must be greater than 0")
	}

	return nil
}

// ValidateHlrID validates HLR check ID correctness
func (v *Validator) ValidateHlrID(hlrID int) error {
	if hlrID < 1 {
		return NewValidationError("hlr_id", "HLR check ID must be greater than 0")
	}

	return nil
}

// ValidateCardID validates card ID correctness
func (v *Validator) ValidateCardID(cardID int) error {
	if cardID < 1 {
		return NewValidationError("card_id", "card ID must be greater than 0")
	}

	return nil
}

// ValidateAmount validates amount correctness
func (v *Validator) ValidateAmount(amount float64) error {
	if amount <= 0 {
		return NewValidationError("amount", "amount must be greater than 0")
	}

	return nil
}
