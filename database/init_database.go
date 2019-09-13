package database

import (
	"context"
	"log"
	"time"
	"fmt"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"crypto/sha256"
	"github.com/tranhuy-dev/IStockGolang/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const (
	dbName = "IStock"
)

var Client *mongo.Client

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

func InsertCustomer(req models.CustomerReq) interface{} {
	newCustomer := models.Customer{
		FirstName: req.FirstName,
		LastName: req.LastName,
		Phone: req.Phone,
		Address: req.Address,
		Age: req.Age,
		Status:1,
		Email:req.Email,
		Password:req.Password}
	customerCollection := Client.Database("IStock").Collection("customer")
	_, errorQueryInsert := customerCollection.InsertOne(context.TODO(), newCustomer)
	if errorQueryInsert != nil {
		log.Fatal(errorQueryInsert)
	}
	hashToken := sha256.Sum256([]byte(newCustomer.Email))
	responseBody := map[string]interface{}{}
	responseBody["token"] = hashToken[:]
	return responseBody
}

func RetrieveAllCustomer() interface{} {
	var customer []*models.Customer
	customerCollection := Client.Database("IStock").Collection("customer")
	findOptions := options.Find()
	findOptions.SetLimit(100)
	cur, err := customerCollection.Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		log.Fatal(err)
	}
	for cur.Next(context.TODO()) {

		// create a value into which the single document can be decoded
		var elem models.Customer
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		customer = append(customer, &elem)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	cur.Close(context.TODO())
	responseBody := map[string]interface{}{}
	responseBody["customer"] = customer
	responseBody["size"] = len(customer)
	return responseBody
}

func UpdateCustomer(req models.CustomerReq , email string) (*models.Customer , error){
	customerCollection := Client.Database("IStock").Collection("customer")

	filter := bson.D{{"email",email}}

	updateBody := bson.D{
		{"$set" , bson.D{
			{"first_name" , req.FirstName},
			{"last_name",req.LastName},
			{"age",req.Age},
			{"phone",req.Phone},
			{"address",req.Address},
		}},
	}
	var customer models.Customer
	err := customerCollection.FindOneAndUpdate(context.TODO() , filter , updateBody).Decode(&customer)
	if err != nil {
		return nil , errors.New("User not found")
	}
	return &customer, nil
}

func DeleteCustomer(email string) (*models.Customer , error) {
	customerCollection := Client.Database("IStock").Collection("customer")
	filter := bson.D{{"email" , email}}
	updateBody := bson.D{
		{"$set", bson.D{
			{"status",0},
		}},
	}
	var customer models.Customer
	err := customerCollection.FindOneAndUpdate(context.TODO() , filter , updateBody).Decode(&customer)
	if err != nil {
		return nil,errors.New("User not found")
	}
	return &customer,nil
}

func FindUserByEmail(email string) (*models.Customer, error) {
	var customer models.Customer
	customerCollection := Client.Database("IStock").Collection("customer")
	err := customerCollection.FindOne(context.TODO() , bson.D{
		{"email" , email},
	}).Decode(&customer)
	if err != nil {
		return nil, errors.New("ko the tim user")
	}
	return &customer,nil
}
func IncID() interface{} {
	var sequenceID models.SequenceID
	seCollection := Client.Database("IStock").Collection("sequence")
	filter := bson.D{{"sequence_type","sequence_id"}}
	updateBody := bson.D{
		{"$inc", bson.D{
			{"count",1},
		}},
	}
	seCollection.UpdateOne(context.TODO() , filter , updateBody)

	err := seCollection.FindOne(context.TODO() , filter).Decode(&sequenceID)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("mysequence %+v",sequenceID)
	return sequenceID.Count
}

func Syncnorize() {
	customerCollection := Client.Database("IStock").Collection("customer")
	updateBody := bson.D{
		{"$set",bson.D{
			{"password","123456"},
		}},
	}
	customerCollection.UpdateMany(context.TODO(),bson.D{{}} , updateBody)
}

func ChangePassword(passwordReq models.ChangePasswordReq) (*models.Customer , error) {
	customerCollection := Client.Database("IStock").Collection("customer")
	filter := bson.D{
		{"password",passwordReq.OldPassword},
		{"email",passwordReq.Email},
	}

	updateBody := bson.D{
		{"$set", bson.D{
			{"password",passwordReq.NewPassword},
		}},
	}

	var customer models.Customer
	err := customerCollection.FindOneAndUpdate(context.TODO() , filter , updateBody).Decode(&customer)
	if err != nil {
		return nil,errors.New("Change password fail")
	}
	return &customer,nil
}