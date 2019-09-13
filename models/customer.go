package models


type Customer struct {
	FirstName string `bson:"first_name" json:"first_name"`
	LastName string `bson:"last_name" json:"last_name"`
	Age int `bson:"age" json:"age"`
	Phone string `bson:"phone" json:"phone"`
	Address string `bson:"address" json:"address"`
	Status int `bson:"status" json:"status"`
	Email string `bson:"email" json:"email"`
	Password string `bson:"password" json:"password"`
}