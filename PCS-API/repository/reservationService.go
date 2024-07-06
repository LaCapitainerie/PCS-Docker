package repository

import (
	"PCS-API/models"
	"PCS-API/utils"
	"github.com/google/uuid"
)

func ReservationServiceListCreate(service models.ReservationService) (models.ReservationService, error) {
	err := utils.DB.Create(&service).Error
	return service, err
}

func ReservationServiceGetAllByAReservationId(reservationId uuid.UUID) ([]models.ServiceDTO, error) {
	var reservationsService []models.ReservationService
	err := utils.DB.Find(&reservationsService, "reservation_id = ?", reservationId).Error
	if err != nil {
		return []models.ServiceDTO{}, err
	}
	servicesDTO := make([]models.ServiceDTO, len(reservationsService))
	for i, reservation := range reservationsService {
		service, err := ServiceGetWithServiceId(reservation.ServiceId)
		if err != nil {
			return []models.ServiceDTO{}, err
		}
		servicesDTO[i] = models.ServiceDTO{
			Service: service,
			Date:    reservation.Date,
		}
	}
	return servicesDTO, nil
}
