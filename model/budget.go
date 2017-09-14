package model

import (
	"time"
	"gopkg.in/mgo.v2/bson"
)

type(
	Budget struct {
		Id bson.ObjectId `json:"id" bson:"_id"`
		Name string `json:"name"`
		Description string `json:"description"`
    Transactions Transactions `json:"transactions"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}
	Budgets []Budget
)
