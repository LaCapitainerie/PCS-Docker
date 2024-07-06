package models

import "github.com/google/uuid"

// ChatDTO structure spécifiant un Data Transfer Object un objet de donnée de transfert avec le front
// sert au chat
type ChatDTO struct {
	ID      uuid.UUID  `json:"id"`
	View    bool       `json:"view"`
	Ticket  Ticket     `json:"ticket"`
	UserId  []UsersDTO `json:"userId"`
	Message []Message  `json:"message"`
}
