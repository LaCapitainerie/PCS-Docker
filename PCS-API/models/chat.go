package models

import "github.com/google/uuid"

// Chat est la structure spécifiant les données du chat, la structure la plus généraliste établissant l'échange
// entre deux utilisateurs
type Chat struct {
	ID   uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	View bool      `gorm:"type:boolean;default:false" json:"view"`
}

// TableName Chat Spécifie à gorm le nom de la table
func (Chat) TableName() string {
	return "chat"
}
