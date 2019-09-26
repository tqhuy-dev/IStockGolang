package handles

import (
	"github.com/labstack/echo"
	"net/http"
	"github.com/tranhuy-dev/IStockGolang/models"
	"github.com/tranhuy-dev/IStockGolang/core/constant"
)

func AddProductionHandles(c echo.Context) error {
	return c.JSON(http.StatusOK , models.SuccessReponse{
		Code: constant.Success,
		Message: "Add stock success",
		Data: ""})
}

func GetProductionByTokenHandles(c echo.Context) error {
	return c.JSON(http.StatusOK , models.SuccessReponse{
		Code: constant.Success,
		Message: "Retrieve production success",
		Data: ""})
}