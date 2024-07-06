package models

import (
	"github.com/google/uuid"
	"time"
)

type Reservation struct {
	ID         uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	TravelerId uuid.UUID `gorm:"type:uuid;notnull" json:"travelerId"`
	PropertyId uuid.UUID `gorm:"type:uuid;notnull" json:"propertyId"`
	BeginDate  time.Time `gorm:"type:timestamp;notnull" json:"beginDate"`
	EndDate    time.Time `gorm:"type:timestamp;notnull" json:"endDate"`
	Annulation bool      `gorm:"type:boolean;default:false" json:"annulation"`
}

func (Reservation) TableName() string {
	return "reservation"
}
