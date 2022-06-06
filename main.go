package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go_auth/configs"
	"github.com/go_auth/controller"
	"github.com/go_auth/middle_ware_login"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func init() {
	configs.InitDatabase()
}

func main() {

	server := echo.New()

	server.Use(middleware.Logger())

	isLogIn := middleware.JWT([]byte("mysecretkey"))
	isAdmin := middle_ware_login.IsAdmin

	server.POST("/login", controller.Login, middleware.BasicAuth(middle_ware_login.BasicAuth))

	groupNormalUser := server.Group("api/user")
	groupNormalUser.GET("/get", controller.GetUser, isLogIn)

	groupAminUser := server.Group("api/admin")
	groupAminUser.GET("/get", controller.GetUser, isLogIn, isAdmin)

	server.Logger.Fatal(server.Start(":8888"))
}
