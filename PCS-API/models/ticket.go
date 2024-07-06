package models

import (
	"github.com/google/uuid"
)

const (
	TICKET_STATE_CLOSE    string = "close"
	TICKET_STATE_OPEN     string = "open"
	TICKET_STATE_PROGRESS string = "progress"
)

const (
	TICKET_TYPE_PAIEMENT  string = "paiement"
	TICKET_TYPE_TECHNIQUE string = "technique"
)

// Ticket est la structure spécifiant les données des tickets, un type de chat
type Ticket struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Type        string    `gorm:"type:varchar(64);notnull" json:"type"`
	State       string    `gorm:"type:varchar(16);notnull" json:"state"`
	Description string    `gorm:"type:text;notnull" json:"description"`
	ChatId      uuid.UUID `gorm:"type:uuid" json:"chatId"`
}

// TableName Ticket Spécifie à gorm le nom de la table
func (Ticket) TableName() string {
	return "ticket"
}
