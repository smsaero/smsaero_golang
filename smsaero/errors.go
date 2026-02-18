package smsaero_golang

import (
	"errors"
	"fmt"
)

// SmsAeroError represents a base error type for SMS Aero operations
type SmsAeroError struct {
	Message string `json:"message"`
	Code    string `json:"code"`
}

func (e *SmsAeroError) Error() string {
	if e.Code != "" {
		return fmt.Sprintf("[%s] %s", e.Code, e.Message)
	}

	return e.Message
}

// SmsAeroConnectionError represents connection-related errors
type SmsAeroConnectionError struct {
	*SmsAeroError
	URL string `json:"url"`
}

func (e *SmsAeroConnectionError) Error() string {
	if e.URL != "" {
		return fmt.Sprintf("connection error to %s: %s", e.URL, e.Message)
	}

	return "connection error: " + e.Message
}

// SmsAeroNoMoneyError represents insufficient funds errors
type SmsAeroNoMoneyError struct {
	*SmsAeroError
	Balance float64 `json:"balance"`
}

func (e *SmsAeroNoMoneyError) Error() string {
	return fmt.Sprintf("insufficient funds: %s (balance: %.2f)", e.Message, e.Balance)
}

// SmsAeroAPIError represents API-related errors
type SmsAeroAPIError struct {
	*SmsAeroError
	StatusCode int `json:"statusCode"`
}

func (e *SmsAeroAPIError) Error() string {
	return fmt.Sprintf("API error (status %d): %s", e.StatusCode, e.Message)
}

// SmsAeroValidationError represents validation errors
type SmsAeroValidationError struct {
	*SmsAeroError
	Field string `json:"field"`
}

func (e *SmsAeroValidationError) Error() string {
	if e.Field != "" {
		return fmt.Sprintf("validation error for field '%s': %s", e.Field, e.Message)
	}

	return "validation error: " + e.Message
}

// SmsAeroAuthError represents authentication errors
type SmsAeroAuthError struct {
	*SmsAeroError
}

func (e *SmsAeroAuthError) Error() string {
	return "authentication error: " + e.Message
}

// Constructor functions
func NewConnectionError(message string) *SmsAeroConnectionError {
	return &SmsAeroConnectionError{
		SmsAeroError: &SmsAeroError{Message: message, Code: "CONNECTION_ERROR"},
	}
}

func NewNoMoneyError(message string, balance float64) *SmsAeroNoMoneyError {
	return &SmsAeroNoMoneyError{
		SmsAeroError: &SmsAeroError{Message: message, Code: "NO_MONEY"},
		Balance:      balance,
	}
}

func NewValidationError(field, message string) *SmsAeroValidationError {
	return &SmsAeroValidationError{
		SmsAeroError: &SmsAeroError{Message: message, Code: "VALIDATION_ERROR"},
		Field:        field,
	}
}

func NewAPIError(message string, statusCode int) *SmsAeroAPIError {
	return &SmsAeroAPIError{
		SmsAeroError: &SmsAeroError{Message: message, Code: "API_ERROR"},
		StatusCode:   statusCode,
	}
}

func NewAuthError(message string) *SmsAeroAuthError {
	return &SmsAeroAuthError{
		SmsAeroError: &SmsAeroError{Message: message, Code: "AUTH_ERROR"},
	}
}

// Predefined errors
var (
	ErrInvalidPhoneNumber = &SmsAeroValidationError{
		SmsAeroError: &SmsAeroError{
			Message: "invalid phone number format",
			Code:    "VALIDATION_ERROR",
		},
	}
	ErrEmptyMessage = &SmsAeroValidationError{
		SmsAeroError: &SmsAeroError{
			Message: "message cannot be empty",
			Code:    "VALIDATION_ERROR",
		},
	}
	ErrMessageTooLong = &SmsAeroValidationError{
		SmsAeroError: &SmsAeroError{
			Message: "message is too long",
			Code:    "VALIDATION_ERROR",
		},
	}
	ErrInvalidSign = &SmsAeroValidationError{
		SmsAeroError: &SmsAeroError{
			Message: "invalid sender signature",
			Code:    "VALIDATION_ERROR",
		},
	}
	ErrInvalidCode = &SmsAeroValidationError{
		SmsAeroError: &SmsAeroError{
			Message: "invalid Telegram code",
			Code:    "VALIDATION_ERROR",
		},
	}
	ErrEmptyCredentials = &SmsAeroValidationError{
		SmsAeroError: &SmsAeroError{
			Message: "credentials not specified",
			Code:    "VALIDATION_ERROR",
		},
	}
	ErrAllGatesUnavailable = &SmsAeroConnectionError{
		SmsAeroError: &SmsAeroError{
			Message: "all gateways unavailable",
			Code:    "CONNECTION_ERROR",
		},
	}
	ErrInvalidResponse = &SmsAeroAPIError{
		SmsAeroError: &SmsAeroError{
			Message: "invalid server response",
			Code:    "API_ERROR",
		},
	}
)

// Error type checking functions
func IsValidationError(err error) bool {
	var validationErr *SmsAeroValidationError
	ok := errors.As(err, &validationErr)

	return ok
}

func IsNoMoneyError(err error) bool {
	var noMoneyErr *SmsAeroNoMoneyError
	ok := errors.As(err, &noMoneyErr)

	return ok
}

func IsAPIError(err error) bool {
	var apiErr *SmsAeroAPIError
	ok := errors.As(err, &apiErr)

	return ok
}

func IsConnectionError(err error) bool {
	var connErr *SmsAeroConnectionError
	ok := errors.As(err, &connErr)

	return ok
}

func IsAuthError(err error) bool {
	var authErr *SmsAeroAuthError
	ok := errors.As(err, &authErr)

	return ok
}
