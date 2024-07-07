# UPGRADE.md

## Обновление с предыдущей версии до текущей

Этот документ содержит инструкции по обновлению с предыдущей версии `smsaero_golang` до текущей версии.

### Изменения в импорте пакета

В новой версии изменился путь импорта пакета. Ранее пакет импортировался как `github.com/smsaero/smsaero_golang`, теперь же необходимо использовать `github.com/smsaero/smsaero_golang/smsaero`.

#### Как обновить:

- Во всех файлах вашего проекта, где используется пакет `smsaero_golang`, обновите путь импорта с `github.com/smsaero/smsaero_golang` на `github.com/smsaero/smsaero_golang/smsaero`.

#### Пример обновления:

```go
// Старый импорт
import "github.com/smsaero/smsaero_golang"

// Новый импорт
import "github.com/smsaero/smsaero_golang/smsaero"
```



### Изменения в API

1. **Добавление контекста в клиентские запросы**

В новой версии была добавлена возможность передачи `context.Context` в запросы к API. Это позволяет более гибко управлять таймаутами и отменой запросов.

#### Как обновить:

- Для использования контекста в запросах, теперь необходимо использовать функциональные опции при создании клиента. Пример использования:

  ```go
  httpClient := &http.Client{Timeout: time.Second * 10}
  ctx := context.Background()

  client := smsaero_golang.NewSmsAeroClient(
      "your_email@example.com",
      "your_api_key",
      smsaero_golang.WithHTTPClient(httpClient),
      smsaero_golang.WithContext(ctx),
  )
  ```

2. **Добавление необязательного параметра `sign` в функцию `SendSms`**

В новой версии была добавлена возможность указывать необязательный параметр `sign` при отправке SMS. Если параметр не указан, используется значение по умолчанию.

#### Пример использования:

```go
client := smsaero_golang.NewSmsAeroClient(
  "your_email@example.com",
  "your_api_key",
  smsaero_golang.WithSign("YourCustomSign"),
)

// Или напрямую при отправке SMS
send, err := client.SendSms(phoneInt, *message, smsaero_golang.WithSign("YourCustomSign"))
```
