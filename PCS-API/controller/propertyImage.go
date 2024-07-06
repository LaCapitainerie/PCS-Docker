package controller

import (
	"PCS-API/service"

	"github.com/gin-gonic/gin"
)

// Property_image # Controller
//
// Property_image réceptionne toutes les requêtes ayant pour endpoint '/Property_image'
// Il les envoie aux fonctions services liés
func Property_image(api *gin.RouterGroup) {
	api.GET("/property_image", service.GetAllPropertyImage)
}
