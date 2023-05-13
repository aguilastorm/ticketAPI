package handlers

import (
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/aguilastorm/ticketAPI/models"
	"github.com/aguilastorm/ticketAPI/store"
)


func UpdateTicket(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var updatedTicket models.Ticket
	err := json.NewDecoder(r.Body).Decode(&updatedTicket)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	updatedTicket.ID = params["id"]

	err = store.UpdateTicket(updatedTicket)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedTicket)
}

