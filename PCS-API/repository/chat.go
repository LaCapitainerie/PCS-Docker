package repository

import (
	"PCS-API/models"
	"PCS-API/utils"

	"github.com/google/uuid"
)

func VerifyExistenceChat(users []string) (string, error) {
	var chatId string
	subReq := utils.DB.Model(&models.ChatUser{}).
		Select("user_id").
		Where("user_id IN ?", users)

	err := utils.DB.Model(&models.ChatUser{}).
		Select("chat_id").
		Where("user_id IN (?)", subReq).
		Order("chat_id").
		Limit(1).
		Pluck("chat_id", &chatId).
		Error
	return chatId, err
}

func CreateChat(chat models.Chat, users []models.ChatUser) (models.Chat, error) {
	result := utils.DB.Create(&chat)
	if result.Error != nil {
		return chat, result.Error
	}
	for i := range users {
		result = utils.DB.Create(&users[i])
		if result.Error != nil {
			return chat, result.Error
		}
	}
	return chat, nil
}

func CreateMessage(message models.Message) (models.Message, error) {
	result := utils.DB.Create(&message)
	if result.Error != nil {
		return message, result.Error
	}
	return message, nil
}

func VerifyExistenceUserInAChat(idUser string, idChat string) bool {
	var count int64
	result := utils.DB.Model(&models.ChatUser{}).
		Where("chat_id = ? AND user_id = ?", idChat, idUser).
		Count(&count)
	if result.Error != nil {
		return false
	}
	return count > 0
}

func GetChat(idChat string) (models.Chat, error) {
	var chat models.Chat
	err := utils.DB.First(&chat, idChat).Error
	return chat, err
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

func UserGetByIdComplet(userid uuid.UUID) models.UsersDTO {

	var userDTO models.UsersDTO
	user, _ := UsersGetById(userid)

	switch user.Type {
	case models.TravelerType:
		provider := ProviderGetByUserId(user.ID)
		userDTO = CreateUserDTOwithUserAndProvider(user, provider)
	case models.ProviderType:
		traveler := TravelerGetByUserId(user.ID)
		userDTO = CreateUserDTOwithUserAndTraveler(user, traveler)
	case models.LessorType:
		lessor := LessorGetByUserId(user.ID)
		userDTO = CreateUserDTOwithUserAndLessor(user, lessor)
	}
	userDTO.Password = ""

	return userDTO
}

func GetAllChatUserOfAChat[R []models.UsersDTO](idChat string) R {
	var chatUsers []models.ChatUser

	utils.DB.Where("chat_id = ?", idChat).Find(&chatUsers)
	userId := make(R, len(chatUsers))
	for i := range chatUsers {
		userId[i] = UserGetByIdComplet(chatUsers[i].UserID)
	}
	return userId
}

func GetAllMessageOfAChat(idChat string) ([]models.Message, error) {
	var message []models.Message
	err := utils.DB.Where("chat_id", idChat).Find(&message).Error
	return message, err
}

func GetTicketOfAChat(idChat string) (models.Ticket, error) {
	var ticket models.Ticket
	err := utils.DB.Where("chat_id", idChat).First(&ticket).Error
	return ticket, err
}

func GetEverythingAboutAChat(idChat string) struct {
	Chat     models.Chat
	Users    []models.Users
	Messages []models.Message
	Tickets  models.Ticket
} {
	var result struct {
		Chat     models.Chat
		Users    []models.Users
		Messages []models.Message
		Tickets  models.Ticket
	}
	var chatUsers []models.ChatUser

	utils.DB.Where("id = ?", idChat).First(&result.Chat)
	utils.DB.Where("chat_id = ?", idChat).Find(&chatUsers)

	result.Users = make([]models.Users, len(chatUsers))
	for i, value := range chatUsers {
		result.Users[i], _ = UsersGetById(value.UserID)
	}
	utils.DB.Where("chat_id = ?", idChat).Find(&result.Messages).Joins("JOIN user ON message.user_id = user.id as user")
	utils.DB.Where("chat_id = ?", idChat).Find(&result.Tickets)

	return result
}

func GetAllChatByUser(id string) []models.ChatUser {
	var chatUsers []models.ChatUser
	utils.DB.Where("user_id = ?", id).Find(&chatUsers)
	return chatUsers
}

func chatUserDeleteByIdUser(idUser uuid.UUID) {
	utils.DB.Where("id_user = ?", idUser).Delete(&models.ChatUser{})
}
