package repository

import (
	"PCS-API/models"
	"PCS-API/utils"
	"github.com/google/uuid"
)

func BillCreate(bill models.Bill, idReservation uuid.UUID) (models.Bill, error) {
	err := utils.DB.Create(&bill).Error
	if err != nil {
		return bill, err
	}
	reservationBill := models.ReservationBill{
		ReservationId: idReservation,
		BillId:        bill.ID,
	}
	err = utils.DB.Create(&reservationBill).Error
	return bill, err
}

func BillGetByReservationId(id uuid.UUID) (models.Bill, error) {
	var bill models.Bill
	err := utils.DB.
		Where("id = (?)", utils.DB.Model(&models.ReservationBill{}).
			Select("bill_id").
			Where("reservation_id = ?", id)).
		First(&bill).Error
	return bill, err
}
