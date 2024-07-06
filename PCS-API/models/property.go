package models

import (
	"github.com/google/uuid"
)

// Property est la structure spécifiant les données de la property utilisé par le front web de l'application
type Property struct {
	ID                      uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	IdStripe                string    `gorm:"type:varchar(32)" json:"idStripe"`
	Name                    string    `gorm:"type:varchar(64);notnull" json:"name"`
	Type                    string    `gorm:"type:varchar(64);notnull" json:"type"`
	Price                   float64   `gorm:"type:numeric(10,2);notnull" json:"price"`
	Surface                 int       `gorm:"type:integer;notnull" json:"surface"`
	Room                    int       `gorm:"type:integer;notnull" json:"room"`
	Bathroom                int       `gorm:"type:integer;notnull" json:"bathroom"`
	Garage                  int       `gorm:"type:integer" json:"garage"`
	Description             string    `gorm:"type:text" json:"description"`
	Address                 string    `gorm:"type:varchar(64);notnull" json:"address"`
	City                    string    `gorm:"type:varchar(64);notnull" json:"city"`
	ZipCode                 string    `gorm:"type:varchar(6);notnull" json:"zipCode"`
	Country                 string    `gorm:"type:varchar(64);notnull" json:"country"`
	Lon                     float64   `gorm:"type:DOUBLE PRECISION" json:"lon"`
	Lat                     float64   `gorm:"type:DOUBLE PRECISION" json:"lat"`
	AdministratorValidation bool      `gorm:"type:boolean" json:"administrationValidation"`
	LessorId                uuid.UUID `gorm:"type:uuid" json:"lessorId"`
}

// TableName Property Spécifie à gorm le nom de la table
func (Property) TableName() string {
	return "property"
}
