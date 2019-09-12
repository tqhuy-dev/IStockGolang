package models

type CustomerReq struct {
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Phone string `json:"phone"`
	Age int `json:"age"`
	Address string `json:"address"`
}