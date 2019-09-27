package response_models

import (
	"github.com/tranhuy-dev/IStockGolang/models"
)

type CustomerResponse struct {
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Address string `json:"address"`
	Age int `json:"age"`
	Email string `json:"email"`
	Stock interface{} `json:"stock"`
}

type StockResponse struct {
	Name string `json:"name"`
	Description string `json:"description"`
	Price models.Price `json:"price"`
	Product interface{} `json:"product"`
	ID int `json:"id_stock"`
	Status string `json:"status"`
}