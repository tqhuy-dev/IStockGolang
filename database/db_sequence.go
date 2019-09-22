package database

import (
	"context"
	"errors"
	"github.com/tranhuy-dev/IStockGolang/models"
	"go.mongodb.org/mongo-driver/bson"
)

func GetSequenceStock(typeSequence string) (int, error) {
	sequenceStock := Client.Database("IStock").Collection("sequence")
	var sequence models.SequenceID
	filter := bson.D{
		{"sequence", typeSequence}}

	err := sequenceStock.FindOne(context.TODO(), filter).Decode(&sequence)
	if err != nil {
		return 0, errors.New("Sequence not found")
	}

	updateSequence , err := UpdateSequence("stock")
	if !updateSequence {
		return 0 , errors.New(err.Error())
	}
	return sequence.Count, nil
}

func UpdateSequence(typeSequence string) (bool, error) {
	sequenceCollection := Client.Database("IStock").Collection("sequence")
	var sequence models.SequenceID
	filter := bson.D{
		{"sequence", typeSequence}}
	updateBody := bson.D{
		{"$inc", bson.D{
			{"count", 1},
		}},
	}
	err := sequenceCollection.FindOneAndUpdate(context.TODO(), filter , updateBody).Decode(&sequence)
	if err != nil {
		return false , errors.New("Update sequence fail")
	}

	return true , nil
}
