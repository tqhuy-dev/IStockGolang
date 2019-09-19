package handles

import (
	"net/http"
	"github.com/labstack/echo"
	models "github.com/tranhuy-dev/IStockGolang/models"
	"github.com/tranhuy-dev/IStockGolang/core/constant"
	"github.com/tranhuy-dev/IStockGolang/database"
)

func CreateStockHandles(c echo.Context) error{
	var stockReq models.Stock
	err := c.Bind(&stockReq)
	if err != nil {
		return c.JSON(http.StatusBadRequest , models.ErrorResponse{
			Code: constant.BadRequest,
			Message: "Bad parameters"})
	}
	dataStock , err := database.CreateStock(stockReq)
	if err != nil {
		return c.JSON(http.StatusForbidden , models.ErrorResponse{
			Code: constant.ExpectedError,
			Message: err.Error()})
	}

	return c.JSON(http.StatusOK , models.SuccessReponse{
		Code: constant.Success,
		Message: "Add stock success",
		Data: dataStock})
}