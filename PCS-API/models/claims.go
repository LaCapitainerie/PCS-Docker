package models

import "github.com/dgrijalva/jwt-go"

// Claims est la structure spécifiant les données utilisées pour les tokens de connexion
type Claims struct {
	IdUser string `json:"idUser"`
	jwt.StandardClaims
}
