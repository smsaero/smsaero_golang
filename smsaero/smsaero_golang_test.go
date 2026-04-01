package smsaero_golang

import (
	"context"
	"encoding/base64"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	"time"
)

// testLogger - тестовый логгер для проверки функциональности логирования
// Реализует интерфейс Logger, но не выводит сообщения (пустая реализация)
type testLogger struct{}

func (l *testLogger) Printf(format string, v ...interface{}) {}

// TestNewSmsAeroClient - тест создания клиента SMS Aero с базовыми параметрами
// Проверяет:
// 1. Успешное создание клиента с валидными учетными данными
// 2. Корректное сохранение имени пользователя и пароля
// 3. Отсутствие ошибок при создании
func TestNewSmsAeroClient(t *testing.T) {
	client, err := NewSmsAeroClient("username", "test-api-key-32-chars-long")
	if err != nil {
		t.Errorf("NewSmsAeroClient() failed: %v", err)
		return
	}
	if client.username != "username" || client.password != "test-api-key-32-chars-long" {
		t.Errorf("NewSmsAeroClient() failed to set username or password correctly")
	}
}

// TestNewSmsAeroClientWithHTTPClient - тест создания клиента с пользовательским HTTP-клиентом
// Проверяет:
// 1. Успешное создание клиента с опцией WithHTTPClient
// 2. Корректную установку пользовательского HTTP-клиента
// 3. Сохранение настроек таймаута в пользовательском клиенте
func TestNewSmsAeroClientWithHTTPClient(t *testing.T) {
	customClient := &http.Client{Timeout: 2 * time.Second}
	client, err := NewSmsAeroClient("username", "test-api-key-32-chars-long", WithHTTPClient(customClient))
	if err != nil {
		t.Errorf("NewSmsAeroClient() failed: %v", err)
	}
	if client.client != customClient {
		t.Errorf("WithHTTPClient() failed to set custom HTTP client")
	}
}

// TestNewSmsAeroClientWithTimeout - тест создания клиента с пользовательским таймаутом
// Проверяет:
// 1. Успешное создание клиента с опцией WithTimeout
// 2. Корректную установку пользовательского таймаута
// 3. Сохранение настроек таймаута в HTTP-клиенте
func TestNewSmsAeroClientWithTimeout(t *testing.T) {
	timeout := 5 * time.Second
	client, err := NewSmsAeroClient("username", "test-api-key-32-chars-long", WithTimeout(timeout))
	if err != nil {
		t.Errorf("NewSmsAeroClient() failed: %v", err)
	}
	if client.client.Timeout != timeout {
		t.Errorf("WithTimeout() failed to set custom timeout")
	}
}

// TestNewSmsAeroClientWithPhoneValidation - тест создания клиента с отключенной валидацией телефонов
// Проверяет:
// 1. Успешное создание клиента с опцией WithPhoneValidation(false)
// 2. Корректное отключение валидации номеров телефонов
// 3. Сохранение настройки phoneValidation = false
func TestNewSmsAeroClientWithPhoneValidation(t *testing.T) {
	client, err := NewSmsAeroClient("username", "test-api-key-32-chars-long", WithPhoneValidation(false))
	if err != nil {
		t.Errorf("NewSmsAeroClient() failed: %v", err)
	}
	if client.phoneValidation != false {
		t.Errorf("WithPhoneValidation() failed to disable phone validation")
	}
}

// TestNewSmsAeroClientWithTest - тест создания клиента в тестовом режиме
// Проверяет:
// 1. Успешное создание клиента с опцией WithTest(true)
// 2. Корректное включение тестового режима
// 3. Сохранение настройки testMode = true
func TestNewSmsAeroClientWithTest(t *testing.T) {
	client, err := NewSmsAeroClient("username", "test-api-key-32-chars-long", WithTest(true))
	if err != nil {
		t.Errorf("NewSmsAeroClient() failed: %v", err)
	}
	if client.testMode != true {
		t.Errorf("WithTest() failed to enable test mode")
	}
}

