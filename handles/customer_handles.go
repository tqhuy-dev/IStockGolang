package handles

import (
	"net/http"
	"github.com/labstack/echo"
	"github.com/tranhuy-dev/IStockGolang/database"
	"github.com/tranhuy-dev/IStockGolang/core/constant"
	models "github.com/tranhuy-dev/IStockGolang/models"
)
// Get customer
func GetCustomer(c echo.Context) error {
	customer := database.RetrieveAllCustomer()
	return c.JSON(http.StatusOK, models.SuccessReponse{
		Code: 200 ,
		Message: "Retrieve customers success" ,
		Data: customer})
}
// Create customer
func CreateCustomer(c echo.Context) error {
	var req models.CustomerReq
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, models.ErrorResponse{Code: constant.BadRequest, Message: "Bad parameter"})
	}
	customer := database.InsertCustomer(req)
	return c.JSON(http.StatusOK, models.SuccessReponse{Code: constant.Success , Message: "Create success" , Data:customer})
}

func UpdateCustomer(c echo.Context) error {
	var req models.CustomerReq
	idCustomer := c.Param("email")
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest , models.ErrorResponse{Code: constant.BadRequest , Message: "Bad request"})
	}
	updateResult := database.UpdateCustomer(req , idCustomer)

	if updateResult == 0  {
		return c.JSON(http.StatusNotFound , models.ErrorResponse{Code: constant.NotFound , Message:"Not found customer"})
	} else if updateResult == 1{
		return c.JSON(http.StatusOK , models.SuccessReponse{
			Code: constant.Success,
			Message:"update success"})
	}

	return c.JSON(http.StatusOK , updateResult) 
}

func DeleteCustomer(c echo.Context) error {
	idCustomer := c.Param("email")
	deleteResult := database.DeleteCustomer(idCustomer)
	return c.JSON(http.StatusOK , deleteResult)
}

func GetCustomerByEmail(c echo.Context) error {
	idCustomer := c.Param("email")
	customerData,err := database.FindUserByEmail(idCustomer)
	if err != nil {
		return c.JSON(http.StatusOK , models.ErrorResponse{Code: constant.NotFound, Message: err.Error()})
	}
	return c.JSON(http.StatusOK , customerData)
}