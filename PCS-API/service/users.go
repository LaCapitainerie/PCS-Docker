// Package service spécifie le code "métier" de l'API
package service

import (
	"PCS-API/models"
	"PCS-API/repository"
	"PCS-API/utils"
	"net/http"
	"regexp"
	"strconv"
	"unicode"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// CreateUser Crée un utilisateur
// @Summary User
// @Schemes
// @Description Crée un nouvel utilisateur
// @Tags Création
// @Produce json
// @Param user body models.UsersDTO true "User to create"
// @Success 200 {object} models.UsersDTO "Retourne l'utilisateur crée"
// @Failure 400 {string} error "Requête incorrecte - données invalides"
// @Failure 409 {string} error "Conflit - L'email ou le mot de passe existe déjà"
// @Router /api/user/register [post]
func CreateUser(c *gin.Context) {
	var user models.UsersDTO
	var err error
	if err = c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if !validityPassword(user.Password) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "1"})
		return
	}

	if !validityEmail(user.Mail) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "2"})
		return
	}

	if repository.UsersVerifyEmail(user.Mail) {
		c.JSON(http.StatusConflict, gin.H{"error": "5"})
		return
	}

	if user.TypeUser != models.AdminType && repository.UsersVerifyPhone(user.PhoneNumber) {
		c.JSON(http.StatusConflict, gin.H{"error": "6"})
		return
	}

	user.ID = uuid.New()
	user.Password, err = utils.HashPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if user.TypeUser == models.TravelerType {
		createTraveler(c, user)
	} else if user.TypeUser == models.ProviderType {
		createProvider(c, user)
	} else if user.TypeUser == models.LessorType {
		createLessor(c, user)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "3"})
	}
}

// LoginUser Se connecte à un utilisateur
// @Summary User
// @Schemes
// @Description Se connecte à un utilisateur
// @Tags Connexion
// @Produce json
// @Param mail body string true "Mail de l'utilisateur"
// @Param password body string true "Mot de passe de l'utilisateur"
// @Success 200 {string} token "Retourne un token de connexion"
// @Failure 400 {string} error "Requête incorrecte - données invalides"
// @Failure 401 {string} error "L'email ou le mot de passe est invalide"
// @Router /api/user/login [post]
func LoginUser(c *gin.Context) {
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
	if user.Type == models.ProviderType {
		userDTO = CreateUserDTOwithUserAndProvider(user, repository.ProviderGetByUserId(user.ID))
	} else if user.Type == models.LessorType {
		userDTO = CreateUserDTOwithUserAndLessor(user, repository.LessorGetByUserId(user.ID))
	} else if user.Type == models.TravelerType {
		userDTO = CreateUserDTOwithUserAndTraveler(user, repository.TravelerGetByUserId(user.ID))
	} else if user.Type == models.AdminType {
		userDTO = createUserDTOwithUserAndAdmin(user, repository.AdminGetByUserId(user.ID))
	}
	userDTO.Token = tokenString

	c.JSON(http.StatusOK, gin.H{"user": userDTO})
}

// convertUserDTOtoUser Crée un utilisateur à partir d'un UserDTO
func convertUserDTOtoUser(userDTO models.UsersDTO, typeUser string) models.Users {
	return models.Users{
		ID:                 userDTO.ID,
		Mail:               userDTO.Mail,
		Password:           userDTO.Password,
		Type:               typeUser,
		RegisterDate:       userDTO.RegisterDate,
		LastConnectionDate: userDTO.LastConnectionDate,
		PhoneNumber:        userDTO.PhoneNumber,
		Avatar:             userDTO.Avatar,
		Description:        userDTO.Description,
	}
}

// CreateUserDTOwithUserAndLessor Crée un userDTO à partir d'un utilisateur et d'un bailleur
func CreateUserDTOwithUserAndLessor(users models.Users, lessor models.Lessor) models.UsersDTO {
	return models.UsersDTO{
		ID:                 users.ID,
		TypeUser:           models.LessorType,
		Mail:               users.Mail,
		Password:           users.Password,
		RegisterDate:       users.RegisterDate,
		LastConnectionDate: users.LastConnectionDate,
		FirstName:          lessor.FirstName,
		LastName:           lessor.LastName,
		PhoneNumber:        users.PhoneNumber,
		Avatar:             users.Avatar,
		Description:        users.Description,
	}
}

// CreateUserDTOwithUserAndTraveler Crée un userDTO à partir d'un utilisateur et d'un voyageur
func CreateUserDTOwithUserAndTraveler(users models.Users, traveler models.Traveler) models.UsersDTO {
	return models.UsersDTO{
		ID:                 users.ID,
		TypeUser:           models.TravelerType,
		Mail:               users.Mail,
		Password:           users.Password,
		RegisterDate:       users.RegisterDate,
		LastConnectionDate: users.LastConnectionDate,
		FirstName:          traveler.FirstName,
		LastName:           traveler.LastName,
		PhoneNumber:        users.PhoneNumber,
		Avatar:             users.Avatar,
		Description:        users.Description,
	}
}

// CreateUserDTOwithUserAndTraveler Crée un userDTO à partir d'un utilisateur et d'un prestataire
func CreateUserDTOwithUserAndProvider(users models.Users, provider models.Provider) models.UsersDTO {
	return models.UsersDTO{
		ID:                 users.ID,
		TypeUser:           models.ProviderType,
		Mail:               users.Mail,
		Password:           users.Password,
		RegisterDate:       users.RegisterDate,
		LastConnectionDate: users.LastConnectionDate,
		Nickname:           provider.Nickname,
		FirstName:          provider.FirstName,
		LastName:           provider.LastName,
		PhoneNumber:        users.PhoneNumber,
		Avatar:             users.Avatar,
		Description:        users.Description,
	}
}

