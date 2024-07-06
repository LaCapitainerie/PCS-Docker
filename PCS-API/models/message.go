package models

import (
	"time"

	"github.com/google/uuid"
)

// Message est la structure spécifiant les données des messages
type Message struct {
	ID      uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Content string    `gorm:"type:text;notnull" json:"content"`
	Date    time.Time `gorm:"type:timestamp;notnull;default:current_timestamp" json:"date"`
	Type    string    `gorm:"type:varchar(5)" json:"type"`
	UserId  uuid.UUID `gorm:"type:uuid" json:"userId"`
	ChatId  uuid.UUID `gorm:"type:uuid" json:"chatId"`
}

// TableName Message Spécifie à gorm le nom de la table
func (Message) TableName() string {
	return "message"
}
