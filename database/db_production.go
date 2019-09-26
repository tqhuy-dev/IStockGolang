package database

import (
	"context"
	"errors"
	"log"
	"strconv"
	"time"

	"github.com/tranhuy-dev/IStockGolang/core/constant"
	"github.com/tranhuy-dev/IStockGolang/core/mathematic"
	"github.com/tranhuy-dev/IStockGolang/models"
	"go.mongodb.org/mongo-driver/bson"
	// "fmt"
)

func AddProduction(token string, newProduction models.Production, stockID int) (interface{}, error) {
	dataSession, errSession := CheckToken(token)
	if errSession != nil {
		return nil, errSession
	}
	_, errCheckStock := CheckUserStock(dataSession.Customer, stockID)
	if errCheckStock != nil {
		return nil, errCheckStock
	}
	newProduction.Customer = dataSession.Customer
	newProduction.Stock = stockID
	newProduction.ID = mathematic.GetHash(newProduction.Name + strconv.Itoa(int(time.Now().Unix())))
	newProduction.CreateAt = strconv.Itoa(int(time.Now().Unix()))
	if newProduction.Status == "" {
		newProduction.Status = constant.STATUS_PROD_OPEN
	}

	productionCollection := Client.Database(DatabaseName).Collection("production")
	dataProduction, err := productionCollection.InsertOne(context.TODO(), newProduction)
	if err != nil {
		return nil, errors.New("Insert production fail")
	}
	return dataProduction, nil
}

func GetProduction(token string, stockID int) (interface{}, error) {
	productionCollection := Client.Database(DatabaseName).Collection("production")

	dataSession, errSession := CheckToken(token)
	if errSession != nil {
		return nil, errSession
	}
	filter := bson.M{}
	filter["customer"] = dataSession.Customer
	if stockID != -1 {
		_, errCheckStock := CheckUserStock(dataSession.Customer, stockID)
		if errCheckStock != nil {
			return nil, errCheckStock
		}
		filter["stock"] = stockID
	}
	// var production []*models.Production
	cur, errQueryProduction := productionCollection.Find(context.TODO(), filter)
	var production []*models.Production
	if errQueryProduction != nil {
		return nil, errors.New("Retrive production fail")
	}
	for cur.Next(context.TODO()) {

		// create a value into which the single document can be decoded
		var elem models.Production
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		production = append(production, &elem)
	}
	return production, nil
}

func UpdateProduct(token string, idStock int, idProduct string, product models.Production) (interface{}, error) {
	productionCollection := Client.Database(DatabaseName).Collection("production")

	dataSession, errSession := CheckToken(token)
	if errSession != nil {
		return nil, errSession
	}

	filter := bson.D{
		{"customer", dataSession.Customer},
		{"stock", idStock},
		{"id", idProduct},
	}

	updateInfor := bson.M{}

	if product.Name != "" {
		updateInfor["name"] = product.Name
	}

	if product.Description != "" {
		updateInfor["description"] = product.Description
	}

	if product.Price != 0 {
		updateInfor["price"] = product.Price
	}

	if product.Status != "" {
		updateInfor["status"] = product.Status
	}

	updateBody := bson.D{
		{"$set", updateInfor},
	}

	var newProduct models.Production
	err := productionCollection.FindOneAndUpdate(context.TODO(), filter, updateBody).Decode(&newProduct)

	if err != nil {
		return nil, errors.New("Update product fail")
	}

	return newProduct, nil
}