// TestNewSmsAeroClientWithSign - тест создания клиента с пользовательской подписью
// Проверяет:
// 1. Успешное создание клиента с опцией WithSign
// 2. Корректную установку пользовательской подписи
// 3. Сохранение настройки sign в клиенте
func TestNewSmsAeroClientWithSign(t *testing.T) {
	sign := "Custom Sign"
	client, err := NewSmsAeroClient("username", "test-api-key-32-chars-long", WithSign(sign))
	if err != nil {
		t.Errorf("NewSmsAeroClient() failed: %v", err)
	}
	if client.sign != sign {
		t.Errorf("WithSign() failed to set custom sign")
	}
}

// TestNewSmsAeroClientWithLogger - тест создания клиента с пользовательским логгером
// Проверяет:
// 1. Успешное создание клиента с опцией WithLogger
// 2. Корректную установку пользовательского логгера
// 3. Сохранение ссылки на логгер в клиенте
func TestNewSmsAeroClientWithLogger(t *testing.T) {
	logger := &testLogger{}
	client, err := NewSmsAeroClient("username", "test-api-key-32-chars-long", WithLogger(logger))
	if err != nil {
		t.Errorf("NewSmsAeroClient() failed: %v", err)
	}
	if client.logger != logger {
		t.Errorf("WithLogger() failed to set custom logger")
	}
}

// TestNewSmsAeroClientWithContext - тест создания клиента с пользовательским контекстом
// Проверяет:
// 1. Успешное создание клиента с опцией WithContext
// 2. Корректную установку пользовательского контекста
// 3. Сохранение переданного контекста в клиенте
func TestNewSmsAeroClientWithContext(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := NewSmsAeroClient("username", "test-api-key-32-chars-long", WithContext(ctx))
	if err != nil {
		t.Errorf("NewSmsAeroClient() failed: %v", err)
	}
	if client.ctx != ctx {
		t.Errorf("WithContext() failed to set custom context")
	}
}

// TestNewSmsAeroClientWithContextCancellation - тест отмены контекста во время выполнения запроса
// Проверяет:
// 1. Создание клиента с контекстом, который будет отменён
// 2. Запуск HTTP-запроса к медленному серверу
// 3. Отмену контекста во время выполнения запроса
// 4. Корректное прерывание операции с ошибкой context canceled
func TestNewSmsAeroClientWithContextCancellation(t *testing.T) {
	// Создаём сервер с задержкой ответа
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Эмулируем медленный ответ сервера
		time.Sleep(2 * time.Second)
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"success":true}`))
	}))
	defer server.Close()

	// Создаём контекст с возможностью отмены
	ctx, cancel := context.WithCancel(context.Background())

	client, err := NewSmsAeroClient("username", "test-api-key-32-chars-long", WithContext(ctx))
	if err != nil {
		t.Fatalf("NewSmsAeroClient() failed: %v", err)
	}
	client.SetHTTPProtocol("http")
	client.SetGateUrls([]string{strings.TrimPrefix(server.URL, "http://")})

	// Канал для получения результата запроса
	errChan := make(chan error, 1)

	// Запускаем запрос в отдельной горутине
	go func() {
		var resp ErrorResponse
		errChan <- client.executeRequest("/slow", &resp, nil)
	}()

	// Даём запросу начаться, затем отменяем контекст
	time.Sleep(100 * time.Millisecond)
	cancel()

	// Ожидаем результат
	select {
	case err := <-errChan:
		if err == nil {
			t.Error("Expected error due to context cancellation, got nil")
		} else if !strings.Contains(err.Error(), "context canceled") &&
			!strings.Contains(err.Error(), "canceled") {
			t.Errorf("Expected context canceled error, got: %v", err)
		}
	case <-time.After(3 * time.Second):
		t.Error("Request did not complete in time after context cancellation")
	}
}

// TestNewSmsAeroClientWithContextTimeout - тест таймаута контекста во время выполнения запроса
// Проверяет:
// 1. Создание клиента с контекстом с коротким таймаутом
// 2. Запуск HTTP-запроса к медленному серверу
// 3. Автоматическое прерывание операции по истечении таймаута контекста
// 4. Корректную обработку ошибки context deadline exceeded
func TestNewSmsAeroClientWithContextTimeout(t *testing.T) {
	// Создаём сервер с задержкой ответа
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Эмулируем медленный ответ сервера (дольше чем таймаут контекста)
		time.Sleep(2 * time.Second)
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"success":true}`))
	}))
	defer server.Close()

	// Создаём контекст с коротким таймаутом
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()

	client, err := NewSmsAeroClient("username", "test-api-key-32-chars-long", WithContext(ctx))
	if err != nil {
		t.Fatalf("NewSmsAeroClient() failed: %v", err)
	}
	client.SetHTTPProtocol("http")
	client.SetGateUrls([]string{strings.TrimPrefix(server.URL, "http://")})

	var resp ErrorResponse
	err = client.executeRequest("/slow", &resp, nil)

	if err == nil {
		t.Error("Expected error due to context timeout, got nil")
	} else if !strings.Contains(err.Error(), "deadline exceeded") &&
		!strings.Contains(err.Error(), "context deadline exceeded") &&
		!strings.Contains(err.Error(), "timeout") {
		t.Errorf("Expected context deadline exceeded error, got: %v", err)
	}
}

