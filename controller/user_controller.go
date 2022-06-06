package controller

import (
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go_auth/services"
	"github.com/golang/glog"
	"github.com/labstack/echo/v4"
)

func GetUser(c echo.Context) error {
	getUser, err := services.GetUser(c)
	if err != nil {
		glog.Error("get user error: %v", err)
		return err
	}
	return c.JSON(http.StatusOK, getUser)
}
