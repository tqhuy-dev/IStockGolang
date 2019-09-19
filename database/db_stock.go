package database
import (
	"context"
	"errors"
	"github.com/tranhuy-dev/IStockGolang/models"
	"github.com/tranhuy-dev/IStockGolang/core/constant"
)

func CreateStock(stock models.Stock) (interface{} , error){
	stockCollection := Client.Database("IStock").Collection("stock")

	newStock := models.Stock{
		Customer: stock.Customer,
		Description: stock.Description,
		Name: stock.Name,
		Price: stock.Price,
		Status: stock.Status}

	_ , searchUserError := FindUserByEmail(stock.Customer)
	if searchUserError  != nil {
		return nil , errors.New(constant.MessageUserNotFound)
	}
	insertResult , insertError := stockCollection.InsertOne(context.TODO() , newStock)
	if insertError != nil {
		return nil,errors.New("Insert fail") 
	}
	return insertResult, nil
}