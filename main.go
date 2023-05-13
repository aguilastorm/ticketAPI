package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"fmt"
	"github.com/aguilastorm/ticketAPI/handlers"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/tickets", handlers.GetTickets).Methods("GET")
	router.HandleFunc("/tickets/{id}", handlers.GetTicket).Methods("GET")
	router.HandleFunc("/tickets", handlers.CreateTicket).Methods("POST")
	router.HandleFunc("/tickets/{id}", handlers.DeleteTicket).Methods("DELETE")
	router.HandleFunc("/tickets/{id}", handlers.UpdateTicket).Methods("PUT")

	fmt.Println("Server running on port 8080")
	http.ListenAndServe(":8080", router)

}
