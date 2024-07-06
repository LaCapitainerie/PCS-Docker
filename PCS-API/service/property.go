package service

import (
	"PCS-API/models"
	"PCS-API/repository"
	"PCS-API/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stripe/stripe-go/v78"
	price2 "github.com/stripe/stripe-go/v78/price"
	"github.com/stripe/stripe-go/v78/product"
)

//TODO: Problème d'autorisation à gérer dans le service property
//TODO: Clean le code

// @BasePath /api/v1

// GetAllProperty Récupère la liste de tous les Property
// @Summary Property
// @Schemes
// @Description Récupère tous les Property
// @Tags administration
// @Produce json
// @Success 200 {array} models.Property
// @Router /api/Property [get]
func GetAllProperty(c *gin.Context) {
	Propertys := repository.GetAllProperty()

	var propertyDTO []models.PropertyDTO
	for _, v := range Propertys {
		//TODO: On peut optimiser ça
		propertyImage := repository.PropertyImageGetAllByIdProperty(v.ID)
		lessor := repository.LessorGetById(v.LessorId)
		propertyDTO = append(propertyDTO, createPropertyDTOwithProperty(v, propertyImage, lessor.UserId))
	}

	c.JSON(http.StatusOK, gin.H{"Property": propertyDTO})
}

func PostAProperty(c *gin.Context) {
	var err error
	var propertyDTO models.PropertyDTO
	if err = c.BindJSON(&propertyDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	idBrut, exist := c.Get("idUser")
	if !exist {
		c.JSON(http.StatusBadRequest, gin.H{"error": "8"})
		return
	}
	idUser, _ := uuid.Parse(idBrut.(string))

	if !repository.IsALessor(idUser) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "14"})
		return
	}

	if len(propertyDTO.Name) < 1 ||
		len(propertyDTO.Type) < 1 ||
		propertyDTO.Price < 1 ||
		propertyDTO.Surface < 8 ||
		propertyDTO.Room < 1 ||
		len(propertyDTO.ZipCode) < 5 ||
		len(propertyDTO.Address) < 1 ||
		len(propertyDTO.City) < 1 ||
		len(propertyDTO.Country) < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "15"})
		return
	}

	// create Property

	var property models.Property
	property.ID = uuid.New()
	property.Name = propertyDTO.Name
	property.Type = propertyDTO.Type
	property.Price = propertyDTO.Price
	property.Surface = propertyDTO.Surface
	property.Room = propertyDTO.Room
	property.Bathroom = propertyDTO.Bathroom
	property.Garage = propertyDTO.Garage
	property.Description = propertyDTO.Description
	property.Address = propertyDTO.Address
	property.City = propertyDTO.City
	property.ZipCode = propertyDTO.ZipCode
	property.Country = propertyDTO.Country
	property.LessorId = repository.GetLessorIdByUserId(idUser)
	property.Lat, property.Lon, err = utils.LocateWithAddress(
		property.Address,
		property.City,
		property.ZipCode,
		property.Country)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	stripe.Key = "sk_test_51PNwOpRrur5y60cs5Yv2aKu9v6SrJHigo2cLgmxevvozEfzSDWFnaQhMwVH02RLc8R2xHdjkJ6QagZ7KDyYTVxZt00gadizteA"

	// Put the price on Stripe
	prodParams := &stripe.ProductParams{
		Name: stripe.String(property.Name),
	}
	prod, err := product.New(prodParams)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "26"})
		return
	}

	stripe.Key = "sk_test_51PNwOpRrur5y60cs5Yv2aKu9v6SrJHigo2cLgmxevvozEfzSDWFnaQhMwVH02RLc8R2xHdjkJ6QagZ7KDyYTVxZt00gadizteA"

	priceParams := &stripe.PriceParams{
		Product:    stripe.String(prod.ID),
		UnitAmount: stripe.Int64(int64(property.Price * 100)),
		Currency:   stripe.String(string(stripe.CurrencyEUR)),
	}
	price, err := price2.New(priceParams)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "26"})
		return
	}
	property.IdStripe = price.ID

	property, err = repository.PropertyCreate(property)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Property non créée"})
		return
	}

	// image

	var images []models.PropertyImage
	for _, value := range propertyDTO.Images {
		var image models.PropertyImage
		image.ID = uuid.New()
		image.Path = value
		image.PropertyId = property.ID
		image = repository.PropertyImageCreate(image)
		images = append(images, image)
	}

	// DTO Création - Rendue1
	propertyDTO = createPropertyDTOwithProperty(property, []models.PropertyImage{}, idUser)
	c.JSON(http.StatusOK, gin.H{"property": propertyDTO})
}

