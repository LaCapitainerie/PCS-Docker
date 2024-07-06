package service

import (
	"PCS-API/models"
	"PCS-API/repository"
	"PCS-API/utils"
	"github.com/google/uuid"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// GetAllAdmin Récupère la liste de tous les Admin
// @Summary Admin
// @Schemes
// @Description Récupère tous les Admin
// @Tags administration
// @Produce json
// @Success 200 {array} models.Admin
// @Router /api/Admin [get]
func GetAllAdmin(c *gin.Context) {
	Admins := repository.GetAllAdmin()
	c.JSON(http.StatusOK, gin.H{"Admin": Admins})
}

func LoginAdmin(c *gin.Context) {
	var userJson models.UsersDTO
	var err error
	if err = c.BindJSON(&userJson); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := repository.UsersLoginVerify(userJson.Mail)
	if user.Mail == "" || !utils.CheckPassword(user.Password, userJson.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "7"})
		return
	}

	tokenString, err := utils.CreateToken(user.ID.String())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not create token"})
		return
	}

	user.Password = ""
	var userDTO models.UsersDTO
	if user.Type == models.AdminType {
		userDTO = createUserDTOwithUserAndAdmin(user, repository.AdminGetByUserId(user.ID))
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "7"})
		return
	}
	userDTO.Token = tokenString

	c.JSON(http.StatusOK, gin.H{"user": userDTO})
}

func AdminRegister(c *gin.Context) {
	var userDTO models.UsersDTO
	var err error
	if err = c.BindJSON(&userDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if !validityPassword(userDTO.Password) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "1"})
		return
	}

	if !validityEmail(userDTO.Mail) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "2"})
		return
	}

	if repository.UsersVerifyEmail(userDTO.Mail) {
		c.JSON(http.StatusConflict, gin.H{"error": "5"})
		return
	}

	userDTO.ID = uuid.New()
	userDTO.Password, err = utils.HashPassword(userDTO.Password)

	user := convertUserDTOtoUser(userDTO, models.AdminType)
	admin := createAdminWithUserDTO(userDTO)

	if len(admin.Nickname) < 1 &&
		len(admin.Site) < 1 &&
		(len(user.PhoneNumber) < 6 && len(user.PhoneNumber) > 15) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "4"})
		return
	}
	user, err = repository.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	admin, err = repository.AdminCreate(admin)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	tokenString, err := utils.CreateToken(user.ID.String())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not create token"})
		return
	}
	userDTO = createUserDTOwithUserAndAdmin(user, admin)
	userDTO.Token = tokenString
	userDTO.Password = ""
	c.JSON(http.StatusOK, gin.H{"user": userDTO})
}
