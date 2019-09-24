package models

type Session struct {
	Customer string `json:"customer" bson:"customer"`
	Token string `json:"token" bson:"token"`
	LoginTime string `json:"login_time" bson:"login_time"`
	ExpiredTime string `json:"expired_time" bson:"expired_time"`
}