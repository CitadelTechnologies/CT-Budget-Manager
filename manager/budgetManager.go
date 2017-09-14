package manager

import(
	"ct-budget-manager/model"
	"time"
	"gopkg.in/mgo.v2/bson"
)

/*
* POST request to create a new Budget object
*
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

  if err := MongoDBConnection.DB("test").C("budget").Insert(budget); err != nil {
    panic(err)
  }
  return budget
}
