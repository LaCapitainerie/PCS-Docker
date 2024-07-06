package service

import (
	"PCS-API/repository"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @BasePath /api/v1

// GetAllSidebar Récupère la liste de tous les Sidebar
// @Summary Sidebar
// @Schemes
// @Description Récupère tous les Sidebar
// @Tags administration
// @Produce json
// @Success 200 {array} models.Sidebar
// @Router /api/sidebar [get]
func GetAllSidebar(c *gin.Context) {
	sidebars := repository.GetAllSidebar()
	c.JSON(http.StatusOK, gin.H{"Sidebar": sidebars})
}
