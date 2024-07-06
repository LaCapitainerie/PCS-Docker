package repository

import (
	"PCS-API/models"
	"PCS-API/utils"

	"github.com/google/uuid"
)

func ServiceCreateNewService(service models.Service) (models.Service, error) {
	err := utils.DB.Create(&service).Error
	return service, err
}

func ServiceUpdate(service models.Service) models.Service {
	utils.DB.Save(&service)
	return service
}

func ServiceGetWithServiceId(id uuid.UUID) (models.Service, error) {
	var service models.Service
	err := utils.DB.First(&service, id).Error
	return service, err
}

func ServiceGetAll() ([]models.Service, error) {
	var services []models.Service
	if err := utils.DB.Find(&services).Error; err != nil {
		return nil, err
	}
	return services, nil
}

func ServiceDeleteById(id uuid.UUID) error {
	err := utils.DB.Where("id = ?", id).Delete(&models.Service{}).Error
	return err
}
