package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
	"github.com/gorilla/mux"
	"github.com/aguilastorm/ticketAPI/models"
	"github.com/aguilastorm/ticketAPI/store"
	"github.com/aguilastorm/ticketAPI/handlers"
)

func TestDeleteTicket(t *testing.T) {
	// Configura el ticket de prueba
	ticket := models.Ticket{
		ID:           "3",
		User:         "Test User 3",
		CreationDate: time.Now(),
		UpdateDate:   time.Now(),
		Status:       "open",
	}
	store.AddTicket(ticket)

	// Configura la solicitud de prueba
	req, err := http.NewRequest("DELETE", "/tickets/3", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Configura el grabador de respuesta de prueba
	rr := httptest.NewRecorder()

	// Configura el enrutador y el controlador de prueba
	router := mux.NewRouter()
	router.HandleFunc("/tickets/{id}", handlers.DeleteTicket).Methods("DELETE")
	router.ServeHTTP(rr, req)

	// Verifica el código de estado HTTP de la respuesta
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Verifica que el ticket se eliminó correctamente
	if _, err := store.GetTicket("3"); err == nil {
		t.Errorf("ticket with ID 3 was not deleted")
	}
}
