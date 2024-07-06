package service

import (
	"PCS-API/models"
	"PCS-API/repository"
	"PCS-API/utils"
	"github.com/google/uuid"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// GetAllPropertyImage Récupère la liste de tous les PropertyImage
// @Summary PropertyImage
// @Schemes
// @Description Récupère tous les PropertyImage
// @Tags PropertyImage
// @Produce json
// @Success 200 {array} models.PropertyImage
// @Router /api/Property_image [get]
func GetAllPropertyImage(c *gin.Context) {
	Property_images := repository.GetAllPropertyImage()
	c.JSON(http.StatusOK, gin.H{"Property_image": Property_images})
}

func propertyImageGetArrayPathFromArray(array []models.PropertyImage) []string {
	str := make([]string, len(array))
	for i, v := range array {
		str[i] = v.Path
	}
	return str
}

func propertyImageClean(propertyImage []models.PropertyImage, idProperty uuid.UUID) {
	imagesOrigin := repository.PropertyImageGetAllByIdProperty(idProperty)
	pathOrigin := propertyImageGetArrayPathFromArray(imagesOrigin)
	pathProperty := propertyImageGetArrayPathFromArray(propertyImage)
	for i, v := range pathOrigin {
		if !utils.IsInArrayString(v, pathProperty) {
			repository.PropertyImageDeleteById(imagesOrigin[i].ID)
		}
	}
}
