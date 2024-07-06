package models

import (
	"github.com/google/uuid"
)

// Admin est la structure spécifiant les données de la Admin utilisé par le front web de l'application
type Admin struct {
	ID       uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Site     string    `gorm:"type:varchar(64)" json:"site"`
	Nickname string    `gorm:"type:varchar(64);notnull" json:"nickname"`
	UserId   uuid.UUID `gorm:"type:uuid;notnull" json:"userId"`
}

// TableName Admin Spécifie à gorm le nom de la table
func (Admin) TableName() string {
	return "administrator"
}
