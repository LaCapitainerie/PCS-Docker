// Package middleware
// package contenant les middleware soit un code s'executant avant le controller pour vérifier la validité de la requête
package middleware

import (
	"PCS-API/models"
	"PCS-API/repository"
	"PCS-API/utils"
	"errors"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// CORS édite le header HTTP pour accepter les requêtes venant d'autres serveurs/proxy source.
func CORS(router *gin.Engine) {
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "UPDATE", "PATCH"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
}

// AuthMiddleware Vérifie le token de connexion pour accèder à certaines parties de l'API
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")

		claims := &models.Claims{}

		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return utils.TokenKey, nil
		})

		if err != nil {
			if errors.Is(err, jwt.ErrSignatureInvalid) {
				c.AbortWithStatus(http.StatusUnauthorized)
				return
			}
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		if !token.Valid {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		idUser, err := uuid.Parse(claims.IdUser)
		if err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		/*		logEntry := models.Log{
					ID:       uuid.New(),
					UserID:   idUser,
					Action:   c.Request.Method,
					Endpoint: c.Request.URL.Path,
					Time:     time.Now(),
				}
				repository.CreateLogEntry(logEntry)*/

		repository.UsersUpdateLastConnectionDate(idUser)
		c.Set("idUser", claims.IdUser)
		c.Next()
	}
}

func BlockTypeMiddleware(userType string) gin.HandlerFunc {
	return func(c *gin.Context) {
		idBrut, exist := c.Get("idUser")
		if !exist {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		idUser, _ := uuid.Parse(idBrut.(string))
		typeUser := repository.UsersGetTypeById(idUser)
		if typeUser != userType {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		c.Next()
	}
}