func createPropertyDTOwithProperty(property models.Property, images []models.PropertyImage, idUser uuid.UUID) models.PropertyDTO {
	imagesPath := make([]string, len(images))
	for i, v := range images {
		imagesPath[i] = v.Path
	}

	return models.PropertyDTO{
		ID:                      property.ID,
		IdStripe:                property.IdStripe,
		Name:                    property.Name,
		Type:                    property.Type,
		Price:                   property.Price,
		Surface:                 property.Surface,
		Room:                    property.Room,
		Bathroom:                property.Bathroom,
		Garage:                  property.Garage,
		Description:             property.Description,
		Address:                 property.Address,
		City:                    property.City,
		ZipCode:                 property.ZipCode,
		Lon:                     property.Lon,
		Lat:                     property.Lat,
		Images:                  imagesPath,
		Country:                 property.Country,
		AdministratorValidation: property.AdministratorValidation,
		UserId:                  idUser,
		LessorId:                property.LessorId,
	}
}

func PropertyDeleteById(c *gin.Context) {
	IDUSER, exist := c.Get("idUser")
	idUser, _ := uuid.Parse(IDUSER.(string))
	idProperty, _ := uuid.Parse(c.Param("id"))
	if !exist {
		c.JSON(http.StatusBadRequest, gin.H{"error": "8"})
		return
	}

	lessor := repository.LessorGetByUserId(idUser)
	supp := repository.PropertyDeleteWithIdUserAndPropertyId(idProperty, lessor.ID)
	if supp != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": supp.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

func GetPropertyById(c *gin.Context) {
	idProperty, _ := uuid.Parse(c.Param("id"))
	property, err := repository.PropertyGetById(idProperty)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{})
		return
	}

	propertyImages := repository.PropertyImageGetAllByIdProperty(idProperty)
	userId := repository.LessorGetById(property.LessorId).UserId

	propertyDTO := createPropertyDTOwithProperty(property, propertyImages, userId)
	c.JSON(http.StatusOK, gin.H{"property": propertyDTO})
}

func PutPropertyById(c *gin.Context) {
	idProperty, _ := uuid.Parse(c.Param("id"))
	propertyOrigin, err := repository.PropertyGetById(idProperty)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{})
		return
	}

	var propertyDTO models.PropertyDTO
	if err = c.BindJSON(&propertyDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	idBrut, exist := c.Get("idUser")
	if !exist {
		c.JSON(http.StatusBadRequest, gin.H{"error": "8"})
		return
	}
	idUser, _ := uuid.Parse(idBrut.(string))

	if !repository.IsALessor(idUser) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "14"})
		return
	}

	if len(propertyDTO.Name) < 1 ||
		len(propertyDTO.Type) < 1 ||
		propertyDTO.Price < 1 ||
		propertyDTO.Surface < 8 ||
		propertyDTO.Room < 1 ||
		len(propertyDTO.ZipCode) < 5 ||
		len(propertyDTO.Address) < 1 ||
		len(propertyDTO.City) < 1 ||
		len(propertyDTO.Country) < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "15"})
		return
	}

	var property models.Property
	property.ID = propertyOrigin.ID
	property.IdStripe = propertyOrigin.IdStripe
	property.Name = propertyDTO.Name
	property.Type = propertyDTO.Type
	property.Surface = propertyDTO.Surface
	property.Room = propertyDTO.Room
	property.Bathroom = propertyDTO.Bathroom

	property.Price = propertyDTO.Price

	property.Garage = propertyDTO.Garage
	property.Description = propertyDTO.Description
	property.Address = propertyDTO.Address
	property.City = propertyDTO.City
	property.ZipCode = propertyDTO.ZipCode
	property.Country = propertyDTO.Country
	property.LessorId = repository.GetLessorIdByUserId(idUser)
	property.Lat, property.Lon, err = utils.LocateWithAddress(
		property.Address,
		property.City,
		property.ZipCode,
		property.Country)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if propertyOrigin.Price != propertyDTO.Price {

		stripe.Key = "sk_test_51PNwOpRrur5y60cs5Yv2aKu9v6SrJHigo2cLgmxevvozEfzSDWFnaQhMwVH02RLc8R2xHdjkJ6QagZ7KDyYTVxZt00gadizteA"

		priceParams := &stripe.PriceParams{
			Product:    stripe.String(property.IdStripe),
			UnitAmount: stripe.Int64(int64(propertyDTO.Price * 100)),
			Currency:   stripe.String(string(stripe.CurrencyEUR)),
		}
		_, err = price2.New(priceParams)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"property": "26"})
			return
		}
	}

	property, err = repository.PropertyUpdate(property)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Property non mis à jour"})
		return
	}

	// Vérification ajout
	imagesOrigin := repository.PropertyImageGetAllByIdProperty(property.ID)
	pathOrigin := propertyImageGetArrayPathFromArray(imagesOrigin)
	var propertyImage []models.PropertyImage
	for i, value := range propertyDTO.Images {
		if utils.IsInArrayString(value, pathOrigin) {
			propertyImage = append(propertyImage, imagesOrigin[i])
			continue
		}
		var image models.PropertyImage
		image.ID = uuid.New()
		image.Path = value
		image.PropertyId = property.ID
		image = repository.PropertyImageCreate(image)
		propertyImage = append(propertyImage, image)
	}
	propertyImageClean(propertyImage, property.ID)

	// Création DTO
	propertyDTO = createPropertyDTOwithProperty(property, propertyImage, idUser)
	c.JSON(http.StatusOK, gin.H{"property": propertyDTO})
}
