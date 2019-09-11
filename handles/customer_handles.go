package handles

import (
	"net/http"

	"github.com/labstack/echo"
)
// Get customer
func GetCustomer(c echo.Context) error {
	return c.String(http.StatusOK, "Get customer")
}
// Create customer
func CreateCustomer(c echo.Context) error {
	return c.String(http.StatusOK, "Create customer")
}