package handles

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/tranhuy-dev/IStockGolang/core/constant"
	"github.com/tranhuy-dev/IStockGolang/core/mathematic"
	"github.com/tranhuy-dev/IStockGolang/database"
	"github.com/tranhuy-dev/IStockGolang/models"
	// "fmt"
)

func AddProductionHandles(c echo.Context) error {
	token := GetTokenHeader(c)
	idStock, _ := strconv.Atoi(c.Param("stock"))
	var newProduction models.Production
	ProductionArr := []string {constant.STATUS_PROD_BLOCK,constant.STATUS_PROD_CLOSED,constant.STATUS_PROD_OPEN}
	err := c.Bind(&newProduction)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Code:    constant.BadRequest,
			Message: "Bad parameters"})
	}

	if newProduction.Status != "" {
		dataStatus , errStatus := mathematic.FindElementString(newProduction.Status , ProductionArr)
		if errStatus != nil {
			return c.JSON(http.StatusBadRequest, models.ErrorResponse{
				Code:    constant.BadRequest,
				Message: "Wrong status"})
		}
		newProduction.Status = dataStatus
	}
	resultInsert, errInsert := database.AddProduction(token, newProduction, idStock)
	if errInsert != nil {
		return c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Code:    constant.BadRequest,
			Message: errInsert.Error()})
	}

	return c.JSON(http.StatusOK, models.SuccessReponse{
		Code:    constant.Success,
		Message: "Add stock success",
		Data:    resultInsert})
}

func GetProductionByTokenHandles(c echo.Context) error {
	token := GetTokenHeader(c)
	stockID := c.Param("stock")
	stockIDParse, err := strconv.Atoi(stockID)
	if err != nil {
		stockIDParse = -1
	}
	dataProduction, errProduction := database.GetProduction(token, stockIDParse)
	if errProduction != nil {
		return c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Code:    constant.BadRequest,
			Message: errProduction.Error()})
	}
	return c.JSON(http.StatusOK, models.SuccessReponse{
		Code:    constant.Success,
		Message: "Retrieve production success",
		Data:    dataProduction})
}

func UpdateProduction(c echo.Context) error {
	var productUpdateBody models.Production
	err := c.Bind(&productUpdateBody)
	token := GetTokenHeader(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Code:    constant.BadRequest,
			Message: "Bad parameters"})
	}

	if c.Param("stock") == "" || c.Param("product") == "" {
		return c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Code:    constant.BadRequest,
			Message: "Bad parameters"})
	}

	convertStockID, errConvert := strconv.Atoi(c.Param("stock"))
	if errConvert != nil {
		return c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Code:    constant.BadRequest,
			Message: "Bad parameters"})
	}

	dataResultUpdate, erruUpdateProduction := database.UpdateProduct(token, convertStockID, c.Param("product"), productUpdateBody)
	if erruUpdateProduction != nil {
		return c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Code:    constant.BadRequest,
			Message: erruUpdateProduction.Error()})
	}

	return c.JSON(http.StatusOK, models.SuccessReponse{
		Code:    constant.Success,
		Message: "Update success",
		Data: dataResultUpdate})
}
