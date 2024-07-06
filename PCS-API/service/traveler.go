package service

import (
	"PCS-API/models"
	"PCS-API/repository"
	"PCS-API/utils"
	"net/http"

	"github.com/google/uuid"

	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// GetAllTraveler Récupère la liste de tous les Traveler
// @Summary Traveler
// @Schemes
// @Description Récupère tous les Traveler
// @Tags administration
// @Produce json
// @Success 200 {array} models.Traveler
// @Router /api/Traveler [get]
func GetAllTraveler(c *gin.Context) {
	Travelers := repository.GetAllTraveler()
	c.JSON(http.StatusOK, gin.H{"Traveler": Travelers})
}

// createTraveler crée un nouveau voyageur
// la fonction ne peut être appelé hors du package
func createTraveler(c *gin.Context, userDTO models.UsersDTO) {
	user := convertUserDTOtoUser(userDTO, models.TravelerType)
	traveler := createTravelerWithUserDTO(userDTO)
	var err error

	if len(traveler.FirstName) < 1 &&
		len(traveler.LastName) < 1 &&
		(len(user.PhoneNumber) < 6 && len(user.PhoneNumber) > 15) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "4"})
		return
	}

	user, err = repository.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	traveler, err = repository.CreateTraveler(traveler)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	tokenString, err := utils.CreateToken(user.ID.String())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not create token"})
		return
	}

	userDTO = CreateUserDTOwithUserAndTraveler(user, traveler)
	userDTO.Token = tokenString
	userDTO.Password = ""

	c.JSON(http.StatusOK, gin.H{"user": userDTO})
}

// createTravelerWithUserDTO Crée un voyageur à partir d'un UserDTO
func createTravelerWithUserDTO(dto models.UsersDTO) models.Traveler {
	return models.Traveler{
		ID:        uuid.New(),
		FirstName: dto.FirstName,
		LastName:  dto.LastName,
		UserId:    dto.ID,
	}
}

func createAdminWithUserDTO(dto models.UsersDTO) models.Admin {
	return models.Admin{
		ID:       uuid.New(),
		Site:     dto.Site,
		Nickname: dto.Nickname,
		UserId:   dto.ID,
	}
}

// Update Traveler
// @Summary Traveler
// @Schemes
// @Description Met à jour un Traveler
// @Tags administration
// @Accept json
// @Produce json
// @Param id path string true "ID du Traveler"
// @Param Traveler body models.Traveler true "Traveler"
// @Success 200 {string} string
// @Router /api/Traveler/{id} [put]
func UpdateTraveler(t models.Traveler, id uuid.UUID) {
	utils.DB.Find(&models.Traveler{}, "user_id = ?", id).Update("Traveler", t)
}
