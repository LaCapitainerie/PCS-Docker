package controller

import (
	"PCS-API/service"
	"github.com/gin-gonic/gin"
)

// Sidebar # Controller
//
// Sidebar réceptionne toutes les requêtes ayant pour endpoint '/sidebar'
// Il les envoie aux fonctions services liés
func Sidebar(api *gin.RouterGroup) {
	api.GET("/sidebar", service.GetAllSidebar)
}
