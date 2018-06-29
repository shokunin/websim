package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/shokunin/websim/handlers"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthCheck(t *testing.T) {
	type ApiMessage struct {
		Message string
	}
	w := httptest.NewRecorder()
	r := gin.Default()
	r.GET("/health", websim.HealthCheck)
	req, _ := http.NewRequest("GET", "/health", nil)
	r.ServeHTTP(w, req)
	if w.Code != 200 {
		t.Errorf("HTTP response code should be 200, was: %d", w.Code)
	}
	var message ApiMessage
	err := json.Unmarshal(w.Body.Bytes(), &message)
	if err != nil {
		t.Errorf("JSON parse error: %d", w.Code)
	}

	if message.Message != "OK" {
		t.Errorf("JSON message should be OK, but is: %s", message.Message)
	}
}
