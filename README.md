# SmsAero GoLang Api


### Installation:

    $ go get github.com/smsaero/smsaero_golang


### Usage:
    
    package main
    
    
    import (
        "fmt"
        "github.com/smsaero/smsaero_golang"
    )
    
    
    const (
        // Get credentials from account settings page: https://smsaero.ru/cabinet/settings/apikey/
        email  = "your email"
        apiKey = "your api key"
    )
    
    
    func main() {
        client := smsaero_golang.NewSmsAeroClient(email, apiKey)
    
        balance, err := client.Balance()
        if err != nil {
            panic(err)
        }
    
        if balance <= 0 {
            panic("Insufficient balance")
        }
    
        if sendResult, err := client.Send("70000000000", "Hello, World!", "SMSAERO"); err == nil {
            fmt.Println(sendResult.Id)
        } else {
            panic(err)
        }
    }


### License

    MIT License
