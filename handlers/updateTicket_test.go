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
)

func TestUpdateTicket(t *testing.T) {
	tickets = append(store.GeTtickets(), models.Ticket{ID: "4", User: "Test User 4", CreationDate: time.Now(), UpdateDate: time.Now(), Status: "open"})

	newTicket := &models.Ticket{
		ID:           "4",
		User:         "Updated User",
		CreationDate: time.Now(),
		UpdateDate:   time.Now(),
		Status:       "closed",
	}

	jsonTicket, _ := json.Marshal(newTicket)
	req, err := http.NewRequest("PUT", "/tickets/4", bytes.NewBuffer(jsonTicket))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/tickets/{id}", UpdateTicket)
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
