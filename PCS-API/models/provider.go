package models

import "github.com/google/uuid"

// Provider est la structure spécifiant les données des prestataires utilisé par le front web de l'application
type Provider struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	FirstName string    `gorm:"type:varchar(64);notnull" json:"firstName"`
	LastName  string    `gorm:"type:varchar(64);notnull" json:"lastName"`
	Nickname  string    `gorm:"type:varchar(64);notnull" json:"nickname"`
	UserId    uuid.UUID `gorm:"type:uuid" json:"userId"`
}

// TableName Provider Spécifie à gorm le nom de la table
func (Provider) TableName() string {
	return "provider"
}
