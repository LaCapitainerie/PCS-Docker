// Package utils spécifie toutes les fonctions utilitaire à l'API
package utils

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	"github.com/stripe/stripe-go/v78"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB                       // Contient toutes les données pour l'interaction avec la base de données
var PortApp string                    // Contient le port utilisé par l'API
var TokenExpirationTime time.Duration // Temps d'expiration d'un token généré en heure
var TokenKey []byte

// LoadConfig
// Charge toutes les données nécessaire à l'API
// Les variables d'environnement et la connexion avec la bdd
func LoadConfig() {
	err := godotenv.Load("config.env")
	if err != nil {
		log.Fatal("The \"config.env\" file is invalid (please rename 'config.example.env' to 'config.env'\nIf you want to use docker, please insert '--env-file config.env' in docker command")
		os.Exit(1)
	}

	dsn := "host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s"
	dsn = fmt.Sprintf(dsn, os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"), os.Getenv("DB_PORT"), os.Getenv("DB_SSLMODE"), os.Getenv("DB_TIMEZONE"))

	log.Println("Connexion à la base de donnée...")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Database connection error, please recheck the information in the config.env")
		os.Exit(2)
	}

	PortApp = os.Getenv("PORT_APP")
	stripe.Key = os.Getenv("STRIPE_KEY")

	expirationTime, err := strconv.Atoi(os.Getenv("TOKEN_EXPIRATION"))
	if err != nil {
		log.Fatal("token expiration time error, please recheck the information in the config.env")
		os.Exit(3)
	}
	TokenExpirationTime = time.Hour * time.Duration(expirationTime)

	key := os.Getenv("TOKEN_KEY")
	if key == "" {
		log.Fatal("token key error, please recheck the information in the config.env")
		os.Exit(4)
	}
	TokenKey = []byte(key)
}
