package models

type Stock struct {
	Customer string `json:"customer" bson:"customer"`
	Name string `json:"name" bson:"name"`
	Description string `json:"description" bson:"description"`
	Status string `json:"status" bson:"status"`
	Price string `json:"price" bson:"price"`
	ID int `json:"id" bson:"id"`
}