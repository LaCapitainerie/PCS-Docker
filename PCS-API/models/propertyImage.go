package models

import (
	"github.com/google/uuid"
)

// PropertyImage est la structure spécifiant les données de la Property_image utilisé par le front web de l'application
type PropertyImage struct {
	ID         uuid.UUID `gorm:"type:uuid;primaryKey" json:"ID"`
	Path       string    `gorm:"type:varchar(255);notnull" json:"Path"`
	PropertyId uuid.UUID `gorm:"type:uuid;notnull" json:"property_id"`
}

// TableName PropertyImage Spécifie à gorm le nom de la table
func (PropertyImage) TableName() string {
	return "property_image"
}
