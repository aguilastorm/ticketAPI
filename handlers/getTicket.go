package handlers

import (
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/aguilastorm/ticketAPI/store"
	"github.com/aguilastorm/ticketAPI/models"
)

func GetTicket(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range store.GetTickets() {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&models.Ticket{})
}
