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

	fmt.Println("Server running on port 8080")
	http.ListenAndServe(":8080", router)

}
