package controller

import (
	"log"
	"net/http"

	"github.com/go_auth/models"
	"github.com/go_auth/services"
	"github.com/labstack/echo/v4"
)

func Login(c echo.Context) error {
	token, err := services.Login(c)

	if err != nil {
		log.Printf("Signed token err %v\n", err)
		return err
	}
	return c.JSON(http.StatusOK, &models.LoginRepon{
		Token: token,
	})
}
