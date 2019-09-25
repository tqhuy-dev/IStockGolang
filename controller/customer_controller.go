package controller

import (
	"github.com/labstack/echo"
	"github.com/tranhuy-dev/IStockGolang/handles"
)
// API Customer
func CustomerController(e *echo.Echo) {
	publicRoute := e.Group("/v1/customer")
	publicRoute.GET("/", handles.GetCustomer)
	publicRoute.GET("/filter", handles.FindUserByFilter)
	publicRoute.GET("/:email", handles.GetCustomerByEmail)
	publicRoute.POST("/password" , handles.ChangePassword)
	publicRoute.POST("/login" , handles.LoginAccount)
	publicRoute.POST("/", handles.CreateCustomer)
	publicRoute.PUT("/", handles.UpdateCustomer)
	publicRoute.DELETE("/" , handles.DeleteCustomer)
}