func createUserDTOwithUserAndAdmin(users models.Users, admin models.Admin) models.UsersDTO {
	return models.UsersDTO{
		ID:                 users.ID,
		TypeUser:           models.AdminType,
		Mail:               users.Mail,
		Password:           users.Password,
		RegisterDate:       users.RegisterDate,
		LastConnectionDate: users.LastConnectionDate,
		Nickname:           admin.Nickname,
		Site:               admin.Site,
		PhoneNumber:        users.PhoneNumber,
		Avatar:             users.Avatar,
		Description:        users.Description,
	}
}

// validityPassword Vérifie la validité d'un mot de passe
func validityPassword(password string) bool {
	var check [4]bool
	if len(password) < 8 || len(password) > 128 {
		return false
	}
	for _, char := range password {
		if unicode.IsUpper(char) {
			check[0] = true
		} else if unicode.IsLower(char) {
			check[1] = true
		} else if !unicode.IsLetter(char) && !unicode.IsDigit(char) {
			check[2] = true
		} else if unicode.IsDigit(char) {
			check[3] = true
		}

		if check[0] &&
			check[1] &&
			check[2] &&
			check[3] {
			return true
		}
	}
	return false
}

func validityEmail(email string) bool {
	regex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return regex.MatchString(email)
}

func UserGetById(c *gin.Context) {
	id, _ := uuid.Parse(c.Param("id"))
	var userDTO models.UsersDTO
	user, err := repository.UsersGetById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"user": "10"})
		return
	}
	switch user.Type {
	case models.TravelerType:
		provider := repository.ProviderGetByUserId(user.ID)
		userDTO = CreateUserDTOwithUserAndProvider(user, provider)
	case models.ProviderType:
		traveler := repository.TravelerGetByUserId(user.ID)
		userDTO = CreateUserDTOwithUserAndTraveler(user, traveler)
	case models.LessorType:
		lessor := repository.LessorGetByUserId(user.ID)
		userDTO = CreateUserDTOwithUserAndLessor(user, lessor)
	}
	userDTO.Password = ""

	c.JSON(http.StatusOK, gin.H{"user": userDTO})
}

func UserDeleteById(c *gin.Context) {
	idUserDelete, _ := uuid.Parse(c.Param("id"))
	user, err := repository.UsersGetById(idUserDelete)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{})
		return
	}

	idBrut, exist := c.Get("idUser")
	if !exist {
		c.JSON(http.StatusBadRequest, gin.H{"error": "8"})
		return
	}

	idUser, _ := uuid.Parse(idBrut.(string))

	userAuthor, error := repository.UsersGetById(idUser)
	if error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": error.Error()})
		return
	}

	if userAuthor.Type != models.AdminType && idUser != idUserDelete {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "18"})
		return
	}

	err = repository.UsersDelete(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}
	c.JSON(http.StatusOK, gin.H{})
}

func UserUpdateById(c *gin.Context) {
	idCurrentUser, _ := uuid.Parse(c.Param("id"))
	idBrut, exist := c.Get("idUser")
	if !exist {
		c.JSON(http.StatusBadRequest, gin.H{"error": "8"})
		return
	}
	idUser, _ := uuid.Parse(idBrut.(string))
	if idUser != idCurrentUser {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "18"})
		return
	}

	var userDTO models.UsersDTO
	var err error
	if err = c.BindJSON(&userDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userSelected, _ := repository.UsersGetById(idUser)

	if userDTO.Password != "" {
		if !validityPassword(userDTO.Password) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "1"})
			return
		}
		userSelected.Password, err = utils.HashPassword(userDTO.Password)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}

	//TODO: Gérer le problème de redondance
	userDTO.ID = userSelected.ID
	userDTO.Mail = userSelected.Mail
	userDTO.Password = userSelected.Password
	userDTO.RegisterDate = userSelected.RegisterDate
	userDTO.LastConnectionDate = userSelected.LastConnectionDate
	userDTO.TypeUser = userSelected.Type

	if userDTO.TypeUser == models.TravelerType {

		traveler := repository.TravelerGetByUserId(userSelected.ID)
		traveler.FirstName = userDTO.FirstName
		traveler.LastName = userDTO.LastName

		UpdateTraveler(traveler, userDTO.ID)
	} else if userDTO.TypeUser == models.ProviderType {
		createProvider(c, userDTO)
	} else if userDTO.TypeUser == models.LessorType {
		createLessor(c, userDTO)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "3"})
	}
}

func UserGetAll(c *gin.Context) {

	limit, err := strconv.Atoi(c.Query("limit"))
	if err != nil {
		limit = 100 // default limit
	}

	offset, err := strconv.Atoi(c.Query("offset"))
	if err != nil {
		offset = 0 // default offset
	}

	users, _ := repository.UsersGetAll(limit, offset)
	var usersDTO []models.UsersDTO
	for _, user := range users {
		switch user.Type {
		case models.TravelerType:
			traveler := repository.TravelerGetByUserId(user.ID)
			usersDTO = append(usersDTO, CreateUserDTOwithUserAndTraveler(user, traveler))
		case models.ProviderType:
			provider := repository.ProviderGetByUserId(user.ID)
			usersDTO = append(usersDTO, CreateUserDTOwithUserAndProvider(user, provider))
		case models.LessorType:
			lessor := repository.LessorGetByUserId(user.ID)
			usersDTO = append(usersDTO, CreateUserDTOwithUserAndLessor(user, lessor))
		}
	}
	c.JSON(http.StatusOK, gin.H{"users": usersDTO})
}
