package handles

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/tranhuy-dev/IStockGolang/core/constant"
	"github.com/tranhuy-dev/IStockGolang/core/mathematic"
	"github.com/tranhuy-dev/IStockGolang/database"
	models "github.com/tranhuy-dev/IStockGolang/models"
)

func CreateStockHandles(c echo.Context) error {
	var stockReq models.Stock
	err := c.Bind(&stockReq)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Code:    constant.BadRequest,
			Message: "Bad parameters"})
	}
	stockStatusArr := []string{constant.STATUS_BLOCK, constant.STATUS_CLOSE, constant.STATUS_OPEN}
	data, err := mathematic.FindElementString(stockReq.Status, stockStatusArr)
	if err != nil {
		return c.JSON(http.StatusForbidden, models.ErrorResponse{
			Code:    constant.ExpectedError,
			Message: err.Error()})
	}
	stockReq.Status = data

	if stockReq.Customer == "" {
		return c.JSON(http.StatusBadRequest , models.ErrorResponse{
			Code: constant.BadRequest,
			Message: "Customer is required"})
	}

	dataStock, err := database.CreateStock(stockReq)
	if err != nil {
		return c.JSON(http.StatusForbidden, models.ErrorResponse{
			Code:    constant.ExpectedError,
			Message: err.Error()})
	}

	return c.JSON(http.StatusOK, models.SuccessReponse{
		Code:    constant.Success,
		Message: "Add stock success",
		Data:    dataStock})
}

func RetriveStockByEmail(c echo.Context) error {
	email := c.FormValue("email")

	dataStock, err := database.RetrieveStockUser(email)

	if err != nil {
		return c.JSON(http.StatusNotFound, models.ErrorResponse{
			Code:    constant.NotFound,
			Message: constant.MessageUserNotFound})
	}

	return c.JSON(http.StatusOK, models.SuccessReponse{
		Code:    constant.Success,
		Message: "Retrice success",
		Data:    dataStock})
}
