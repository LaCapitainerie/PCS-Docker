package models

import "github.com/google/uuid"

type PropertyDTO struct {
	ID                      uuid.UUID `json:"id"`
	IdStripe                string    `json:"idStripe"`
	Name                    string    `json:"name"`
	Type                    string    `json:"type"`
	Price                   float64   `json:"price"`
	Surface                 int       `json:"surface"`
	Room                    int       `json:"room"`
	Bathroom                int       `json:"bathroom"`
	Garage                  int       `json:"garage"`
	Description             string    `json:"description"`
	Address                 string    `json:"address"`
	City                    string    `json:"city"`
	ZipCode                 string    `json:"zipCode"`
	Lon                     float64   `json:"lon"`
	Lat                     float64   `json:"lat"`
	Images                  []string  `json:"images"`
	Country                 string    `json:"country"`
	AdministratorValidation bool      `json:"administrationValidation"`
	UserId                  uuid.UUID `json:"userId"`
	LessorId                uuid.UUID `json:"lessorId"`
}
