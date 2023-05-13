package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/aguilastorm/ticketAPI/store"
)

// GetTickets obtiene todos los tickets
func GetTickets(w http.ResponseWriter, r *http.Request) {
	// Establecer el encabezado de la respuesta con el tipo de contenido JSON
	w.Header().Set("Content-Type", "application/json")

	// Obtener todos los tickets desde el almacenamiento y codificarlos en formato JSON
	// para enviarlos en la respuesta
	json.NewEncoder(w).Encode(store.GetTickets())
}
