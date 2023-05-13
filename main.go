package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"fmt"
	"github.com/aguilastorm/ticketAPI/handlers"
)

func main() {
	// Crear un nuevo enrutador utilizando gorilla/mux
	router := mux.NewRouter().StrictSlash(true)

	// Registrar las rutas y los controladores
	router.HandleFunc("/tickets", handlers.GetTickets).Methods("GET")
	router.HandleFunc("/tickets/{id}", handlers.GetTicket).Methods("GET")
	router.HandleFunc("/tickets", handlers.CreateTicket).Methods("POST")
	router.HandleFunc("/tickets/{id}", handlers.DeleteTicket).Methods("DELETE")
	router.HandleFunc("/tickets/{id}", handlers.UpdateTicket).Methods("PUT")

	// Imprimir un mensaje indicando que el servidor está en ejecución en el puerto 8080
	fmt.Println("Server running on port 8080")

	// Iniciar el servidor HTTP y escuchar en el puerto 8080 utilizando el enrutador definido
	http.ListenAndServe(":8080", router)
}
