package utils

import (
	"github.com/beego/beego/v2/client/orm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go_auth/models"
	"github.com/golang/glog"
	"github.com/labstack/echo/v4"
)

func GetUserByEmail(email string, c echo.Context) (interface{}, error) {
	o := orm.NewOrm()
	user := &models.User{
		Email: email,
	}

	err := o.Read(user)
	if err != nil {
		glog.Error("get user error: %v", err)
		return user, err
	}

	return user, err
}
