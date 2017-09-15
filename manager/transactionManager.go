package manager

import(
	"ct-budget-manager/model"
	"time"
	"gopkg.in/mgo.v2/bson"
)

/*
* @param string budgetId
* @param interface{} wording
* @param interface{} description
* @param interface{} tType
* @param interface{} sector
* @param interface{} amount
* @parem *model.Transaction
*/
func CreateTransaction(budgetId string, wording interface{}, description interface{}, tType interface{}, sector interface{}, amount interface{}) *model.Transaction {
	sectorName := sector.(map[string]interface{})["name"]

	transaction := model.Transaction{
		Id: bson.NewObjectId(),
		Wording: wording.(string),
		Description: description.(string),
		Type: tType.(bool),
		Sector: model.Sector{Name: sectorName.(string)},
		Amount: amount.(float64),
		CreatedAt: time.Now(),
	}
	if !AddTransactionToBudget(budgetId, transaction) {
		return nil
	}
	return &transaction
}

/*
* @param string id
* @return *model.Transaction
*/
func GetTransaction(id string) *model.Transaction {
	var raw interface{}

	if err := MongoDBConnection.
	DB(MongoDBName).
	C("budget").
	Find(bson.M{"transactions._id": bson.ObjectIdHex(id)}).
	Select(bson.M{"_id": 0, "transactions.$": 1}).
	One(&raw); err != nil {
		if err.Error() == "not found" {
			return nil
		}
		panic(err)
	}
	var data map[string]model.Transactions
	bytes, _ := bson.Marshal(raw)
	bson.Unmarshal(bytes, &data)
	return &data["transactions"][0]
}

/*
* @return model.Transactions
*/
func GetTransactions() model.Transactions {
	transactions := make(model.Transactions, 0)
	if err := MongoDBConnection.DB(MongoDBName).C("budget").Find(nil).All(&transactions); err != nil {
		panic(err)
	}
	return transactions
}
