package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
	"github.com/gorilla/mux"
	"github.com/aguilastorm/ticketAPI/models"
)

func TestDeleteTicket(t *testing.T) {
	tickets = append(tickets, models.Ticket{ID: "3", User: "Test User 3", CreationDate: time.Now(), UpdateDate: time.Now(), Status: "open"})
	req, err := http.NewRequest("DELETE", "/tickets/3", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/tickets/{id}", DeleteTicket)
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
