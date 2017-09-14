package manager

import(
	"ct-budget-manager/model"
	"time"
	"gopkg.in/mgo.v2/bson"
)

/*
* @return model.Budget
*/
func GetBudgets() model.Budgets {
  	budgets := make(model.Budgets, 0)
  	if err := MongoDBConnection.DB(MongoDBName).C("budget").Find(nil).All(&budgets); err != nil {
      panic(err)
    }
    return budgets
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
