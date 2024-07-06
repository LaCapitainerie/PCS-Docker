package repository

import (
	"PCS-API/models"
	"PCS-API/utils"
	"github.com/google/uuid"
)

// GetAllAdmin
// Renvoie la liste de tous les "Admin"
func GetAllAdmin() []models.Admin {
	var Admins []models.Admin
	if err := utils.DB.Find(&Admins); err.Error != nil {
		panic("Unable to get Admins " + err.Error.Error())
	}
	return Admins
}

func AdminGetByUserId(id uuid.UUID) models.Admin {
	var admin models.Admin
	utils.DB.Where("user_id = ?", id).Find(&admin)
	return admin
}

func AdminCreate(admin models.Admin) (models.Admin, error) {
	err := utils.DB.Save(&admin).Error
	return admin, err
}
