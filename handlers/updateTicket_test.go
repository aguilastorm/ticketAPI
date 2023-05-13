package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
	"github.com/gorilla/mux"
	"github.com/aguilastorm/ticketAPI/models"
	"github.com/aguilastorm/ticketAPI/store"
	"github.com/aguilastorm/ticketAPI/handlers"

)

func TestUpdateTicket(t *testing.T) {
	store.AddTicket(models.Ticket{ID: "4", User: "Test User 4", CreationDate: time.Now(), UpdateDate: time.Now(), Status: "open"})

	newTicket := &models.Ticket{
		ID:           "4",
		User:         "Updated User",
		CreationDate: time.Now(),
		UpdateDate:   time.Now(),
		Status:       "closed",
	}

	jsonTicket, err := json.Marshal(newTicket)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("PUT", "/tickets/4", bytes.NewBuffer(jsonTicket))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/tickets/{id}", handlers.UpdateTicket).Methods("PUT")
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Make a GET request to the same ticket and check that the data was updated correctly
	req, _ = http.NewRequest("GET", "/tickets/4", nil)
	rr = httptest.NewRecorder()
	router.HandleFunc("/tickets/{id}", handlers.GetTicket).Methods("GET")
	router.ServeHTTP(rr, req)

	var updatedTicket models.Ticket
	err = json.NewDecoder(rr.Body).Decode(&updatedTicket)
	if err != nil {
		t.Fatal(err)
	}

	if updatedTicket.User != "Updated User" || updatedTicket.Status != "closed" {
		t.Errorf("ticket was not updated correctly: got %+v", updatedTicket)
	}
}
