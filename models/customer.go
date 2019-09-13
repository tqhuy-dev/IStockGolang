package models

type Customer struct {
	FirstName string `bson:"first_name"`
	LastName string `bson:"last_name"`
	Age int `bson:"age"`
	Phone string `bson:"phone"`
	Address string `bson:"address"`
	Status int `bson:"status"`
}