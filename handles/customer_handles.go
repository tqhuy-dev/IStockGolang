package handles

import (
	"net/http"
	"github.com/labstack/echo"
	"github.com/tranhuy-dev/IStockGolang/database"
	models "github.com/tranhuy-dev/IStockGolang/models"
)
// Get customer
func GetCustomer(c echo.Context) error {
	customer := database.RetrieveAllCustomer()
	return c.JSON(http.StatusOK, customer)
}
// Create customer
func CreateCustomer(c echo.Context) error {
	var req models.CustomerReq
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, models.ErrorResponse{Code: 404, Message: "Bad parameter"})
	}
	customer := database.InsertCustomer(req)
	return c.JSON(http.StatusOK, models.SuccessReponse{Code: 200 , Message: "Create success" , Data:customer})
}

func UpdateCustomer(c echo.Context) error {
	updateResult := database.UpdateCustomer()
	return c.JSON(http.StatusOK , updateResult)
}