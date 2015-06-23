package model

import (
	"time"
	"gopkg.in/mgo.v2/bson"
)

type(
	Transaction struct {
		Id bson.ObjectId `json:"id" bson:"_id"`
		Wording string `json:"wording"`
		Description string `json:"description"`
		Type bool `json:"type"`
		Sector Sector `json:"sector"`
		Amount int `json:"amount"`
		Date time.Time `json:"date"`
	}
	Transactions []Transaction
)