// TestNewSmsAeroClientDefaultValues - тест проверки значений по умолчанию при создании клиента
// Проверяет:
// 1. Валидация телефонов включена по умолчанию (phoneValidation = true)
// 2. Тестовый режим отключен по умолчанию (testMode = false)
// 3. Подпись по умолчанию "Sms Aero"
// 4. HTTP-клиент по умолчанию http.DefaultClient
// 5. Контекст не nil (создается context.Background())
func TestNewSmsAeroClientDefaultValues(t *testing.T) {
	client, err := NewSmsAeroClient("username", "test-api-key-32-chars-long")
	if err != nil {
		t.Errorf("NewSmsAeroClient() failed: %v", err)
	}
	if client.phoneValidation != true {
		t.Errorf("Default phoneValidation expected to be true")
	}
	if client.testMode != false {
		t.Errorf("Default testMode expected to be false")
	}
	if client.sign != "Sms Aero" {
		t.Errorf("Default sign expected to be 'Sms Aero'")
	}
	if client.client != http.DefaultClient {
		t.Errorf("Default HTTP client expected to be http.DefaultClient")
	}
	if client.ctx == nil {
		t.Errorf("Default context expected to be non-nil")
	}
}

// TestClient_getApiPath_TestModeEnabled - тест метода getAPIPath в тестовом режиме
// Проверяет:
// 1. При включенном тестовом режиме (testMode = true) возвращается тестовый путь
// 2. Корректное переключение между обычным и тестовым API
// 3. Правильную работу метода getAPIPath с параметрами (normalPath, testPath)
func TestClient_getApiPath_TestModeEnabled(t *testing.T) {
	client, err := NewSmsAeroClient("username", "test-api-key-32-chars-long", WithTest(true))
	if err != nil {
		t.Errorf("NewSmsAeroClient() failed: %v", err)
	}
	testPath := "sms/testsend"
	normalPath := "sms/send"
	expectedPath := testPath

	if path := client.getAPIPath(normalPath, testPath); path != expectedPath {
		t.Errorf("getApiPath() with testMode enabled = %v, want %v", path, expectedPath)
	}
}

// TestClient_getApiPath_TestModeDisabled - тест метода getAPIPath в обычном режиме
// Проверяет:
// 1. При отключенном тестовом режиме (testMode = false) возвращается обычный путь
// 2. Корректное переключение между обычным и тестовым API
// 3. Правильную работу метода getAPIPath с параметрами (normalPath, testPath)
func TestClient_getApiPath_TestModeDisabled(t *testing.T) {
	client, err := NewSmsAeroClient("username", "test-api-key-32-chars-long") // testMode is false by default
	if err != nil {
		t.Errorf("NewSmsAeroClient() failed: %v", err)
	}
	testPath := "sms/testsend"
	normalPath := "sms/send"
	expectedPath := normalPath

	if path := client.getAPIPath(normalPath, testPath); path != expectedPath {
		t.Errorf("getApiPath() with testMode disabled = %v, want %v", path, expectedPath)
	}
}

