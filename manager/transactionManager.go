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
	floatType := tType.(float64)
	floatAmount := amount.(float64)
	sectorName := sector.(map[string]interface{})["name"]
	
	transaction := model.Transaction{
		Id: bson.NewObjectId(),
		Wording: wording.(string),
		Description: description.(string),
		Type: (int(floatType) > 0),
		Sector: model.Sector{Name: sectorName.(string)},
		Amount: int(floatAmount),
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
	var transaction model.Transaction

	if !bson.IsObjectIdHex(id) {
		return nil
	}
	if err := MongoDBConnection.DB(MongoDBName).C("budget").FindId(bson.ObjectIdHex(id)).One(&transaction); err != nil {
		panic(err)
	}
	return &transaction
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
