package handles

import (
	"github.com/labstack/echo"
	"net/http"
	"github.com/tranhuy-dev/IStockGolang/models"
	"github.com/tranhuy-dev/IStockGolang/database"
	"github.com/tranhuy-dev/IStockGolang/core/constant"
	"strconv"
)

func AddProductionHandles(c echo.Context) error {
	token := GetTokenHeader(c)
	idStock, _ := strconv.Atoi(c.Param("stock"))
	var newProduction models.Production
	err := c.Bind(&newProduction) 
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Code:    constant.BadRequest,
			Message: "Bad parameters"})
	}

	resultInsert , errInsert := database.AddProduction(token , newProduction , idStock)
	if errInsert != nil {
		return c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Code:    constant.BadRequest,
			Message: errInsert.Error()})
	}

	return c.JSON(http.StatusOK , models.SuccessReponse{
		Code: constant.Success,
		Message: "Add stock success",
		Data: resultInsert})
}

func GetProductionByTokenHandles(c echo.Context) error {
	return c.JSON(http.StatusOK , models.SuccessReponse{
		Code: constant.Success,
		Message: "Retrieve production success",
		Data: ""})
}