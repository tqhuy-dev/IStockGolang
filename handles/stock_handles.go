package handles

import (
	"net/http"
	"strconv"
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
	token := GetTokenHeader(c)
	stockStatusArr := []string{constant.STATUS_BLOCK, constant.STATUS_CLOSE, constant.STATUS_OPEN}
	data, err := mathematic.FindElementString(stockReq.Status, stockStatusArr)
	if err != nil {
		return c.JSON(http.StatusForbidden, models.ErrorResponse{
			Code:    constant.ExpectedError,
			Message: err.Error()})
	}
	stockReq.Status = data

	dataStock, err := database.CreateStock(token , stockReq)
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
	token := GetTokenHeader(c)
	dataStock, err := database.RetrieveStockUser(token , email)

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

func RetrieveStockByToken(c echo.Context) error {
	token := GetTokenHeader(c)
	dataStock, err := database.RetriveStockByToken(token)

	if err != nil {
		return c.JSON(http.StatusNetworkAuthenticationRequired, models.ErrorResponse{
			Code:    constant.Authentication,
			Message: err.Error()})
	}

	return c.JSON(http.StatusOK, models.SuccessReponse{
		Code:    constant.Success,
		Message: "Success",
		Data:    dataStock})
}

func UpdateStock(c echo.Context) error {
	token := GetTokenHeader(c)
	var stockUpdateBody models.Stock
	err := c.Bind(&stockUpdateBody)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Code:    constant.BadRequest,
			Message: "Bad parameters"})
	}

	idStock , err := strconv.Atoi(c.FormValue("stock_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Code:    constant.BadRequest,
			Message: "Bad parameters"})
	}
	updateStockData, errUpdateStock := database.UpdateStock(token , idStock , stockUpdateBody)
	if errUpdateStock != nil {
		return c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Code:    constant.BadRequest,
			Message: errUpdateStock.Error()})
	}

	return c.JSON(http.StatusOK, models.SuccessReponse{
		Code:    constant.Success,
		Message: "Success",
		Data:    updateStockData})
}
