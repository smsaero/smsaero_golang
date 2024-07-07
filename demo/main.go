package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	smsaero_golang "github.com/smsaero/smsaero_golang/smsaero"
)

func main() {
	email := flag.String("email", "", "Your email registered with SmsAero")
	apiKey := flag.String("api_key", "", "Your SmsAero API key")
	phone := flag.String("phone", "", "Phone number to send SMS to")
	message := flag.String("message", "", "Message to send")
	sign := flag.String("sign", "Sms Aero", "Sign to send SMS")
	testMode := flag.Bool("test", false, "Enable test mode")
	logging := flag.Bool("logging", false, "Enable logging")
	validate := flag.Bool("validate", true, "Disable phone number validation")

	flag.Parse()

	if *email == "" || *apiKey == "" || *phone == "" || *message == "" {
		fmt.Println("All parameters are required")
		flag.Usage()
		os.Exit(1)
	}

	httpClient := &http.Client{}
	ctx := context.Background()

	phoneInt, err := strconv.Atoi(*phone)
	if err != nil {
		fmt.Printf("Error converting phone number to integer: %v\n", err)
		os.Exit(1)
	}

	var logger *log.Logger
	if *logging {
		logger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	} else {
		logger = log.New(io.Discard, "", 0)
	}

	client := smsaero_golang.NewSmsAeroClient(
		*email,
		*apiKey,
		smsaero_golang.WithHTTPClient(httpClient),
		smsaero_golang.WithContext(ctx),
		smsaero_golang.WithLogger(logger),
		smsaero_golang.WithTimeout(time.Second*10),
		smsaero_golang.WithPhoneValidation(*validate),
		smsaero_golang.WithTest(*testMode),
		smsaero_golang.WithSign(*sign),
	)

	if auth, err := client.IsAuthorized(); err != nil {
		fmt.Printf("An error occurred: %v\n", err)
		os.Exit(1)
	} else if !auth {
		fmt.Println("Authorization failed")
		os.Exit(1)
	}

	if !*testMode {
		if balance, err := client.Balance(); err != nil {
			fmt.Printf("An error occurred: %v\n", err)
			os.Exit(1)
		} else if balance < 5 {
			fmt.Println("Insufficient balance")
			os.Exit(1)
		}
	}

	send, err := client.SendSms(phoneInt, *message)
	if err != nil {
		fmt.Printf("An error occurred: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Message sent successfully: %d\n", send.Id)

	time.Sleep(time.Second * 3)

	status, err := client.SmsStatus(send.Id)
	if err != nil {
		fmt.Printf("An error occurred: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Message status: %s\n", status.ExtendStatus)
}
