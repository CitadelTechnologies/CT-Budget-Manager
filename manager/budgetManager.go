package manager

import(
	"ct-budget-manager/model"
	"time"
	"gopkg.in/mgo.v2/bson"
)

/*
* @return model.Budgets
*/
func GetBudgets() model.Budgets {
  	budgets := make(model.Budgets, 0)
  	if err := MongoDBConnection.DB(MongoDBName).C("budget").Find(nil).All(&budgets); err != nil {
      panic(err)
    }
    return budgets
}

/*
* @return model.Budget
*/
func GetBudget(id string) *model.Budget {
  	var budget model.Budget

  	if !bson.IsObjectIdHex(id) {
  		return nil
  	}
  	if err := MongoDBConnection.DB(MongoDBName).C("budget").FindId(bson.ObjectIdHex(id)).One(&budget); err != nil {
      panic(err)
    }
    return &budget
}

/*
* @param string name
* @param string description
* @return model.Budget
*/
func CreateBudget(name string, description string) model.Budget {
	var budget model.Budget

	budget.Id = bson.NewObjectId()
  budget.Name = name
  budget.Description = description
  budget.Transactions = make(model.Transactions, 0)
	budget.CreatedAt = time.Now()
	budget.UpdatedAt = time.Now()

  if err := MongoDBConnection.DB(MongoDBName).C("budget").Insert(budget); err != nil {
    panic(err)
  }
  return budget
}

/*
* @param string budgetId
* @param model.Transaction transaction
* @return bool
*/
func AddTransactionToBudget(budgetId string, transaction model.Transaction) bool {
  budget := GetBudget(budgetId)
  if budget == nil {
    return false
  }
  budget.Transactions = append(budget.Transactions, transaction)
  if err := MongoDBConnection.DB(MongoDBName).C("budget").UpdateId(budget.Id, budget); err != nil {
    panic(err)
  }
  return true
}
