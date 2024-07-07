package smsaero_golang

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	"time"
)

type testLogger struct{}

func (l *testLogger) Printf(format string, v ...interface{}) {}

func TestNewSmsAeroClient(t *testing.T) {
	client := NewSmsAeroClient("username", "password")
	if client.username != "username" || client.password != "password" {
		t.Errorf("NewSmsAeroClient() failed to set username or password correctly")
	}
}

func TestNewSmsAeroClientWithHTTPClient(t *testing.T) {
	customClient := &http.Client{Timeout: 2 * time.Second}
	client := NewSmsAeroClient("username", "password", WithHTTPClient(customClient))
	if client.client != customClient {
		t.Errorf("WithHTTPClient() failed to set custom HTTP client")
	}
}

func TestNewSmsAeroClientWithContext(t *testing.T) {
	ctx := http.DefaultClient.Timeout
	client := NewSmsAeroClient("username", "password", WithTimeout(ctx))
	if client.client.Timeout != ctx {
		t.Errorf("WithContext() failed to set custom context")
	}
}

func TestNewSmsAeroClientWithTimeout(t *testing.T) {
	timeout := 5 * time.Second
	client := NewSmsAeroClient("username", "password", WithTimeout(timeout))
	if client.client.Timeout != timeout {
		t.Errorf("WithTimeout() failed to set custom timeout")
	}
}

func TestNewSmsAeroClientWithPhoneValidation(t *testing.T) {
	client := NewSmsAeroClient("username", "password", WithPhoneValidation(false))
	if client.phoneValidation != false {
		t.Errorf("WithPhoneValidation() failed to disable phone validation")
	}
}

func TestNewSmsAeroClientWithTest(t *testing.T) {
	client := NewSmsAeroClient("username", "password", WithTest(true))
	if client.testMode != true {
		t.Errorf("WithTest() failed to enable test mode")
	}
}

func TestNewSmsAeroClientWithSign(t *testing.T) {
	sign := "Custom Sign"
	client := NewSmsAeroClient("username", "password", WithSign(sign))
	if client.sign != sign {
		t.Errorf("WithSign() failed to set custom sign")
	}
}

func TestNewSmsAeroClientWithLogger(t *testing.T) {
	logger := &testLogger{}
	client := NewSmsAeroClient("username", "password", WithLogger(logger))
	if client.logger != logger {
		t.Errorf("WithLogger() failed to set custom logger")
	}
}

func TestNewSmsAeroClientDefaultValues(t *testing.T) {
	client := NewSmsAeroClient("username", "password")
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

func TestClient_getApiPath_TestModeEnabled(t *testing.T) {
	client := NewSmsAeroClient("username", "password", WithTest(true))
	testPath := "sms/testsend"
	normalPath := "sms/send"
	expectedPath := testPath

	if path := client.getApiPath(normalPath, testPath); path != expectedPath {
		t.Errorf("getApiPath() with testMode enabled = %v, want %v", path, expectedPath)
	}
}

func TestClient_getApiPath_TestModeDisabled(t *testing.T) {
	client := NewSmsAeroClient("username", "password") // testMode is false by default
	testPath := "sms/testsend"
	normalPath := "sms/send"
	expectedPath := normalPath

	if path := client.getApiPath(normalPath, testPath); path != expectedPath {
		t.Errorf("getApiPath() with testMode disabled = %v, want %v", path, expectedPath)
	}
}

// Mock HTTP server to simulate responses for different scenarios
func setupMockServer() *httptest.Server {
	handler := http.NewServeMux()
	handler.HandleFunc("/v2/success", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"success":true}`))
	})
	handler.HandleFunc("/v2/fail", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(`{"success":false, "message":"Internal Server Error"}`))
	})
	return httptest.NewServer(handler)
}

func TestExecuteRequestDefaultValues(t *testing.T) {
	server := setupMockServer()
	defer server.Close()

	GateUrls = []string{strings.TrimPrefix(server.URL, "http://")}
	fmt.Println(server.URL)

	client := NewSmsAeroClient("username", "password")
	client.httpProtocol = "http"
	var resp ErrorResponse
	err := client.executeRequest("/v2/success", &resp, nil)
	if err != nil {
		t.Errorf("executeRequest() error = %v, wantErr %v", err, false)
	}
	if !resp.Success {
		t.Errorf("executeRequest() response = %v, want %v", resp.Success, true)
	}
}

func TestExecuteRequestURLSwitch(t *testing.T) {
	server := setupMockServer()
	defer server.Close()
	GateUrls = []string{"http://invalid.url", strings.TrimPrefix(server.URL, "http://")}

	client := NewSmsAeroClient("username", "password", WithTimeout(1*time.Second))
	client.httpProtocol = "http"
	var resp ErrorResponse
	err := client.executeRequest("/v2/success", &resp, nil)
	if err != nil {
		t.Errorf("executeRequest() should not fail when switching URLs, got error: %v", err)
	}
	if !resp.Success {
		t.Errorf("executeRequest() should succeed on the second URL, got: %v", resp.Success)
	}
}

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
	GateUrls = []string{strings.TrimPrefix(server.URL, "http://")}

	client := NewSmsAeroClient("username", "password")
	var resp ErrorResponse
	_ = client.executeRequest("", &resp, nil)
}

func TestExecuteRequestWithParams(t *testing.T) {
	server := setupMockServer()
	defer server.Close()
	GateUrls = []string{strings.TrimPrefix(server.URL, "http://")}

	client := NewSmsAeroClient("username", "password")
	client.httpProtocol = "http"

	params := url.Values{}
	params.Add("test", "value")
	var resp ErrorResponse
	err := client.executeRequest("/v2/success", &resp, nil)
	if err != nil {
		t.Errorf("executeRequest() with params error = %v, wantErr %v", err, false)
	}
}

func TestExecuteRequestBasicAuth(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		if auth == "" {
			t.Fatal("Authorization header not found")
		}
		expectedAuth := "Basic " + base64.StdEncoding.EncodeToString([]byte("username:password"))
		if auth != expectedAuth {
			t.Errorf("Authorization header = %v, want %v", auth, expectedAuth)
		}
	}))
	defer server.Close()

	GateUrls = []string{strings.TrimPrefix(server.URL, "http://")}
	client := NewSmsAeroClient("username", "password")
	client.httpProtocol = "http"
	_ = client.executeRequest("", &ErrorResponse{}, nil)
}
