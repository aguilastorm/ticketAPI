package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/aguilastorm/ticketAPI/store"
	"github.com/aguilastorm/ticketAPI/models"
)

// GetTicket obtiene un ticket por su ID
func GetTicket(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	// Obtener todos los tickets desde el almacenamiento
	tickets := store.GetTickets()

	for _, item := range tickets {
		if item.ID == params["id"] {
			// Si se encuentra el ticket, codificarlo en formato JSON y enviarlo en la respuesta
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	// Si no se encuentra el ticket, establecer el encabezado de la respuesta con el tipo de contenido JSON
	w.Header().Set("Content-Type", "application/json")

	// Codificar un ticket vacío en formato JSON y enviarlo en la respuesta
	json.NewEncoder(w).Encode(&models.Ticket{})
}
