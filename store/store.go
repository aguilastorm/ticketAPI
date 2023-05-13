package store

import "github.com/aguilastorm/ticketAPI/models"

var Tickets []models.Ticket

// GetTickets devuelve una copia de los tickets.
func GetTickets() []models.Ticket {
	return Tickets
}

// AddTicket agrega un ticket a la lista de tickets.
func AddTicket(ticket models.Ticket) {
	Tickets = append(Tickets, ticket)
}

// Limpia los tickets
func ClearTickets() {
	Tickets = []models.Ticket{}
}
