package service

import (
	"PCS-API/models"
	"PCS-API/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

func TicketGetAll(c *gin.Context) {
	tickets, err := repository.TicketGetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}
	chatDTO := make([]models.ChatDTO, len(tickets))
	for i, ticket := range tickets {
		result := repository.GetEverythingAboutAChat(ticket.ChatId.String())
		chatDTO[i] = createChatDTOWithAttribut(result.Chat, result.Tickets, result.Users, result.Messages)
	}
	c.JSON(http.StatusOK, gin.H{"chat": chatDTO})
}

func TicketUpdateById(c *gin.Context) {
	var err error

	var ticket models.Ticket
	if err = c.BindJSON(&ticket); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ticket, err = repository.TicketUpdateById(ticket)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"ticket": ticket})
}

func TicketCreate(c *gin.Context) {

	// JSON send is only

	var err error

	// Parse json from the body to a ticket struct
	var ticket models.Ticket
	if err = c.BindJSON(&ticket); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create the ticket in the database

	ticket, err = repository.TicketCreate(ticket)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	// Return the created ticket
	c.JSON(http.StatusOK, gin.H{"ticket": ticket})
}

/*
func GetEverythingAboutAChat(idChat string) struct {
	Chat      models.Chat
	ChatUsers []models.ChatUser
	Messages  []models.Message
	Tickets   models.Ticket
} {
*/
