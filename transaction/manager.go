package transaction

import(
	"ct-budget-manager/exception"
	"ct-budget-manager/server"
	"time"
	"gopkg.in/mgo.v2/bson"
)

func CreateTransaction(budgetId string, wording interface{}, description interface{}, tType interface{}, amount interface{}) *Transaction {
	transaction := Transaction{
		Id: bson.NewObjectId(),
		Wording: wording.(string),
		Description: description.(string),
		Type: tType.(bool),
		Amount: amount.(float64),
		CreatedAt: time.Now(),
	}
	//if !AddTransactionToBudget(budgetId, transaction) {
	//	panic(exception.New(500, "Transaction could not be created"))
	//}
	return &transaction
}

func GetTransaction(id string) *Transaction {
	var raw interface{}

	if err := server.App.Database.
	C("budget").
	Find(bson.M{"transactions._id": bson.ObjectIdHex(id)}).
	Select(bson.M{"_id": 0, "transactions.$": 1}).
	One(&raw); err != nil {
		if err.Error() == "not found" {
			panic(exception.New(404, "Transaction not found"))
		}
		panic(exception.New(500, "Transaction could not be retrieved"))
	}
	var data map[string]Transactions
	bytes, _ := bson.Marshal(raw)
	bson.Unmarshal(bytes, &data)
	return &data["transactions"][0]
}

func GetTransactions() Transactions {
	transactions := make(Transactions, 0)
	if err := server.App.Database.C("budget").Find(nil).All(&transactions); err != nil {
		panic(exception.New(500, "Transactions could not be retrieved"))
	}
	return transactions
}
