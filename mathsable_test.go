package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHome(t *testing.T) {
	req, _ := http.NewRequest("GET", "http://localhost:8080/", nil)
	w := httptest.NewRecorder()

	handler(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("Home page didn't return %v", http.StatusOK)
	}
}
