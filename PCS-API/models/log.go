package models

import (
	"time"

	"github.com/google/uuid"
)

// Log est la structure spécifiant les données des logs utilisé par l'application
type Log struct {
	ID       uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	UserID   uuid.UUID `gorm:"type:uuid" json:"userId"`
	Action   string    `gorm:"type:varchar(64);notnull" json:"action"`
	Endpoint string    `gorm:"type:varchar(64);notnull" json:"endpoint"`
	Time     time.Time `gorm:"type:timestamp;notnull" json:"time"`
}

// TableName Log Spécifie à gorm le nom de la table
func (Log) TableName() string {
	return "log"
}
