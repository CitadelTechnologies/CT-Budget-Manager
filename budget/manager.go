package budget

import(
	"ct-budget-manager/exception"
	"ct-budget-manager/server"
	"time"
	"gopkg.in/mgo.v2/bson"
)

func GetBudgets() Budgets {
  	budgets := make(Budgets, 0)
  	if err := server.App.Database.C("budget").Find(nil).All(&budgets); err != nil {
  		panic(exception.New(500, "Budgets retrieval query failed"))
    }
    return budgets
}

func GetBudget(id string) *Budget {
  	var budget Budget

  	if bson.IsObjectIdHex(id) && server.App.Database.C("budget").FindId(bson.ObjectIdHex(id)).One(&budget) == nil {
      	return &budget
    }
	panic(exception.New(404, "Budget not found"))
}

func CreateBudget(name string, description string) *Budget {
	budget := &Budget{
		Id: bson.NewObjectId(),
		Name: name,
		Description: description,
		Sectors: make(Sectors, 0),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	if err := server.App.Database.C("budget").Insert(budget); err != nil {
		panic(exception.New(500, "Budget creation failed"))
	}
	return budget
}
