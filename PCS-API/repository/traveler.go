package repository

import (
	"PCS-API/models"
	"PCS-API/utils"

	"github.com/google/uuid"
)

// GetAllTraveler
// Renvoie la liste de tous les "Traveler"
func GetAllTraveler() []models.Traveler {
	var Travelers []models.Traveler
	if err := utils.DB.Find(&Travelers); err.Error != nil {
		panic("Unable to get Travelers " + err.Error.Error())
	}
	return Travelers
}

// CreateTraveler reçoit en argument un traveler
// Crée un "traveler" dans la table et renvoie le voyageur mis à jour
func CreateTraveler(traveler models.Traveler) (models.Traveler, error) {
	err := utils.DB.Save(&traveler)
	return traveler, err.Error
}

func TravelerGetByUserId(id uuid.UUID) models.Traveler {
	var traveler models.Traveler
	utils.DB.Where("user_id = ?", id).Find(&traveler)
	return traveler
}

func TravelerGetIdByUserId(id uuid.UUID) uuid.UUID {
	var traveler models.Traveler
	utils.DB.Where("user_id = ?", id).Find(&traveler)
	return traveler.ID
}

func travelerDeleteByUserId(id uuid.UUID) {
	utils.DB.Where("user_id = ?", id).Delete(&models.Traveler{})
}

func UpdateTraveler(traveler models.Traveler) (models.Traveler, error) {
	err := utils.DB.Save(&traveler)
	return traveler, err.Error
}
