package controller

import (
	"github.com/labstack/echo"
	"github.com/tranhuy-dev/IStockGolang/handles"
)

func ProductionController (e *echo.Echo) {
	publicRoute := e.Group("/v1/production")
	publicRoute.POST("/" , handles.AddProductionHandles)
	publicRoute.GET("/" , handles.GetProductionByTokenHandles)
}