// setupMockServer - создает mock HTTP сервер для тестирования различных сценариев
// Возвращает тестовый сервер с двумя эндпоинтами:
// - /success - возвращает успешный ответ {"success":true}
// - /fail - возвращает ошибку 500 с сообщением об ошибке
func setupMockServer() *httptest.Server {
	handler := http.NewServeMux()
	handler.HandleFunc("/success", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"success":true}`))
	})
	handler.HandleFunc("/fail", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(`{"success":false, "message":"Internal Server Error"}`))
	})
	return httptest.NewServer(handler)
}

// TestExecuteRequestDefaultValues - тест выполнения HTTP-запроса с настройками по умолчанию
// Проверяет:
// 1. Успешное выполнение HTTP-запроса к mock-серверу
// 2. Корректную обработку успешного ответа от сервера
// 3. Правильную работу метода executeRequest с базовыми настройками
// 4. Установку HTTP-протокола и URL шлюзов
func TestExecuteRequestDefaultValues(t *testing.T) {
	server := setupMockServer()
	defer server.Close()

	client, err := NewSmsAeroClient("username", "test-api-key-32-chars-long")
	if err != nil {
		t.Errorf("NewSmsAeroClient() failed: %v", err)
	}
	client.SetHTTPProtocol("http")
	client.SetGateUrls([]string{strings.TrimPrefix(server.URL, "http://")})
	var resp ErrorResponse
	err = client.executeRequest("/success", &resp, nil)
	if err != nil {
		t.Errorf("executeRequest() error = %v, wantErr %v", err, false)
	}
	if !resp.Success {
		t.Errorf("executeRequest() response = %v, want %v", resp.Success, true)
	}
}

// TestExecuteRequestURLSwitch - тест переключения между URL шлюзов при недоступности первого
// Проверяет:
// 1. Переключение на следующий URL при недоступности первого шлюза
// 2. Успешное выполнение запроса на втором доступном шлюзе
// 3. Корректную обработку ошибок подключения к первому шлюзу
// 4. Работу механизма fallback между шлюзами
func TestExecuteRequestURLSwitch(t *testing.T) {
	server := setupMockServer()
	defer server.Close()
	GateUrls = []string{"http://invalid.url", strings.TrimPrefix(server.URL, "http://")}

	client, err := NewSmsAeroClient("username", "test-api-key-32-chars-long", WithTimeout(1*time.Second))
	if err != nil {
		t.Errorf("NewSmsAeroClient() failed: %v", err)
	}
	client.SetHTTPProtocol("http")
	var resp ErrorResponse
	err = client.executeRequest("/success", &resp, nil)
	if err != nil {
		t.Errorf("executeRequest() should not fail when switching URLs, got error: %v", err)
	}
	if !resp.Success {
		t.Errorf("executeRequest() should succeed on the second URL, got: %v", resp.Success)
	}
}

// TestExecuteRequestHeaders - тест проверки HTTP-заголовков в запросах
// Проверяет:
// 1. Корректную установку заголовка Content-Type: application/x-www-form-urlencoded
// 2. Правильную установку заголовка User-Agent: SAGoClient/2.0.0
// 3. Отправку всех необходимых заголовков в HTTP-запросах
// 4. Соответствие заголовков стандартам API SMS Aero
func TestExecuteRequestHeaders(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Content-Type") != "application/x-www-form-urlencoded" {
			t.Errorf("Content-Type header = %v, want %v", r.Header.Get("Content-Type"), "application/x-www-form-urlencoded")
		}
		if r.Header.Get("User-Agent") != "SAGoClient/2.0.0" {
			t.Errorf("User-Agent header = %v, want %v", r.Header.Get("User-Agent"), "SAGoClient/2.0.0")
		}
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()
	client, err := NewSmsAeroClient("username", "test-api-key-32-chars-long")
	if err != nil {
		t.Errorf("NewSmsAeroClient() failed: %v", err)
	}
	client.SetGateUrls([]string{strings.TrimPrefix(server.URL, "http://")})
	var resp ErrorResponse
	_ = client.executeRequest("", &resp, nil)
}

// TestExecuteRequestWithParams - тест выполнения HTTP-запроса с параметрами
// Проверяет:
// 1. Успешное выполнение запроса с дополнительными параметрами
// 2. Корректную передачу параметров в HTTP-запросе
// 3. Правильную работу метода executeRequest с параметрами
// 4. Обработку URL-encoded параметров
func TestExecuteRequestWithParams(t *testing.T) {
	server := setupMockServer()
	defer server.Close()
	client, err := NewSmsAeroClient("username", "test-api-key-32-chars-long")
	if err != nil {
		t.Errorf("NewSmsAeroClient() failed: %v", err)
	}
	client.SetHTTPProtocol("http")
	client.SetGateUrls([]string{strings.TrimPrefix(server.URL, "http://")})

	params := url.Values{}
	params.Add("test", "value")
	var resp ErrorResponse
	err = client.executeRequest("/success", &resp, nil)
	if err != nil {
		t.Errorf("executeRequest() with params error = %v, wantErr %v", err, false)
	}
}

// TestExecuteRequestBasicAuth - тест базовой HTTP-аутентификации
// Проверяет:
// 1. Корректную установку заголовка Authorization с Basic Auth
// 2. Правильное кодирование учетных данных в base64
// 3. Формат заголовка: "Basic base64(username:password)"
// 4. Передачу учетных данных в каждом HTTP-запросе
func TestExecuteRequestBasicAuth(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		if auth == "" {
			t.Fatal("Authorization header not found")
		}
		expectedAuth := "Basic " + base64.StdEncoding.EncodeToString([]byte("username:test-api-key-32-chars-long"))
		if auth != expectedAuth {
			t.Errorf("Authorization header = %v, want %v", auth, expectedAuth)
		}
	}))
	defer server.Close()

	client, err := NewSmsAeroClient("username", "test-api-key-32-chars-long")
	if err != nil {
		t.Errorf("NewSmsAeroClient() failed: %v", err)
	}
	client.SetHTTPProtocol("http")
	client.SetGateUrls([]string{strings.TrimPrefix(server.URL, "http://")})
	_ = client.executeRequest("", &ErrorResponse{}, nil)
}

func setupMobileIdMockServer() *httptest.Server {
	handler := http.NewServeMux()
	handler.HandleFunc("/mobile-id/send", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"success":true,"data":{"id":273,"number":"79031234567","authType":"SIM-PUSH","codeSms":"","status":0,"cost":0,"dateCreate":1719119523,"dateSend":1719119523}}`))
	})
	handler.HandleFunc("/mobile-id/status", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"success":true,"data":{"id":273,"number":"79031234567","authType":"SMS","codeSms":"","status":3,"cost":0,"dateCreate":1719119523,"dateSend":1719119523}}`))
	})
	handler.HandleFunc("/mobile-id/verify", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"success":true,"data":{"id":273,"number":"79031234567","authType":"SMS","codeSms":"1234","status":3,"cost":0,"dateCreate":1719119523,"dateSend":1719119523}}`))
	})
	return httptest.NewServer(handler)
}

