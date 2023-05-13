package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/aguilastorm/ticketAPI/models"
	"github.com/aguilastorm/ticketAPI/store"
)

// CreateTicket crea un nuevo ticket
func CreateTicket(w http.ResponseWriter, r *http.Request) {
	var ticket models.Ticket

	// Decodificar el cuerpo de la solicitud en la estructura del ticket
	err := json.NewDecoder(r.Body).Decode(&ticket)
	if err != nil {
		// Si hay un error en la decodificación, enviar una respuesta de error con el código de estado HTTP 400
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Agregar el ticket al almacenamiento
	store.AddTicket(ticket)

	// Establecer el encabezado de la respuesta con el tipo de contenido JSON
	w.Header().Set("Content-Type", "application/json")

	// Codificar el ticket en formato JSON y enviarlo en la respuesta
	json.NewEncoder(w).Encode(ticket)
}
