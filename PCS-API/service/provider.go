package service

import (
	"PCS-API/models"
	"PCS-API/repository"
	"PCS-API/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// createProvider crée un nouveau prestataire
// la fonction ne peut être appelé hors du package
func createProvider(c *gin.Context, userDTO models.UsersDTO) {
	user := convertUserDTOtoUser(userDTO, models.ProviderType)
	provider := createProviderWithUserDTO(userDTO)
	var err error

	if len(provider.FirstName) < 1 &&
		len(provider.LastName) < 1 &&
		(len(user.PhoneNumber) < 6 && len(user.PhoneNumber) > 15) &&
		len(provider.Nickname) < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "4"})
		return
	}

	user, err = repository.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	provider, err = repository.CreateProvider(provider)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	tokenString, err := utils.CreateToken(user.ID.String())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not create token"})
		return
	}
	userDTO = CreateUserDTOwithUserAndProvider(user, provider)
	userDTO.Token = tokenString
	userDTO.Password = ""

	c.JSON(http.StatusOK, gin.H{"user": userDTO})
}

// createProviderWithUserDTO Crée un prestataire à partir d'un UserDTO
func createProviderWithUserDTO(dto models.UsersDTO) models.Provider {
	return models.Provider{
		ID:        uuid.New(),
		FirstName: dto.FirstName,
		LastName:  dto.LastName,
		Nickname:  dto.Nickname,
		UserId:    dto.ID,
	}
}
