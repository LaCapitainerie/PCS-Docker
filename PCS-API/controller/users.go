// Package controller
// package contenant le code qui réceptionne toutes les requêtes, les envoie au "middleware" et aux fonctions services associées
package controller

import (
	"PCS-API/middleware"
	"PCS-API/service"

	"github.com/gin-gonic/gin"
)

// Users réceptionne toutes les requêtes ayant pour endpoint '/users'
// Il les envoie aux fonctions services liées
func Users(api *gin.RouterGroup) {
	api.POST("/user/register", service.CreateUser)
	api.POST("/user/login", service.LoginUser)
	api.GET("/user/id/:id", service.UserGetById)
	api.GET("/user/all", service.UserGetAll)

	userManagement := api.Group("/user/management")
	userManagement.Use(middleware.AuthMiddleware())
	{
		userManagement.DELETE("/:id", service.UserDeleteById)
		userManagement.PUT("/:id", service.UserUpdateById)
	}
}
