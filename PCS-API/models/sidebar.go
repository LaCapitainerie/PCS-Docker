package models

import (
	"github.com/google/uuid"
)

// Sidebar est la structure spécifiant les données de la sidebar utilisé par le front web de l'application
type Sidebar struct {
	ID         uuid.UUID `gorm:"type:uuid;primaryKey" json:"Id_tab"`
	Permission int       `json:"Permission"`
	Icon       string    `json:"Icon"`
	Hover      string    `json:"Hover"`
	Href       string    `json:"Href"`
}

// TableName Sidebar Spécifie à gorm le nom de la base de donnée
func (Sidebar) TableName() string {
	return "sidebar"
}
