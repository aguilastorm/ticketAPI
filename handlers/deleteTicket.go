package handlers

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/aguilastorm/ticketAPI/store"
)

func DeleteTicket(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	err := store.DeleteTicket(params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
}
