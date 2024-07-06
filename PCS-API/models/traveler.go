package models

import (
	"github.com/google/uuid"
)

// Traveler est la structure spécifiant les données de la Traveler utilisé par le front web de l'application
type Traveler struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey" json:"Id"`
	FirstName string    `gorm:"type:varchar(255);not null" json:"firstName"`
	LastName  string    `gorm:"type:varchar(255);not null" json:"lastName"`
	UserId    uuid.UUID `gorm:"type:uuid;not null" json:"userId"`
}

// TableName Traveler Spécifie à gorm le nom de la base de donnée
func (Traveler) TableName() string {
	return "traveler"
}
