package controller

import (
	"PCS-API/middleware"
	"PCS-API/service"

	"github.com/gin-gonic/gin"
)

// Chat # Controller
//
// Chat réceptionne toutes les requêtes ayant pour endpoint '/chat' demande en header un token
// Il les envoie aux fonctions services liés
func Chat(api *gin.RouterGroup) {
	chat := api.Group("/chat")
	chat.Use(middleware.AuthMiddleware())
	{
		chat.POST("", service.ChatPostMessage)
		chat.GET("/:id", service.ChatGetAllMessages)
		chat.GET("/allchatbyuser", service.GetAllChatByUser)
		/*		chat.GET("/connect", service.GetChatConnect)*/
	}
}
