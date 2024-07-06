package service

import (
	"PCS-API/models"
	"PCS-API/repository"
	"PCS-API/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// createLessor crée un nouveau bailleur
// la fonction ne peut être appelé hors du package
func createLessor(c *gin.Context, userDTO models.UsersDTO) {
	user := convertUserDTOtoUser(userDTO, models.LessorType)
	lessor := createLessorWithUserDTO(userDTO)
	var err error

	if len(lessor.FirstName) < 1 &&
		len(lessor.LastName) < 1 &&
		(len(user.PhoneNumber) < 6 && len(user.PhoneNumber) > 15) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "4"})
		return
	}

	user, err = repository.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	lessor, err = repository.CreateLessor(lessor)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	tokenString, err := utils.CreateToken(user.ID.String())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not create token"})
		return
	}
	userDTO = CreateUserDTOwithUserAndLessor(user, lessor)
	userDTO.Token = tokenString
	userDTO.Password = ""

	c.JSON(http.StatusOK, gin.H{"user": userDTO})
}

// createLessorWithUserDTO Crée un bailleur à partir d'un UserDTO
func createLessorWithUserDTO(dto models.UsersDTO) models.Lessor {
	return models.Lessor{
		ID:        uuid.New(),
		FirstName: dto.FirstName,
		LastName:  dto.LastName,
		UserId:    dto.ID,
	}
}
