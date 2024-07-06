package repository

import (
	"PCS-API/models"
	"PCS-API/utils"
	"errors"
	"github.com/google/uuid"
)

// GetAllProperty
// Renvoie la liste de tous les "Property"
func GetAllProperty() []models.Property {
	var Propertys []models.Property
	if err := utils.DB.Find(&Propertys); err.Error != nil {
		panic("Unable to get Propertys " + err.Error.Error())
	}
	return Propertys
}

// TODO: Proposer un modèle unifier entre créer et update
func PropertyCreate(property models.Property) (models.Property, error) {
	err := utils.DB.Create(&property)
	return property, err.Error
}

func PropertyUpdate(property models.Property) (models.Property, error) {
	err := utils.DB.Save(&property)
	return property, err.Error
}

func PropertyDeleteWithIdUserAndPropertyId(propertyId uuid.UUID, lessorId uuid.UUID) error {
	if !propertyVerifOwnerById(propertyId, lessorId) {
		return errors.New("17")
	}

	utils.DB.Where("property_id = ?", propertyId).Delete(&models.PropertyImage{})
	propertyImageDeleteAllByIdProperty(propertyId)
	return nil
}

func propertyVerifOwnerById(propertyId uuid.UUID, lessorId uuid.UUID) bool {
	var count int64
	utils.DB.Model(models.Property{}).Where("lessor_id = ? AND id = ?", lessorId, propertyId).Count(&count)
	return count > 0
}

func PropertyGetById(propertyId uuid.UUID) (models.Property, error) {
	var property models.Property
	err := utils.DB.First(&property, propertyId).Error
	if err != nil {
		return property, err
	}
	return property, nil
}
