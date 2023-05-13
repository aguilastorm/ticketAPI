package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetTickets(t *testing.T) {
	req, err := http.NewRequest("GET", "/tickets", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetTickets)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
