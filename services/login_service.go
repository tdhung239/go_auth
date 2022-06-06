package services

import (
	"log"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func Login(c echo.Context) (string, error) {
	email := c.Get("email").(string)
	admin := c.Get("admin").(bool)

	hash_token := jwt.New(jwt.SigningMethodHS256)
	claims := hash_token.Claims.(jwt.MapClaims)
	claims["email"] = email
	claims["admin"] = admin
	claims["exp"] = time.Now().Add(30 * time.Minute).Unix()

	token, err := hash_token.SignedString([]byte("mysecretkey"))

	if err != nil {
		log.Printf("Signed token err %v\n", err)
		return token, err
	}
	return token, err
}
