package budget

import (
	"ct-budget-manager/transaction"
	"time"
	"gopkg.in/mgo.v2/bson"
)

type(
	Budget struct {
		Id bson.ObjectId `json:"id" bson:"_id"`
		Name string `json:"name"`
		Slug string `json:"slug"`
		Description string `json:"description"`
		Sectors Sectors `json:"sectors"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}
	Budgets []Budget

	Sector struct {
		Name string `json: "name"`
		Slug string `json: "slug"`
    	Transactions transaction.Transactions `json:"transactions"`
	}
	Sectors []Sector
)
