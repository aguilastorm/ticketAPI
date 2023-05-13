package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
	"github.com/gorilla/mux"
	"github.com/aguilastorm/ticketAPI/models"
	"github.com/aguilastorm/ticketAPI/store"
	"github.com/aguilastorm/ticketAPI/handlers"
)

func TestGetTicket(t *testing.T) {
	tickets = append(store.GetTickets(), models.Ticket{ID: "1", User: "Test User", CreationDate: time.Now(), UpdateDate: time.Now(), Status: "open"})
	req, err := http.NewRequest("GET", "/tickets/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/tickets/{id}", handlers.GetTicket)
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
