package database
import (
	"go.mongodb.org/mongo-driver/bson"
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/mongo/options"
	"github.com/tranhuy-dev/IStockGolang/models"
	"github.com/tranhuy-dev/IStockGolang/core/constant"
	// "fmt"
)

func CreateStock(stock models.Stock) (interface{} , error){
	stockCollection := Client.Database(DatabaseName).Collection("stock")
	idStock , err := GetSequenceStock("stock")
	if err != nil {
		return nil , errors.New(err.Error())
	}
	newStock := models.Stock{
		Customer: stock.Customer,
		Description: stock.Description,
		Name: stock.Name,
		Status: stock.Status,
		ID: idStock}

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

func RetrieveStockUser(email string) ([]*models.Stock, error) {
	stockCollection := Client.Database(DatabaseName).Collection("stock")
	var listStock []*models.Stock
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

		if err != nil {
			return nil , errors.New(constant.MessageUnexpectedError)
		}

		listStock = append(listStock  , &elementStock)
	}

	return listStock , nil
}

func RetriveStockByToken(token string) ([]*models.Stock , error) {
	dataSession , err := CheckToken(token)
	if err != nil {
		return nil , err
	}
	dataStock , err := RetrieveStockUser(dataSession.Customer)
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