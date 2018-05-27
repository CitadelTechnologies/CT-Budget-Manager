package transaction

import (
	"time"
	"gopkg.in/mgo.v2/bson"
)

const(
	TypeExpense = "expense"
	TypeIncome = "income"
)

type(
	Transaction struct {
		Id bson.ObjectId `json:"id" bson:"_id"`
		Wording string `json:"wording"`
		Description string `json:"description"`
		Type string `json:"type"`
		Amount float64 `json:"amount"`
		CreatedAt time.Time `json:"created_at"`
		ProcessedAt time.Time `json:"processed_at"`
	}
	Transactions []Transaction
)
