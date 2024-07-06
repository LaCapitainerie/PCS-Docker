package models

import (
	"github.com/google/uuid"
)

// Lessor est la structure spécifiant les données des bailleurs utilisé par le front web de l'application
type Lessor struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	FirstName string    `gorm:"type:varchar(64);notnull" json:"firstName"`
	LastName  string    `gorm:"type:varchar(64);notnull" json:"lastName"`
	UserId    uuid.UUID `gorm:"type:uuid" json:"userId"`
}

// TableName Lessor Spécifie à gorm le nom de la table
func (Lessor) TableName() string {
	return "lessor"
}