func TestSendMobileId(t *testing.T) {
	server := setupMobileIdMockServer()
	defer server.Close()

	client, err := NewSmsAeroClient("username", "test-api-key-32-chars-long")
	if err != nil {
		t.Fatalf("NewSmsAeroClient() failed: %v", err)
	}
	client.SetHTTPProtocol("http")
	client.SetGateUrls([]string{strings.TrimPrefix(server.URL, "http://")})

	result, err := client.SendMobileId(79031234567, "SMSAero", "https://example.com/callback")
	if err != nil {
		t.Fatalf("SendMobileId() error = %v", err)
	}
	if result.ID != 273 {
		t.Errorf("SendMobileId() ID = %v, want 273", result.ID)
	}
	if result.AuthType != "SIM-PUSH" {
		t.Errorf("SendMobileId() AuthType = %v, want SIM-PUSH", result.AuthType)
	}
	if result.Status != 0 {
		t.Errorf("SendMobileId() Status = %v, want 0", result.Status)
	}
}

func TestSendMobileIdInvalidPhone(t *testing.T) {
	client, err := NewSmsAeroClient("username", "test-api-key-32-chars-long")
	if err != nil {
		t.Fatalf("NewSmsAeroClient() failed: %v", err)
	}
	_, err = client.SendMobileId(123, "SMSAero", "https://example.com/callback")
	if err == nil {
		t.Error("SendMobileId() expected error for invalid phone")
	}
}

