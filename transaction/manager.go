package transaction

import(
	"ct-budget-manager/exception"
	"ct-budget-manager/server"
	"time"
	"gopkg.in/mgo.v2/bson"
)

func CreateTransaction(budgetSlug string, sectorSlug string, wording string, description string, tType string, amount float64, processedAt time.Time) *Transaction {
	if tType != TypeExpense && tType != TypeIncome {
		panic(exception.New(400, "Invalid transaction type", nil))
	}
	transaction := &Transaction{
		Id: bson.NewObjectId(),
		Wording: wording,
		Description: description,
		Type: tType,
		Amount: amount,
		CreatedAt: time.Now(),
		ProcessedAt: processedAt,
	}
	change := bson.M{
		"$push": bson.M{"sectors.$.transactions": transaction},
	}
	if err := server.App.Database.C("budget").Update(bson.M{"slug": budgetSlug, "sectors.slug": sectorSlug}, change); err != nil {
		panic(exception.New(404, "Budget or sector not found", err))
	}
	return transaction
}

func GetTransaction(budgetSlug string, sectorSlug string, id string) *Transaction {
	var transaction Transaction

    pipe := server.App.Database.C("budget").Pipe(
        []bson.M{
			bson.M{
				"$match": bson.M{"slug": budgetSlug},
			},
			bson.M{
				"$unwind": "$sectors",
			},
			bson.M{
				"$match": bson.M{"sectors.slug": sectorSlug},
			},
			bson.M{
				"$unwind": "$sectors.transactions",
			},
			bson.M{
				"$match": bson.M{"sectors.transactions._id": bson.ObjectIdHex(id)},
			},
			bson.M{
				"$project": bson.M{
					"_id": "$sectors.transactions._id",
					"wording": "$sectors.transactions.wording",
					"description": "$sectors.transactions.description",
					"amount": "$sectors.transactions.amount",
					"type": "$sectors.transactions.type",
					"createdat": "$sectors.transactions.createdat",
				},
			},
        },
    )
    if err := pipe.One(&transaction); err != nil {
		panic(exception.New(404, "Not found", err))
	}
	return &transaction
}

func GetTransactions(budgetSlug string, sectorSlug string) Transactions {
  	transactions := make(Transactions, 0)

    pipe := server.App.Database.C("budget").Pipe(
        []bson.M{
			bson.M{
				"$match": bson.M{"slug": budgetSlug},
			},
            bson.M{
                "$unwind": "$sectors",
            },
			bson.M{
				"$match": bson.M{"sectors.slug": sectorSlug},
			},
            bson.M{
                "$unwind": "$sectors.transactions",
            },
            bson.M{
                "$project": bson.M{
					"_id": "$sectors.transactions._id",
                    "wording": "$sectors.transactions.wording",
                    "description": "$sectors.transactions.description",
                    "amount": "$sectors.transactions.amount",
                    "type": "$sectors.transactions.type",
					"createdat": "$sectors.transactions.createdat",
                },
            },
        },
    )
    if err := pipe.All(&transactions); err != nil {
		panic(exception.New(404, "Budget or sector not found", err))
	}
    return transactions
}
