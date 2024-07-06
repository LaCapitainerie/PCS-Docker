package models

import (
	"time"

	"github.com/google/uuid"
)

// Service est la structure spécifiant les données des prestations proposé par les prestataires
type Service struct {
	ID             uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	IdStripe       string    `gorm:"type:varchar(32)" json:"idStripe"`
	Name           string    `gorm:"type:varchar(64)" json:"name"`
	Price          float64   `gorm:"type:numeric(10,2);notnull" json:"price"`
	TargetCustomer string    `gorm:"type:varchar(8);notnull" json:"targetCustomer"`
	Address        string    `gorm:"type:varchar(64);notnull" json:"address"`
	City           string    `gorm:"type:varchar(64);notnull" json:"city"`
	ZipCode        string    `gorm:"type:varchar(6);notnull" json:"zipCode"`
	Country        string    `gorm:"type:varchar(64);notnull" json:"country"`
	Lat            float64   `gorm:"type:DOUBLE PRECISION" json:"lat"`
	Lon            float64   `gorm:"type:DOUBLE PRECISION" json:"lon"`
	RangeAction    int       `gorm:"type:integer" json:"rangeAction"`
	Description    string    `gorm:"type:text;notnull" json:"description"`
	ProviderId     uuid.UUID `gorm:"type:uuid;notnull" json:"providerId"`
	Type           string    `gorm:"type:varchar(64);notnull" json:"type"`
}

type ServiceDTO struct {
	Service
	UserId uuid.UUID `json:"userId"`
	Date   time.Time `json:"date"`
}

// TableName Service Spécifie à gorm le nom de la table
func (Service) TableName() string {
	return "service"
}
