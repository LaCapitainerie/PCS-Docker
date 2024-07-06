package repository

import (
	"PCS-API/models"
	"PCS-API/utils"
	"github.com/google/uuid"
)

// GetAllPropertyImage
// Renvoie la liste de tous les "Property_image"
func GetAllPropertyImage() []models.PropertyImage {
	var PropertyImages []models.PropertyImage
	if err := utils.DB.Find(&PropertyImages); err.Error != nil {
		panic("Unable to get Property_images " + err.Error.Error())
	}
	return PropertyImages
}

func PropertyImageCreate(image models.PropertyImage) models.PropertyImage {
	utils.DB.Create(&image)
	return image
}

func propertyImageDeleteAllByIdProperty(propertyId uuid.UUID) {
	utils.DB.Where("id = ?", propertyId).Delete(&models.Property{})
}

func PropertyImageGetAllByIdProperty(propertyId uuid.UUID) []models.PropertyImage {
	var images []models.PropertyImage
	utils.DB.Where("property_id = ?", propertyId).Find(&images)
	return images
}

func PropertyImageDeleteById(propertyImageId uuid.UUID) {
	utils.DB.Delete(models.PropertyImage{}, propertyImageId)
}
