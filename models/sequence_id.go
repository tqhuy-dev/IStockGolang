package models

type SequenceID struct {
	Sequence string `bson:"sequence"`
	Count int `bson:"count"`
}