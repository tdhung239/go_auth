package middle_ware_login

import (
	"github.com/go_auth/services"
	"github.com/labstack/echo/v4"
)

func BasicAuth(email string, password string, c echo.Context) (bool, error) {
	check_admin, err := services.CheckAdmin(email, password, c)
	if err != nil {
		return false, nil
	}
	if check_admin {
		c.Set("email", email)
		c.Set("admin", true)
		return true, nil
	}

	c.Set("email", email)
	c.Set("admin", false)

	return true, nil
}
