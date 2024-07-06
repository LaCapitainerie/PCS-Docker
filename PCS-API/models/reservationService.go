package models

import (
	"github.com/google/uuid"
	"time"
)

type ReservationService struct {
	ReservationId uuid.UUID `gorm:"type:uuid;primaryKey" json:"reservationId"`
	ServiceId     uuid.UUID `gorm:"type:uuid;primaryKey" json:"serviceId"`
	Date          time.Time `gorm:"type:timestamp;notnull" json:"date"`
}

func (ReservationService) TableName() string {
	return "reservation_service"
}
