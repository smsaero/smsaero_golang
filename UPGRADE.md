# UPGRADE.md

## Миграция с v1.x на v2.0

Версия 2.0 содержит **breaking changes**, требующие обновления вашего кода.

### 1. Изменение пути импорта

```go
// v1.x
import "github.com/smsaero/smsaero_golang/smsaero"

// v2.0
import smsaero "github.com/smsaero/smsaero_golang/v2/smsaero"
```

Обновите `go.mod`:

```bash
go get github.com/smsaero/smsaero_golang/v2@latest
go mod tidy
```

### 2. Изменение конструктора клиента

**Breaking Change:** `NewSmsAeroClient` теперь возвращает `(*Client, error)` вместо `*Client`.

```go
// v1.x — БЕЗ обработки ошибки
client := smsaero.NewSmsAeroClient(email, apiKey)

// v2.0 — ОБЯЗАТЕЛЬНАЯ обработка ошибки
client, err := smsaero.NewSmsAeroClient(email, apiKey)
if err != nil {
    log.Fatal("Ошибка создания клиента:", err)
}
```

Конструктор теперь валидирует учетные данные и возвращает ошибку при некорректных значениях.

### 3. Изменение метода ViberSend

**Breaking Change:** Сигнатура метода полностью изменена на functional options pattern.

```go
// v1.x — 13 позиционных параметров
result, err := client.ViberSend(
    sign, channel, text, number, groupId,
    imageSource, textButton, linkButton, dateSend,
    signSms, channelSms, textSms, priceSms,
)

// v2.0 — 3 обязательных + опции
result, err := client.ViberSend(
    sign, channel, text,
    smsaero.WithViberSendNumber("79031234567"),
    smsaero.WithViberSendGroupID(123),
    smsaero.WithViberSendImageSource("https://example.com/image.png"),
    smsaero.WithViberSendTextButton("Подробнее"),
    smsaero.WithViberSendLinkButton("https://example.com"),
)
```

### 4. Изменение метода ContactAdd

**Breaking Change:** Сигнатура изменена на functional options pattern.

```go
// v1.x — 2 позиционных параметра
result, err := client.ContactAdd(number, groupId)

// v2.0 — 1 обязательный + опции
result, err := client.ContactAdd(
    number,
    smsaero.WithContactAddGroupID(123),
    smsaero.WithContactAddFirstName("Иван"),
    smsaero.WithContactAddLastName("Иванов"),
    smsaero.WithContactAddBirthday("1990-01-15"),
)
```

### 5. Новая система обработки ошибок

В v2.0 добавлены типизированные ошибки для более точной обработки:

```go
result, err := client.SendSms(79031234567, "Привет")
if err != nil {
    switch e := err.(type) {
    case *smsaero.SmsAeroValidationError:
        // Ошибка валидации входных данных
        fmt.Printf("Валидация: %s\n", e.Error())
    case *smsaero.SmsAeroNoMoneyError:
        // Недостаточно средств на балансе
        fmt.Printf("Баланс: %.2f\n", e.Balance)
    case *smsaero.SmsAeroAPIError:
        // Ошибка API (код ответа, сообщение)
        fmt.Printf("API ошибка [%d]: %s\n", e.StatusCode, e.Error())
    case *smsaero.SmsAeroConnectionError:
        // Ошибка соединения с сервером
        fmt.Printf("Соединение: %s\n", e.Error())
    default:
        fmt.Printf("Неизвестная ошибка: %s\n", err.Error())
    }
}
```

### 6. Новые возможности v2.0

- **Валидация входных данных** — автоматическая проверка номеров телефонов, email, API ключей
- **Типизированные ошибки** — точная обработка различных типов ошибок
- **Functional Options** — гибкая конфигурация методов
- **Улучшенная документация** — комментарии ко всем публичным методам

### Чек-лист миграции

- [ ] Обновить путь импорта на `/v2/smsaero`
- [ ] Добавить обработку ошибки в `NewSmsAeroClient`
- [ ] Обновить вызовы `ViberSend` (если используется)
- [ ] Обновить вызовы `ContactAdd` (если используется)
- [ ] Обновить обработку ошибок (опционально, но рекомендуется)
- [ ] Запустить тесты: `go test ./...`

### Поддержка v1.x

Версия v1.x остаётся доступной по адресу:

```go
import "github.com/smsaero/smsaero_golang/smsaero"
```

Однако новые функции и исправления будут выпускаться только для v2.x.
