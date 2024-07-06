package service

import (
	"PCS-API/models"
	"PCS-API/repository"
	"PCS-API/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stripe/stripe-go/v78"
	price2 "github.com/stripe/stripe-go/v78/price"
	"github.com/stripe/stripe-go/v78/product"
)

func serviceConvertToServiceDTO(service models.Service, userId uuid.UUID, date time.Time) models.ServiceDTO {
	return models.ServiceDTO{
		Service: service,
		UserId:  userId,
		Date:    date,
	}
}

func ServiceCreateNewService(c *gin.Context) {
	var service models.Service
	var err error
	if err = c.BindJSON(&service); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if service.Price < 1 ||
		(service.TargetCustomer != models.LessorType && service.TargetCustomer != models.TravelerType) ||
		service.RangeAction < 0 ||
		service.Name == "" ||
		service.Description == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "19"})
		return
	}
	service.Lat, service.Lon, err = utils.LocateWithAddress(service.Address, service.City, service.ZipCode, service.Country)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	idBrut, exist := c.Get("idUser")
	if !exist {
		c.JSON(http.StatusBadRequest, gin.H{"error": "8"})
		return
	}
	idUser, err := uuid.Parse(idBrut.(string))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "18"})
		return
	}

	//TODO: Penser à la sécurité (imaginons que le provider n'existe plus ?
	provider := repository.ProviderGetByUserId(idUser)
	service.ID = uuid.New()
	service.ProviderId = provider.ID

	// Put the price on Stripe

	stripe.Key = "sk_test_51PNwOpRrur5y60cs5Yv2aKu9v6SrJHigo2cLgmxevvozEfzSDWFnaQhMwVH02RLc8R2xHdjkJ6QagZ7KDyYTVxZt00gadizteA"

	prodParams := &stripe.ProductParams{
		Name: stripe.String(service.Name),
	}
	prod, err := product.New(prodParams)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "27"})
		return
	}

	stripe.Key = "sk_test_51PNwOpRrur5y60cs5Yv2aKu9v6SrJHigo2cLgmxevvozEfzSDWFnaQhMwVH02RLc8R2xHdjkJ6QagZ7KDyYTVxZt00gadizteA"

	priceParams := &stripe.PriceParams{
		Product:    stripe.String(prod.ID),
		UnitAmount: stripe.Int64(int64(service.Price * 100)),
		Currency:   stripe.String(string(stripe.CurrencyEUR)),
	}
	price, err := price2.New(priceParams)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"property": "27"})
		return
	}
	service.IdStripe = price.ID

	// Création de la prestation dans la base et renvoie à l'utilisateur

	service, err = repository.ServiceCreateNewService(service)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	serviceDTO := serviceConvertToServiceDTO(service, idUser, time.Time{})
	c.JSON(http.StatusOK, gin.H{"service": serviceDTO})
}

func ServiceUpdate(c *gin.Context) {
	idService, _ := uuid.Parse(c.Param("id"))
	service, err := repository.ServiceGetWithServiceId(idService)
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
	if service.ProviderId != repository.ProviderGetByUserId(idUser).ID {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "20"})
		return
	}

	var serviceTransfert models.Service
	if err = c.BindJSON(&serviceTransfert); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if serviceTransfert.Price < 1 ||
		(serviceTransfert.TargetCustomer != models.LessorType && serviceTransfert.TargetCustomer != models.TravelerType) ||
		serviceTransfert.RangeAction < 0 ||
		service.Name == "" ||
		serviceTransfert.Description == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "19"})
		return
	}
	serviceTransfert.Lat, serviceTransfert.Lon, err = utils.LocateWithAddress(service.Address, service.City, service.ZipCode, service.Country)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	serviceTransfert.ID = service.ID
	serviceTransfert.ProviderId = service.ProviderId
	serviceTransfert.IdStripe = service.IdStripe

	// Modification prix service

	if serviceTransfert.Price != service.Price {

		stripe.Key = "sk_test_51PNwOpRrur5y60cs5Yv2aKu9v6SrJHigo2cLgmxevvozEfzSDWFnaQhMwVH02RLc8R2xHdjkJ6QagZ7KDyYTVxZt00gadizteA"

		priceParams := &stripe.PriceParams{
			Product:    stripe.String(service.IdStripe),
			UnitAmount: stripe.Int64(int64(service.Price * 100)),
			Currency:   stripe.String(string(stripe.CurrencyEUR)),
		}
		_, err = price2.New(priceParams)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"property": "27"})
			return
		}
	}

	// Modification du service et renvoie à l'utilisateur

	serviceTransfert = repository.ServiceUpdate(serviceTransfert)
	ServiceDTO := serviceConvertToServiceDTO(serviceTransfert,
		repository.ProviderGetUserIdWithProviderId(serviceTransfert.ProviderId),
		time.Time{})
	c.JSON(http.StatusOK, gin.H{"service": ServiceDTO})
}

func ServiceGetAll(c *gin.Context) {
	services, err := repository.ServiceGetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"service": services})
}

// TODO: Risque de causer problème de clé étrangère lors de sa suppression

func ServiceDelete(c *gin.Context) {
	idService, _ := uuid.Parse(c.Param("id"))
	service, err := repository.ServiceGetWithServiceId(idService)
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
	if service.ProviderId != repository.ProviderGetByUserId(idUser).ID {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "20"})
		return
	}
	err = repository.ServiceDeleteById(idService)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}
