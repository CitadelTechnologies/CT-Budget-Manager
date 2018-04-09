package transaction

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
		Amount float64 `json:"amount"`
		CreatedAt time.Time `json:"created_at"`
	}
	Transactions []Transaction
)