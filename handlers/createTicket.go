package handlers

import (

	"encoding/json"
	"net/http"
	"github.com/aguilastorm/ticketAPI/models"
	"github.com/aguilastorm/ticketAPI/store"
)

func CreateTicket(w http.ResponseWriter, r *http.Request) {
	var ticket models.Ticket
	err := json.NewDecoder(r.Body).Decode(&ticket)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	store.AddTicket(ticket)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ticket)
}