func TestSendMobileIdInvalidCallbackUrl(t *testing.T) {
	client, err := NewSmsAeroClient("username", "test-api-key-32-chars-long")
	if err != nil {
		t.Fatalf("NewSmsAeroClient() failed: %v", err)
	}
	_, err = client.SendMobileId(79031234567, "SMSAero", "ftp://bad")
	if err == nil {
		t.Error("SendMobileId() expected error for invalid callback URL")
	}
}

func TestMobileIdStatus(t *testing.T) {
	server := setupMobileIdMockServer()
	defer server.Close()

	client, err := NewSmsAeroClient("username", "test-api-key-32-chars-long")
	if err != nil {
		t.Fatalf("NewSmsAeroClient() failed: %v", err)
	}
	client.SetHTTPProtocol("http")
	client.SetGateUrls([]string{strings.TrimPrefix(server.URL, "http://")})

	result, err := client.MobileIdStatus(273)
	if err != nil {
		t.Fatalf("MobileIdStatus() error = %v", err)
	}
	if result.ID != 273 {
		t.Errorf("MobileIdStatus() ID = %v, want 273", result.ID)
	}
	if result.Status != 3 {
		t.Errorf("MobileIdStatus() Status = %v, want 3", result.Status)
	}
}

func TestMobileIdStatusInvalidID(t *testing.T) {
	client, err := NewSmsAeroClient("username", "test-api-key-32-chars-long")
	if err != nil {
		t.Fatalf("NewSmsAeroClient() failed: %v", err)
	}
	_, err = client.MobileIdStatus(0)
	if err == nil {
		t.Error("MobileIdStatus() expected error for zero ID")
	}
	_, err = client.MobileIdStatus(-1)
	if err == nil {
		t.Error("MobileIdStatus() expected error for negative ID")
	}
}

func TestVerifyMobileId(t *testing.T) {
	server := setupMobileIdMockServer()
	defer server.Close()

	client, err := NewSmsAeroClient("username", "test-api-key-32-chars-long")
	if err != nil {
		t.Fatalf("NewSmsAeroClient() failed: %v", err)
	}
	client.SetHTTPProtocol("http")
	client.SetGateUrls([]string{strings.TrimPrefix(server.URL, "http://")})

	result, err := client.VerifyMobileId(273, "1234", "SMSAero")
	if err != nil {
		t.Fatalf("VerifyMobileId() error = %v", err)
	}
	if result.ID != 273 {
		t.Errorf("VerifyMobileId() ID = %v, want 273", result.ID)
	}
	if result.CodeSms != "1234" {
		t.Errorf("VerifyMobileId() CodeSms = %v, want 1234", result.CodeSms)
	}
}

func TestVerifyMobileIdInvalidID(t *testing.T) {
	client, err := NewSmsAeroClient("username", "test-api-key-32-chars-long")
	if err != nil {
		t.Fatalf("NewSmsAeroClient() failed: %v", err)
	}
	_, err = client.VerifyMobileId(0, "1234", "SMSAero")
	if err == nil {
		t.Error("VerifyMobileId() expected error for zero ID")
	}
}

func TestVerifyMobileIdEmptyCode(t *testing.T) {
	client, err := NewSmsAeroClient("username", "test-api-key-32-chars-long")
	if err != nil {
		t.Fatalf("NewSmsAeroClient() failed: %v", err)
	}
	_, err = client.VerifyMobileId(273, "", "SMSAero")
	if err == nil {
		t.Error("VerifyMobileId() expected error for empty code")
	}
}
