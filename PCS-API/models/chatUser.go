package models

import (
	"github.com/google/uuid"
)

// ChatUser est la structure spécifiant les données reliant un chat et un user
type ChatUser struct {
	User   UsersDTO  `gorm:"foreignKey:UserID" json:"user"`
	UserID uuid.UUID `gorm:"type:uuid;primaryKey;foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"userId"`
	ChatID uuid.UUID `gorm:"type:uuid;primaryKey;foreignKey:ChatID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"chatId"`
}

// TableName ChatUser Spécifie à gorm le nom de la table
func (ChatUser) TableName() string {
	return "chat_user"
}
