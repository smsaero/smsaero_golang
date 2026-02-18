# SmsAero GoLang Api клиент v2

Библиотека для отправки SMS сообщений с использованием SmsAero API.

[![Go Version](https://img.shields.io/badge/go-1.21+-blue.svg)](https://golang.org/)
[![License](https://img.shields.io/badge/license-MIT-green.svg)](LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/smsaero/smsaero_golang/v2)](https://goreportcard.com/report/github.com/smsaero/smsaero_golang/v2)

## Возможности

- Полная поддержка SMS Aero API
- Валидация входных данных и обработка ошибок
- Тестовый режим для безопасной разработки
- Отправка SMS, Telegram и Viber сообщений
- Управление контактами и группами
- HLR проверка номеров
- Управление балансом

## Установка

```bash
go mod init smsaero-example
go get github.com/smsaero/smsaero_golang/v2/smsaero@latest
```

## Пример использования

Получите учетные данные на странице настроек аккаунта: https://smsaero.ru/cabinet/settings/apikey/

```go
package main

import (
    "fmt"
    "log"
    "time"

    smsaero "github.com/smsaero/smsaero_golang/v2/smsaero"
)

const (
    Email  = "ваш email"
    ApiKey = "ваш api ключ"
)

func main() {
    // Создание клиента
    client, err := smsaero.NewSmsAeroClient(
        Email, ApiKey,
        smsaero.WithTimeout(time.Second*10),
        smsaero.WithTest(true),
        smsaero.WithPhoneValidation(true),
    )
    if err != nil {
        log.Fatal("Ошибка создания клиента:", err)
    }

    // Отправка SMS
    if sendResult, err := client.SendSms(70000000000, "Привет, Мир!"); err == nil {
        fmt.Printf("SMS отправлено, ID: %d\n", sendResult.ID)
    } else {
        log.Fatal(err)
    }

    // Отправка Telegram кода
    telegramResult, err := client.SendTelegram(
        70000000000,
        1234,
        smsaero.WithSendTelegramSign("SMS Aero"),
        smsaero.WithSendTelegramText("Ваш код: 1234"),
    )
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Telegram ID: %d, Статус: %s\n", telegramResult.ID, telegramResult.ExtendStatus)
}
```

## Конфигурация клиента

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

## Доступные методы

### SMS
- `SendSms(phone int, text string, options ...SendSmsOption) (*SendSms, error)`

### Telegram
- `SendTelegram(phone int, code int, options ...SendTelegramOption) (*SendTelegram, error)`

### Viber
- `ViberSend(sign, channel, text string, options ...ViberSendOption) (interface{}, error)`

### Контакты
- `ContactAdd(phone string, options ...ContactAddOption) (*ContactAdd, error)`
- `ContactList() (interface{}, error)`
- `ContactDelete(phone string) (interface{}, error)`

### Группы
- `GroupAdd(name string) (*GroupAdd, error)`
- `GroupList() (interface{}, error)`
- `GroupDelete(id int) (interface{}, error)`

### Баланс
- `Balance() (float64, error)`
- `AddBalance(amount float64, cardID string) (interface{}, error)`

### HLR
- `HlrCheck(phone int) (*HlrCheck, error)`
- `HlrStatus(id int) (interface{}, error)`

### Другое
- `IsAuthorized() (bool, error)`
- `SignList() (interface{}, error)`
- `Tariffs() (interface{}, error)`

## Обработка ошибок

```go
result, err := client.SendSms(79031234567, "Привет")
if err != nil {
    switch e := err.(type) {
    case *smsaero.SmsAeroValidationError:
        fmt.Printf("Ошибка валидации: %s\n", e.Error())
    case *smsaero.SmsAeroNoMoneyError:
        fmt.Printf("Недостаточно средств: %s\n", e.Error())
    case *smsaero.SmsAeroAPIError:
        fmt.Printf("Ошибка API: %s\n", e.Error())
    case *smsaero.SmsAeroConnectionError:
        fmt.Printf("Ошибка соединения: %s\n", e.Error())
    default:
        fmt.Printf("Неизвестная ошибка: %s\n", e.Error())
    }
}
```

## Установка зависимостей

```bash
go get github.com/smsaero/smsaero_golang/v2@latest
go mod tidy
```

## Запуск в Docker

```bash
docker pull 'smsaero/smsaero_golang:latest'
docker run -it --rm 'smsaero/smsaero_golang:latest' smsaero_send \
    -email "ваш email" \
    -api_key "ваш api ключ" \
    -phone 70000000000 \
    -message 'Привет, Мир!' \
    -test
```

## Лицензия

MIT License
