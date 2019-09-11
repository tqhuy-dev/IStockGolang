package handles

import (
	"net/http"
	"github.com/labstack/echo"
	"github.com/tranhuy-dev/IStockGolang/database"
)
// Get customer
func GetCustomer(c echo.Context) error {
	return c.String(http.StatusOK, "Get customer")
}
// Create customer
func CreateCustomer(c echo.Context) error {
	customer := database.InsertCustomer()
	return c.JSON(http.StatusOK, customer)
}