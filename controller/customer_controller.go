package controller

import (
	"github.com/labstack/echo"
	"github.com/tranhuy-dev/IStockGolang/handles"
)
// API Customer
func CustomerController(e *echo.Echo) {
	publicRoute := e.Group("/v1/customer")
	publicRoute.GET("/", handles.GetCustomer)
	publicRoute.POST("/", handles.CreateCustomer)
	publicRoute.PUT("/:id", handles.UpdateCustomer)
	publicRoute.DELETE("/:id" , handles.DeleteCustomer)
}