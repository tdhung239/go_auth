package services

import (
	"fmt"

	"github.com/beego/beego/v2/client/orm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go_auth/models"
	"github.com/go_auth/utils"
	"github.com/golang-jwt/jwt"
	"github.com/golang/glog"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func init() {
	// models define
	orm.RegisterModel(new(models.User))
}

func CheckAdmin(email string, password string, c echo.Context) (bool, error) {

	o := orm.NewOrm()
	user := &models.User{
		Email: email,
	}

	err := o.Read(user)
	if err != nil {
		glog.Error("get user error: %v", err)
		return false, err
	}

	bsp, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	err = bcrypt.CompareHashAndPassword(bsp, []byte(password))

	if err != nil {
		panic(err)
	} else {
		fmt.Println("password are equal")
	}

	if err != nil {
		panic(err)
	}

	if user.Admin {
		return true, err
	}

	return false, err
}

func GetUser(c echo.Context) (interface{}, error) {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	email := claims["email"].(string)

	info, err := utils.GetUserByEmail(email, c)

	if err != nil {
		glog.Error("get user error: %v", err)
		return info, err
	}

	return info, err
}
