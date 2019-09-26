package models

type Production struct {
	Name string `json:"name" bson:"name"`
	Price int `json:"price" bson:"price"`
	Description string `json:"description" bson:"description"`
	Status string `json:"status" bson:"status"`
	ID string `json:"id" bson:"id"`
	CreateAt string `json:"create_at" bson:"create_at"`
	Stock int `json:"stock" bson:"stock"`
	Customer string `json:"customer" bson:"customer"`
}