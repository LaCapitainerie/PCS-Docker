package repository

import (
	"PCS-API/models"
	"PCS-API/utils"
)

func TicketGetAll() ([]models.Ticket, error) {
	var tickets []models.Ticket
	err := utils.DB.Find(&tickets).Error
	return tickets, err
}

func TicketUpdateById(ticket models.Ticket) (models.Ticket, error) {
	err := utils.DB.Save(&ticket).Error
	return ticket, err
}

func TicketCreate(ticket models.Ticket) (models.Ticket, error) {
	err := utils.DB.Create(&ticket).Error
	return ticket, err
}
