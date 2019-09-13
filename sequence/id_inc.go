package sequence

import (
	"context"
	"time"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var Client *mongo.Client

type SequenceID struct {
	sequence string `bson:"sequence"`
	count int `bson:"count"`
}
func init() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		panic(err)
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		panic(err)
	}
	Client = client
}

func IncID() int {
	var sequenceID SequenceID
	customerCollection := Client.Database("IStock").Collection("sequence_id")
	filter := bson.D{{"sequence","id"}}
	updateBody := bson.D{
		{"$inc", bson.D{
			{"count",1},
		}},
	}
	customerCollection.UpdateOne(context.TODO() , filter , updateBody)
	err := customerCollection.FindOne(context.TODO() , bson.D{
		{"sequence" , "id"},
	}).Decode(&sequenceID)
	if err != nil {
		
	}
	fmt.Println(sequenceID)
	return sequenceID.count
}