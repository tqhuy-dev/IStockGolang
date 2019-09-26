package models

type Production struct {
	Name string `json:"name" bson:"name"`
	Price int `json:"price" bson:"price"`
	Description string `json:"description" bson:"description"`
	ID string `json:"id" bson:"id"`
	Status string `json:"status" bson:"status"`
	CreateAt string `json:"create_at" bson:"create_at"`
	Stock int `json:"stock" bson:"stock"`
	Customer string `json:"customer" bson:"customer"`
}