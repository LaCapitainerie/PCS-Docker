package models

import (
	"github.com/google/uuid"
)

type ReservationBill struct {
	ReservationId uuid.UUID `gorm:"type:uuid;primaryKey" json:"reservationId"`
	BillId        uuid.UUID `gorm:"type:uuid;primaryKey" json:"billId"`
}

func (ReservationBill) TableName() string {
	return "reservation_bill"
}
