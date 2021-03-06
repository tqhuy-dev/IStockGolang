package controller

import (
	"github.com/labstack/echo"
	"github.com/tranhuy-dev/IStockGolang/handles"
)

func ProductionController (e *echo.Echo) {
	publicRoute := e.Group("/v1/production")
	publicRoute.POST("/:stock" , handles.AddProductionHandles)
	publicRoute.PUT("/:product/stock/:stock" , handles.UpdateProduction)
	publicRoute.GET("/:stock" , handles.GetProductionByTokenHandles)
	publicRoute.GET("/" , handles.GetProductionByTokenHandles)
}