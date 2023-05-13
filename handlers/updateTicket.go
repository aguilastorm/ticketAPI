package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/aguilastorm/ticketAPI/models"
	"github.com/aguilastorm/ticketAPI/store"
)

// UpdateTicket actualiza un ticket por su ID
func UpdateTicket(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	// Decodificar el cuerpo de la solicitud en la estructura del ticket actualizado
	var updatedTicket models.Ticket
	err := json.NewDecoder(r.Body).Decode(&updatedTicket)
	if err != nil {
		// Si hay un error en la decodificación, enviar una respuesta de error con el código de estado HTTP 400
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Asignar el ID del ticket desde los parámetros de la ruta
	updatedTicket.ID = params["id"]

	// Actualizar el ticket en el almacenamiento
	err = store.UpdateTicket(updatedTicket)
	if err != nil {
		// Si hay un error al actualizar el ticket, enviar una respuesta de error con el código de estado HTTP 404
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	// Establecer el encabezado de la respuesta con el tipo de contenido JSON
	w.Header().Set("Content-Type", "application/json")

	// Codificar el ticket actualizado en formato JSON y enviarlo en la respuesta
	json.NewEncoder(w).Encode(updatedTicket)
}
