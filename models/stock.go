package models

type Stock struct {
	Customer string `json:"customer" bson:"customer"`
	Name string `json:"name" bson:"name"`
	Description string `json:"description" bson:"description"`
	Status string `json:"status" bson:"status"`
	Price Price `json:"price" bson:"price"`
	ID int `json:"id" bson:"id"`
}

type Price struct {
	Available int `json:"available" bson:"available"`
	Sold int `json:"sold" bson:"sold"`
}