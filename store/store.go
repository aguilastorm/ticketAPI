package store

import (
	"github.com/aguilastorm/ticketAPI/models"
	"fmt"
)

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

// DeleteTicket elimina un ticket de la lista de tickets.
func DeleteTicket(id string) error {
	for index, ticket := range Tickets {
		if ticket.ID == id {
			// Elimina el ticket de la lista
			Tickets = append(Tickets[:index], Tickets[index+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Ticket with ID %s not found", id)
}

// GetTicket retorna un ticket específico por su ID.
func GetTicket(id string) (*models.Ticket, error) {
	for _, ticket := range Tickets {
		if ticket.ID == id {
			return &ticket, nil
		}
	}
	return nil, fmt.Errorf("Ticket with ID %s not found", id)
}

// UpdateTicket actualiza un ticket existente.
func UpdateTicket(updatedTicket models.Ticket) error {
	for index, ticket := range Tickets {
		if ticket.ID == updatedTicket.ID {
			// Actualiza el ticket en la lista.
			Tickets[index] = updatedTicket
			return nil
		}
	}
	return fmt.Errorf("Ticket with ID %s not found", updatedTicket.ID)
}
