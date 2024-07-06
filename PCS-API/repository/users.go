// Package repository
// package spécifiant les fonctions utilisé pour les requêtes avec la base de donnée
package repository

import (
	"PCS-API/models"
	"PCS-API/utils"
	"time"

	"github.com/google/uuid"
)

// CreateUser reçoit en argument un user
// Crée un "users" dans la table et renvoie l'user mis à jour
func CreateUser(users models.Users) (models.Users, error) {
	err := utils.DB.Save(&users)
	return users, err.Error
}

// UsersVerifyEmail reçoit en argument un string
// Vérifie dans la base de donnée si le mail dans user existe déjà
func UsersVerifyEmail(mail string) bool {
	var count int64
	utils.DB.Model(&models.Users{}).Where("mail = ?", mail).Count(&count)
	return count > 0
}

// UsersVerifyPhone reçoit en argument un string
// Vérifie dans la base de donnée si le numéro de téléphone dans user existe déjà
func UsersVerifyPhone(phoneNumber string) bool {
	var count int64
	utils.DB.Model(&models.Users{}).Where("phone_number = ?", phoneNumber).Count(&count)
	return count > 0
}

// UsersLoginVerify reçoit en argument un email
// Vérifie les informations de connexion, renvoie l'utilisateur en question
func UsersLoginVerify(mail string) models.Users {
	var user models.Users
	utils.DB.Where("mail = ?", mail).First(&user)
	return user
}

func UsersGetById(id uuid.UUID) (models.Users, error) {
	var user models.Users
	err := utils.DB.First(&user, id).Error
	return user, err
}

func UsersDelete(user models.Users) error {
	chatUserDeleteByIdUser(user.ID)
	return utils.DB.Delete(models.Users{}, user.ID).Error
}

func UsersUpdateLastConnectionDate(id uuid.UUID) {
	utils.DB.Model(&models.Users{}).Where("id = ?", id).Update("LastConnectionDate", time.Now())
}

func UsersGetTypeById(id uuid.UUID) string {
	var user models.Users
	utils.DB.First(&user, id)
	return user.Type
}

func UsersGetAll(limit, offset int) ([]models.Users, error) {
	var users []models.Users
	err := utils.DB.Limit(limit).Offset(offset).Find(&users).Error
	return users, err
}
