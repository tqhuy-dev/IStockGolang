package database

import (
	"context"
	"errors"
	// "fmt"
	"strconv"
	"time"

	"github.com/tranhuy-dev/IStockGolang/models"
	"go.mongodb.org/mongo-driver/bson"
)

func CheckToken(token string) {

}

func UpdateSession(token, customer string) (interface{}, error) {
	sessionCollection := Client.Database(DatabaseName).Collection("session")
	timestampLogin := strconv.Itoa(int(time.Now().Unix()))
	timestampExpired := strconv.Itoa(int(time.Now().AddDate(0, 1, 0).Unix()))
	filter := bson.D{
		{"customer", customer},
	}

	updateBody := bson.D{
		{"token", token},
		{"login_time", timestampLogin},
		{"expired_time", timestampExpired},
	}

	updateSession := bson.D{
		{"$set" , updateBody}}

	var newSession models.Session
	err := sessionCollection.FindOneAndUpdate(context.TODO(), filter, updateSession).Decode(&newSession)
	if err != nil {
		return nil, errors.New("Update session fail")
	}

	return newSession, nil
}

func CreateSessionToken(token, customer string) (interface{}, error) {
	sessionCollection := Client.Database(DatabaseName).Collection("session")
	var newSession models.Session
	err := sessionCollection.FindOne(context.TODO(), bson.D{
		{"customer", customer}}).Decode(&newSession)

	if err != nil { // user have not login ever
		timestampToday := strconv.Itoa(int(time.Now().Unix()))
		timestampExpiredTime := strconv.Itoa(int(time.Now().AddDate(0, 1, 0).Unix()))
		newSession := models.Session{
			Customer:    customer,
			Token:       token,
			LoginTime:   timestampToday,
			ExpiredTime: timestampExpiredTime}
		dataInsert, insertErr := sessionCollection.InsertOne(context.TODO(), newSession)
		if insertErr != nil {
			return nil, errors.New("Session fail")
		}
		return dataInsert, nil
	}

	dataSession, errUpdateSession := UpdateSession(token, customer)

	if errUpdateSession != nil {
		return nil, errUpdateSession
	}

	return dataSession, nil
}
