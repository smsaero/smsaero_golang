# SmsAero GoLang Api client

Library for sending SMS messages using the SmsAero API. Written in GoLang.

## Installation:

```bash
go get github.com/smsaero/smsaero_golang
```

If you are starting from scratch in a new directory, you will first need to create a `go.mod` file for tracking
dependencies:

```bash
go mod init smsaero-example
```

## Usage example:

Get credentials from account settings page: https://smsaero.ru/cabinet/settings/apikey/

```golang
package main

import (
    "fmt"
    "time"
    "github.com/smsaero/smsaero_golang/smsaero"
)

const (
    Email  = "your email"
    ApiKey = "your api key"
)

func main() {
    client := smsaero_golang.NewSmsAeroClient(
        Email, ApiKey,
        smsaero_golang.WithTimeout(time.Second*10),
        smsaero_golang.WithTest(true),
        smsaero_golang.WithPhoneValidation(false),
    )

    if sendResult, err := client.SendSms(70000000000, "Hello, World!"); err == nil {
        fmt.Println(sendResult.Id)
    } else {
        panic(err)
    }
}
```

## Installation dependencies:

```bash
go get github.com/smsaero/smsaero_golang@latest
go mod tidy
```

## Run on Docker:

```bash
docker pull 'smsaero/smsaero_golang:latest'
docker run -it --rm 'smsaero/smsaero_golang:latest' smsaero_send -email "your email" -api_key "your api key" -phone 70000000000 -message 'Hello, World!' -test
```

## License

```
MIT License
```
