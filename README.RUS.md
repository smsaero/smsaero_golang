# SmsAero GoLang Api клиент

Библиотека для отправки SMS сообщений с использованием SmsAero API. Написана на GoLang.

## Установка:

```bash
go get github.com/smsaero/smsaero_golang
```

Если вы начинаете с нуля в новой директории, сначала необходимо создать файл `go.mod` для отслеживания зависимостей:

```bash
go mod init smsaero-example
```

## Пример использования:

Получите учетные данные на странице настроек аккаунта: https://smsaero.ru/cabinet/settings/apikey/

```golang
package main

import (
    "fmt"
    "time"
    "github.com/smsaero/smsaero_golang/smsaero"
)

const (
    Email  = "ваш email"
    ApiKey = "ваш api ключ"
)

func main() {
    client := smsaero_golang.NewSmsAeroClient(
        Email, ApiKey,
        smsaero_golang.WithTimeout(time.Second*10),
        smsaero_golang.WithTest(true),
        smsaero_golang.WithPhoneValidation(false),
    )

    // Отправка SMS
    if sendResult, err := client.SendSms(70000000000, "Привет, Мир!"); err == nil {
        fmt.Println(sendResult.Id)
    } else {
        panic(err)
    }

    // Отправка Telegram кода
    if telegramResult, err := client.SendTelegram(70000000000, 1234, 
        smsaero_golang.WithSendTelegramSign("SMS Aero"),
        smsaero_golang.WithSendTelegramText("Ваш код: 1234")); err == nil {
        fmt.Printf("Telegram ID: %d, Status: %s\n", telegramResult.Id, telegramResult.ExtendStatus)
    } else {
        panic(err)
    }
}
```

## Установка зависимостей:

```bash
go get github.com/smsaero/smsaero_golang@latest
go mod tidy
```

## Запуск в Docker:

```bash
docker pull 'smsaero/smsaero_golang:latest'
docker run -it --rm 'smsaero/smsaero_golang:latest' smsaero_send -email "ваш email" -api_key "ваш api ключ" -phone 70000000000 -message 'Привет, Мир!' -test
```

## Лицензия

```
MIT License
```
