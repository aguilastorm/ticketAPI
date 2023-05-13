package handlers

import (
	"encoding/json"
	"net/http"
	"github.com/aguilastorm/ticketAPI/store"
)

func GetTickets(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(store.GetTickets())
}
