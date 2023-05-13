package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/aguilastorm/ticketAPI/store"
)

// DeleteTicket elimina un ticket por su ID
func DeleteTicket(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	// Eliminar el ticket del almacenamiento por su ID
	err := store.DeleteTicket(params["id"])
	if err != nil {
		// Si hay un error al eliminar el ticket, enviar una respuesta de error con el código de estado HTTP 404
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	// Establecer el encabezado de la respuesta con el tipo de contenido JSON
	w.Header().Set("Content-Type", "application/json")
}
