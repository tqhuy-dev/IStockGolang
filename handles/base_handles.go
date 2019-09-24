package handles

import (
	"github.com/labstack/echo"
)

func GetTokenHeader(c echo.Context) string {
	token := c.Request().Header.Get("token")
	return token
}