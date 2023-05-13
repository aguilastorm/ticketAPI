package models

import (
	"time"
)

type Ticket struct {
	ID           string    `json:"id"`
	User         string    `json:"user"`
	CreationDate time.Time `json:"creation_date"`
	UpdateDate   time.Time `json:"update_date"`
	Status       string    `json:"status"`
}
