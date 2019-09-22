package controller

import (
	"github.com/labstack/echo"
	"github.com/tranhuy-dev/IStockGolang/handles"
)

// API Customer
func StockController(e *echo.Echo) {
	publicRoute := e.Group("/v1/stock")
	publicRoute.POST("/" , handles.CreateStockHandles)
	publicRoute.GET("/filter" , handles.RetriveStockByEmail)
}