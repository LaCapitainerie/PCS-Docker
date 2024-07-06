package repository

import (
	"PCS-API/models"
	"PCS-API/utils"
)

// GetAllSidebar
// Renvoie la liste de tous les "sidebar"
func GetAllSidebar() []models.Sidebar {
	var sidebars []models.Sidebar
	if err := utils.DB.Find(&sidebars); err.Error != nil {
		panic("Unable to get sidebars " + err.Error.Error())
	}
	return sidebars
}
