package service

import (
	"PCS-API/models"
	"PCS-API/repository"
	"PCS-API/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func reservationDTOCreate(reservation models.Reservation, bill models.Bill, service []models.ServiceDTO) models.ReservationDTO {
	return models.ReservationDTO{
		Reservation: reservation,
		Bill:        bill,
		Service:     service,
	}
}

func reservationDateIntersect(entry models.Reservation, allEntry []models.Reservation) bool {
	for _, value := range allEntry {
		if !value.Annulation {
			if (entry.BeginDate.Before(value.EndDate) || entry.BeginDate.Equal(value.EndDate)) &&
				(entry.EndDate.After(value.BeginDate) || entry.EndDate.Equal(value.BeginDate)) {
				return true
			}
		}
	}
	return false
}

func ReservationPropertyCreate(c *gin.Context) string {

	str, exist := c.Get("idUser")
	if !exist {
		c.JSON(http.StatusBadRequest, gin.H{"error": "8"})
		return ""
	}

	var reservationDTO models.ReservationDTO
	var err error
	if err = c.BindJSON(&reservationDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return ""
	}

	idUser, _ := uuid.Parse(str.(string))
	if idUser != reservationDTO.TravelerId {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "18"})
		return ""
	}

	property, err := repository.PropertyGetById(reservationDTO.PropertyId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "21"})
		return ""
	}

	services, err := reservationGetAllService(&reservationDTO)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "25"})
		return ""
	}

	//TODO: Vérifie que la propriété est dans le rayon d'actiond de tous les services

	var reservation models.Reservation
	reservation.BeginDate = reservationDTO.BeginDate
	reservation.EndDate = reservationDTO.EndDate
	reservation.TravelerId = repository.TravelerGetIdByUserId(idUser)

	if reservation.TravelerId == uuid.Nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "24"})
		return ""
	}

	if reservation.BeginDate.After(reservation.EndDate) {
		tmp := reservation.EndDate
		reservation.EndDate = reservation.BeginDate
		reservation.BeginDate = tmp
	}

	timeNow := time.Now()
	date := time.Date(timeNow.Year(), timeNow.Month(), timeNow.Day(), 0, 0, 0, 0, timeNow.Location())

	if !reservation.BeginDate.After(date) {
		c.JSON(http.StatusConflict, gin.H{"error": "22"})
		return ""
	}

	if reservationDateIntersect(reservation, repository.ReservationGetAllByIdPropertyWithEndDateAfterADate(property.ID, date)) {
		c.JSON(http.StatusConflict, gin.H{"error": "22"})
		return ""
	}

	reservation.ID = uuid.New()
	reservation.PropertyId = property.ID
	reservation.Annulation = true

	reservation, err = repository.ReservationCreate(reservation)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return ""
	}

	serviceDTO, err := reservationServiceListCreate(&reservationDTO, services, &reservation.ID)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "25"})
		return ""
	}

	bill, err := billCreate(property, reservation)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "23"})
		return ""
	}

	reservationDTO = reservationDTOCreate(reservation, bill, serviceDTO)
	return reservationDTO.ID.String()
}

func ReservationValidationPaiement(c *gin.Context) {
	idReservation, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "8"})
		return
	}
	_, err = repository.ReservationValidation(idReservation)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "27"})
		return
	}
	reservationDTO := reservationGetById(c, idReservation.String())
	c.JSON(http.StatusOK, gin.H{"reservation": reservationDTO})
}

func reservationGetById(c *gin.Context, str string) models.ReservationDTO {
	idReservation, err := uuid.Parse(str)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "8"})
		return models.ReservationDTO{}
	}
	var reservationDTO models.ReservationDTO
	reservationDTO.Reservation, err = repository.ReservationGetById(idReservation)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "29"})
		return models.ReservationDTO{}
	}
	reservationDTO.Bill, err = repository.BillGetByReservationId(reservationDTO.ID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "29"})
		return models.ReservationDTO{}
	}
	reservationDTO.Service, err = repository.ReservationServiceGetAllByAReservationId(reservationDTO.ID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "29"})
		return models.ReservationDTO{}
	}
	return reservationDTO
}

func ReservationGetById(c *gin.Context) {
	str, exist := c.Get("id")
	if !exist {
		c.JSON(http.StatusBadRequest, gin.H{"error": "8"})
		return
	}
	dto := reservationGetById(c, str.(string))
	c.JSON(http.StatusOK, gin.H{"reservation": dto})
}

func ReservationGetAllOfAProperty(c *gin.Context) {
	idProperty, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "8"})
		return
	}
	reservations, err := repository.ReservationGetAllByIdProperty(idProperty)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "30"})
		return
	}
	reservationsDTO := make([]models.ReservationDTO, len(reservations))
	for i, reservation := range reservations {
		reservationsDTO[i].Reservation = reservation
		reservationsDTO[i].Bill, _ = repository.BillGetByReservationId(reservation.ID)
		reservationsDTO[i].Service, _ = repository.ReservationServiceGetAllByAReservationId(reservation.ID)
	}
	c.JSON(http.StatusOK, gin.H{"reservation": reservationsDTO})
}

func ReservationPropertyAnnulationWithAId(c *gin.Context) {
	idReservation, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "29"})
		return
	}

	reservation := reservationGetById(c, idReservation.String())
	//TODO: Placer une sécurité ici

	err = repository.ReservationSetAnnulation(idReservation)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "29"})
		return
	}

	reservation.Annulation = true
	c.JSON(http.StatusOK, gin.H{"reservation": reservation})
}

func ReservationPropertyReportReservation(c *gin.Context) {
	var modificationInput models.ReservationDTO
	var err error
	if err = c.BindJSON(&modificationInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	reservationDTO := reservationGetById(c, modificationInput.ID.String())
	if utils.DaysBetweenDates(modificationInput.BeginDate, modificationInput.EndDate) != utils.DaysBetweenDates(reservationDTO.BeginDate, reservationDTO.EndDate) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "31"})
		return
	}

	timeNow := time.Now()
	date := time.Date(timeNow.Year(), timeNow.Month(), timeNow.Day(), 0, 0, 0, 0, timeNow.Location())

	if !modificationInput.BeginDate.After(date) {
		c.JSON(http.StatusConflict, gin.H{"error": "22"})
		return
	}
	reservationOg := models.Reservation{
		BeginDate: modificationInput.BeginDate,
		EndDate:   modificationInput.EndDate,
	}
	if reservationDateIntersect(reservationOg, repository.ReservationGetAllByIdPropertyWithEndDateAfterADate(reservationDTO.PropertyId, date)) {
		c.JSON(http.StatusConflict, gin.H{"error": "22"})
		return
	}

	reservationDTO.BeginDate = modificationInput.BeginDate
	reservationDTO.EndDate = modificationInput.EndDate
	err = repository.ReservationSetReport(reservationDTO.ID, reservationDTO.BeginDate, reservationDTO.EndDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "30"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"reservation": reservationDTO})
}
