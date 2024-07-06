package repository

import (
	"PCS-API/models"
	"PCS-API/utils"
	"github.com/google/uuid"
)

// CreateProvider reçoit en argument un provider
// Crée un "provider" dans la table et renvoie le prestataire mis à jour
func CreateProvider(traveler models.Provider) (models.Provider, error) {
	err := utils.DB.Save(&traveler)
	return traveler, err.Error
}

func ProviderGetByUserId(id uuid.UUID) models.Provider {
	var provider models.Provider
	utils.DB.Where("user_id = ?", id).Find(&provider)
	return provider
}

func ProviderGetUserIdWithProviderId(id uuid.UUID) uuid.UUID {
	var provider models.Provider
	utils.DB.First(&provider, id)
	return provider.UserId
}

func providerDeleteByUserId(id uuid.UUID) {
	utils.DB.Where("user_id = ?", id).Delete(&models.Provider{})
}
