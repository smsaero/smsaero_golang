# SMS Aero Go API Library v2

A Go library for working with the SMS Aero API service.

[![Go Version](https://img.shields.io/badge/go-1.21+-blue.svg)](https://golang.org/)
[![License](https://img.shields.io/badge/license-MIT-green.svg)](LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/smsaero/smsaero_golang/v2)](https://goreportcard.com/report/github.com/smsaero/smsaero_golang/v2)

## Features

- Complete SMS Aero API support
- Input validation and error handling
- Test mode for safe development
- SMS, Telegram, and Viber messaging
- Contact and group management
- HLR number checking
- Balance management

## Installation

```bash
go get github.com/smsaero/smsaero_golang/v2
```

If you are starting from scratch in a new directory, you will first need to create a `go.mod` file for tracking dependencies:

```bash
go mod init smsaero-example
```

## Quick Start

Get your API credentials from: https://smsaero.ru/cabinet/settings/apikey/

```go
package main

import (
    "fmt"
    "log"
    "time"
    
    smsaero "github.com/smsaero/smsaero_golang/v2/smsaero"
)

func main() {
    // Create client
    client, err := smsaero.NewSmsAeroClient(
        "your-email@example.com",
        "your-api-key-32-chars-long",
        smsaero.WithTimeout(30*time.Second),
        smsaero.WithTest(true), // Enable test mode
        smsaero.WithPhoneValidation(true),
    )
    if err != nil {
        log.Fatal("Failed to create client:", err)
    }

    // Send SMS
    result, err := client.SendSms(
        79031234567,
        "Hello from Go!",
        smsaero.WithSendSmsSign("SMS Aero"),
        smsaero.WithSendSmsCallbackURL("https://example.com/callback"),
    )
    if err != nil {
        log.Fatal("Failed to send SMS:", err)
    }
    
    fmt.Printf("SMS sent, ID: %d, Status: %s\n", result.ID, result.ExtendStatus)

    // Send Telegram code
    telegramResult, err := client.SendTelegram(
        79031234567,
        1234,
        smsaero.WithSendTelegramSign("SMS Aero"),
        smsaero.WithSendTelegramText("Your code: 1234"),
    )
    if err != nil {
        log.Fatal("Failed to send Telegram:", err)
    }
    
    fmt.Printf("Telegram code sent, ID: %d\n", telegramResult.ID)
}
```

## Client Configuration

```go
client, err := smsaero.NewSmsAeroClient(
    "email@example.com",
    "api-key-32-chars-long",
    smsaero.WithTimeout(30*time.Second),
    smsaero.WithTest(true),
    smsaero.WithPhoneValidation(true),
    smsaero.WithSign("SMS Aero"),
    smsaero.WithLogger(logger),
    smsaero.WithHTTPClient(customHTTPClient),
)
```

## Available Methods

### SMS
- `SendSms(phone int, text string, options ...SendSmsOption) (*SendSms, error)`

### Telegram
- `SendTelegram(phone int, code int, options ...SendTelegramOption) (*SendTelegram, error)`

### Viber
- `ViberSend(sign, channel, text string, options ...ViberSendOption) (interface{}, error)`

### Contacts
- `ContactAdd(phone string, options ...ContactAddOption) (*ContactAdd, error)`
- `ContactList() (interface{}, error)`
- `ContactDelete(phone string) (interface{}, error)`

### Groups
- `GroupAdd(name string) (*GroupAdd, error)`
- `GroupList() (interface{}, error)`
- `GroupDelete(id int) (interface{}, error)`

### Balance
- `Balance() (float64, error)`
- `AddBalance(amount float64, cardID string) (interface{}, error)`

### HLR
- `HlrCheck(phone int) (*HlrCheck, error)`
- `HlrStatus(id int) (interface{}, error)`

### Other
- `IsAuthorized() (bool, error)`
- `SignList() (interface{}, error)`
- `Tariffs() (interface{}, error)`

## Error Handling

The library provides specific error types for different error conditions:

```go
result, err := client.SendSms(79031234567, "Hello")
if err != nil {
    switch e := err.(type) {
    case *smsaero.SmsAeroValidationError:
        fmt.Printf("Validation error: %s\n", e.Error())
    case *smsaero.SmsAeroNoMoneyError:
        fmt.Printf("Insufficient funds: %s\n", e.Error())
    case *smsaero.SmsAeroAPIError:
        fmt.Printf("API error: %s\n", e.Error())
    case *smsaero.SmsAeroConnectionError:
        fmt.Printf("Connection error: %s\n", e.Error())
    default:
        fmt.Printf("Unknown error: %s\n", e.Error())
    }
}
```

## Installation Dependencies

```bash
go get github.com/smsaero/smsaero_golang/v2@latest
go mod tidy
```

## Run on Docker

```bash
docker pull 'smsaero/smsaero_golang:latest'
docker run -it --rm 'smsaero/smsaero_golang:latest' smsaero_send \
    -email "your email" \
    -api_key "your api key" \
    -phone 70000000000 \
    -message 'Hello, World!' \
    -test
```

## License

MIT License
