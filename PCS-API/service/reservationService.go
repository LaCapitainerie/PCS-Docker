package service

import (
	"PCS-API/models"
	"PCS-API/repository"
	"fmt"
	"github.com/google/uuid"
)

func reservationGetAllService(dto *models.ReservationDTO) ([]models.Service, error) {
	services := make([]models.Service, len(dto.Service))
	var err error
	for i, service := range dto.Service {
		services[i], err = repository.ServiceGetWithServiceId(service.ID)
		if !((service.Date.Equal(dto.BeginDate) || service.Date.After(dto.BeginDate)) &&
			(service.Date.Equal(dto.EndDate) || service.Date.Before(dto.EndDate))) {
			return services, fmt.Errorf("1")
		}

		if err != nil {
			return services, err
		}
	}
	return services, nil
}

func reservationServiceListCreate(dto *models.ReservationDTO, services []models.Service, idReservation *uuid.UUID) ([]models.ServiceDTO, error) {
	var reservationService models.ReservationService
	var err error
	serviceDTO := make([]models.ServiceDTO, len(dto.Service))
	for i, service := range dto.Service {
		serviceDTO[i] = serviceConvertToServiceDTO(
			services[i],
			repository.ProviderGetUserIdWithProviderId(services[i].ProviderId),
			service.Date)

		reservationService.ReservationId = *idReservation
		reservationService.ServiceId = service.ID
		reservationService.Date = service.Date
		reservationService, err = repository.ReservationServiceListCreate(reservationService)
		if err != nil {
			return serviceDTO, err
		}
	}
	return serviceDTO, nil
}
