package controller

import (
	"PCS-API/service"

	"github.com/gin-gonic/gin"
)

// traveler # Controller
//
// traveler réceptionne toutes les requêtes ayant pour endpoint '/traveler'
// Il les envoie aux fonctions services liés
func Traveler(api *gin.RouterGroup) {
	api.GET("/traveler", service.GetAllTraveler)
}
