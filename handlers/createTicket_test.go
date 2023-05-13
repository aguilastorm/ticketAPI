package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
	"github.com/aguilastorm/ticketAPI/models"
)

func TestCreateTicket(t *testing.T) {
	ticket := &models.Ticket{
		ID:           "2",
		User:         "Test User 2",
		CreationDate: time.Now(),
		UpdateDate:   time.Now(),
		Status:       "open",
	}

	jsonTicket, _ := json.Marshal(ticket)
	req, err := http.NewRequest("POST", "/tickets", bytes.NewBuffer(jsonTicket))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateTicket)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
