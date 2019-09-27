package database
import (
	"go.mongodb.org/mongo-driver/bson"
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/mongo/options"
	"github.com/tranhuy-dev/IStockGolang/models"
	"github.com/tranhuy-dev/IStockGolang/models/response_models"
	"github.com/tranhuy-dev/IStockGolang/core/constant"
	// "fmt"
)

func CreateStock(token string , stock models.Stock) (interface{} , error){
	stockCollection := Client.Database(DatabaseName).Collection("stock")
	dataSession, errSession := CheckToken(token)
	if errSession != nil {
		return nil , errSession
	}

	idStock , err := GetSequenceStock("stock")
	if err != nil {
		return nil , errors.New(err.Error())
	}
	newStock := models.Stock{
		Customer: dataSession.Customer,
		Description: stock.Description,
		Name: stock.Name,
		Status: stock.Status,
		ID: idStock}

	insertResult , insertError := stockCollection.InsertOne(context.TODO() , newStock)
	if insertError != nil {
		return nil,errors.New("Insert fail") 
	}
	return insertResult, nil
}

func RetrieveStockUser(token , email string) (interface{}, error) {
	stockCollection := Client.Database(DatabaseName).Collection("stock")
	var listStockRes []response_models.StockResponse
	findOption := options.Find()
	findOption.SetLimit(100)
	filter := bson.D{
		{"customer" , email}}
	
	cur, err := stockCollection.Find(context.TODO() , filter , findOption)
	if err != nil {
		return nil , errors.New("User not found")
	}

	for cur.Next(context.TODO()) {
		var elementStock models.Stock
		err := cur.Decode(&elementStock)
		dataProduction , errDataProduct := GetProduction(token , elementStock.ID)
		if errDataProduct != nil {
			return nil , errDataProduct
		}
		elementStockRes := response_models.StockResponse{
			Name: elementStock.Name,
			Price: elementStock.Price,
			Description: elementStock.Description,
			ID: elementStock.ID,
			Status: elementStock.Status,
			Product: dataProduction,
		}
		if err != nil {
			return nil , errors.New(constant.MessageUnexpectedError)
		}

		listStockRes = append(listStockRes  , elementStockRes)

	}

	return listStockRes , nil
}

func RetriveStockByToken(token string) (interface{} , error) {
	dataSession , err := CheckToken(token)
	if err != nil {
		return nil , err
	}
	dataStock , err := RetrieveStockUser(token , dataSession.Customer)
	if err != nil {
		return nil , err
	}

	return dataStock , nil
}

func UpdateStock(token string, idStock int, stockBody models.Stock) (interface{} , error) {
	stockCollection := Client.Database(DatabaseName).Collection("stock")
	dataSession , errSession := CheckToken(token)
	if errSession != nil {
		return nil , errSession
	}
	filter := bson.D{
		{"id_stock" , idStock},
		{"customer",dataSession.Customer},
}

	updateBody := bson.M{}
	if stockBody.Description != "" {
		updateBody["description"] = stockBody.Description
	}
	if stockBody.Name != "" {
		updateBody["name"] = stockBody.Name
	}
	if stockBody.Status != "" {
		updateBody["status"] = stockBody.Status
	}
	update := bson.D{
		{"$set", updateBody}}

	var updateStock models.Stock
	errUpdateStock := stockCollection.FindOneAndUpdate(context.TODO() , filter , update).Decode(&updateStock)
	if errUpdateStock != nil {
		return nil , errors.New("Update stock fail")
	}

	return &updateStock , nil
}

func CheckUserStock(customer string , idStock int) (interface{} , error) {
	stockCollection := Client.Database(DatabaseName).Collection("stock")
	filter := bson.D{
		{"customer" , customer},
		{"id_stock" , idStock},
	}

	var Stock models.Stock
	err := stockCollection.FindOne(context.TODO() , filter).Decode(&Stock)
	if err != nil {
		return nil , errors.New("Stock not found")
	}

	return Stock , nil
}

func UpdateTotalPriceStock(product models.Production, stockID int) (interface{}  , error) {
	stockCollection := Client.Database(DatabaseName).Collection("stock")
	filter := bson.D{
		{"id_stock" , stockID},
	}

	var Stock models.Stock

	err := stockCollection.FindOne(context.TODO() , filter).Decode(&Stock)
	if err != nil {
		return nil , errors.New("Get stock fail")
	}

	if product.Status != constant.STATUS_PROD_SOLD {
		Stock.Price.Available += product.Price
	} else {
		Stock.Price.Available -= product.Price
		Stock.Price.Sold += product.Price
	}

	updateInfor := bson.D{
		{"price" , Stock.Price},
	}

	updateBody := bson.D{
		{"$set" , updateInfor},
	}

	var newStock models.Stock
	errUpdate := stockCollection.FindOneAndUpdate(context.TODO() , filter , updateBody).Decode(&newStock)
	if errUpdate != nil {
		return nil , errors.New("Update stock fail")
	}

	return newStock , nil
}