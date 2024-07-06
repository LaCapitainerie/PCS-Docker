package utils

import (
	"PCS-API/models"
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

// HashPassword génère le hash bcrypt d'un mot de passe
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// CreateToken génère un token à utiliser, prends en paramètre l'id d'un utilisateur et renvoie le token et l'erreur
func CreateToken(id string) (string, error) {
	expirationTime := time.Now().Add(TokenExpirationTime)
	claims := &models.Claims{
		IdUser: id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	return token.SignedString(TokenKey)
}

// CheckPassword vérifie si le hash et le mot de passe correspondent bien
func CheckPassword(hash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// IsInArrayString vérifie si un chaine ne se trouve pas dans un tableau de chaine
func IsInArrayString(chaine string, tab []string) bool {
	for _, str := range tab {
		if str == chaine {
			return true
		}
	}
	return false
}

func GenerateUniqueFileName(name string) string {
	timestamp := time.Now().UnixNano()
	randomPart := rand.Intn(1000)
	ext := filepath.Ext(name)
	return fmt.Sprintf("%d-%d%s", timestamp, randomPart, ext)
}

func LocateWithAddress(address string, city string, zipcode string, country string) (float64, float64, error) {
	params := strings.Join([]string{address, city, zipcode, country}, ",")
	url := strings.ReplaceAll(fmt.Sprintf("https://nominatim.openstreetmap.org/search?format=json&q=%s", params), " ", "%20")

	resp, err := http.Get(url)
	if err != nil {
		return 0, 0, err
	}
	defer resp.Body.Close()

	var locateJson []struct {
		Lat string `json:"lat"`
		Lon string `json:"lon"`
	}

	err = json.NewDecoder(resp.Body).Decode(&locateJson)
	if err != nil {
		return 0, 0, err
	}

	if len(locateJson) == 0 {
		return 0, 0, fmt.Errorf("aucune coords trouvé")
	}

	lat, err := strconv.ParseFloat(locateJson[0].Lat, 64)
	if err != nil {
		return 0, 0, err
	}
	lon, err := strconv.ParseFloat(locateJson[0].Lon, 64)
	if err != nil {
		return 0, 0, err
	}
	return lat, lon, nil
}

func DaysBetweenDates(date1, date2 time.Time) int {
	diff := date2.Sub(date1)
	return int(diff.Hours() / 24)
}
