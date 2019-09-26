package database

import (
	"github.com/tranhuy-dev/IStockGolang/models"
	"github.com/tranhuy-dev/IStockGolang/core/mathematic"
	"github.com/tranhuy-dev/IStockGolang/core/constant"
	"context"
	"errors"
	"time"
	"strconv"
)

func AddProduction(token string , newProduction models.Production , stockID int) (interface{} , error) {
	dataSession , errSession := CheckToken(token)
	if errSession != nil {
		return nil  , errSession
	}
	_ , errCheckStock := CheckUserStock(dataSession.Customer , stockID)
	if errCheckStock != nil {
		return nil , errCheckStock
	}
	newProduction.Customer  = dataSession.Customer
	newProduction.Stock = stockID
	newProduction.ID = mathematic.GetHash(newProduction.Name + strconv.Itoa(int(time.Now().Unix())))
	newProduction.CreateAt = strconv.Itoa(int(time.Now().Unix()))
	if newProduction.Status == "" {
		newProduction.Status = constant.STATUS_PROD_OPEN
	}
	
	productionCollection := Client.Database(DatabaseName).Collection("production")
	dataProduction , err := productionCollection.InsertOne(context.TODO() , newProduction)
	if err != nil {
		return nil , errors.New("Insert production fail")
	}
	return dataProduction , nil
}