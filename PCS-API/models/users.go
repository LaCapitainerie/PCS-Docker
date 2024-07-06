// Package models
// package spécifiant chaque structure utilisé dans le programme
package models

import (
	"github.com/google/uuid"
	"time"
)

// Users est la structure spécifiant les données utilisateur telles qu'ils sont contenu dans la base
// Il s'agit de la base du code et chaque utilisateur quel qu'il soit (bailleur, presta, admin, voyageur) est affilié à une table Users
type Users struct {
	ID                 uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Mail               string    `json:"mail"`
	Password           string    `json:"password"`
	Avatar             string    `json:"avatar"`
	Type               string    `json:"type"`
	Description        string    `json:"description"`
	PhoneNumber        string    `gorm:"type:varchar(15)" json:"phoneNumber"`
	RegisterDate       time.Time `gorm:"type:timestamp with time zone;not null;default:CURRENT_TIMESTAMP" json:"register_date"`
	LastConnectionDate time.Time `gorm:"type:timestamp with time zone;default:CURRENT_TIMESTAMP" json:"last_connection_date"`
}

// TableName Users Spécifie à gorm le nom de la base de donnée
func (Users) TableName() string {
	return "users"
}
