package handles

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/tranhuy-dev/IStockGolang/core/constant"
	"github.com/tranhuy-dev/IStockGolang/database"
	models "github.com/tranhuy-dev/IStockGolang/models"
)

// Get customer
func GetCustomer(c echo.Context) error {
	customer := database.RetrieveAllCustomer()
	return c.JSON(http.StatusOK, models.SuccessReponse{
		Code:    200,
		Message: "Retrieve customers success",
		Data:    customer})
}

// Create customer
func CreateCustomer(c echo.Context) error {
	var req models.CustomerReq
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Code:    constant.BadRequest,
			Message: "Bad parameter"})
	}
	customer := database.InsertCustomer(req)
	return c.JSON(http.StatusOK, models.SuccessReponse{
		Code:    constant.Success,
		Message: "Create success",
		Data:    customer})
}

func UpdateCustomer(c echo.Context) error {
	var req models.CustomerReq
	idCustomer := c.Param("email")
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Code:    constant.BadRequest,
			Message: "Bad request"})
	}
	updateResult, err := database.UpdateCustomer(req, idCustomer)

	if err != nil {
		return c.JSON(http.StatusNotFound, models.ErrorResponse{
			Code:    constant.NotFound,
			Message: err.Error()})
	}
	return c.JSON(http.StatusOK, models.SuccessReponse{
		Code:    constant.Success,
		Message: "update success",
		Data:    updateResult})
}

func DeleteCustomer(c echo.Context) error {
	idCustomer := c.Param("email")
	deleteResult, err := database.DeleteCustomer(idCustomer)
	if err != nil {
		return c.JSON(http.StatusNotFound, models.ErrorResponse{
			Code:    constant.NotFound,
			Message: err.Error()})
	}
	return c.JSON(http.StatusOK, models.SuccessReponse{
		Code:    constant.Success,
		Message: "Delete success",
		Data:    deleteResult})
}

func GetCustomerByEmail(c echo.Context) error {
	idCustomer := c.Param("email")
	customerData, err := database.FindUserByEmail(idCustomer)
	if err != nil {
		return c.JSON(http.StatusOK, models.ErrorResponse{Code: constant.NotFound, Message: err.Error()})
	}
	return c.JSON(http.StatusOK, customerData)
}

func ChangePassword(c echo.Context) error {
	var req models.ChangePasswordReq
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Code:    constant.BadRequest,
			Message: "Bad request"})
	}

	dataChangePassword, err := database.ChangePassword(req)
	if err != nil {
		return c.JSON(http.StatusNotFound, models.ErrorResponse{
			Code:    401,
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, models.SuccessReponse{
		Code:    201,
		Message: "Change password success",
		Data: dataChangePassword})
}
