package budget

import(
	"ct-budget-manager/exception"
	"ct-budget-manager/server"
	"ct-budget-manager/transaction"
	"github.com/gosimple/slug"
	"time"
	"gopkg.in/mgo.v2/bson"
)

func GetBudgets() Budgets {
  	budgets := make(Budgets, 0)
  	if err := server.App.Database.C("budget").Find(nil).All(&budgets); err != nil {
  		panic(exception.New(500, "Budgets retrieval query failed", err))
    }
    return budgets
}

func GetBudget(slug string) *Budget {
  	var budget Budget

  	if err := server.App.Database.C("budget").Find(bson.M{"slug": slug}).One(&budget); err != nil {
      	panic(exception.New(404, "Budget not found", err))
    }
	return &budget
}

func CreateBudget(name string, description string) *Budget {
	budget := &Budget{
		Id: bson.NewObjectId(),
		Name: name,
		Slug: slug.Make(name),
		Description: description,
		Sectors: make(Sectors, 0),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	if err := server.App.Database.C("budget").Insert(budget); err != nil {
		panic(exception.New(500, "Budget creation failed", err))
	}
	return budget
}

func CreateSector(budgetSlug string, name string) *Sector {
	sector := &Sector{
		Name: name,
		Slug: slug.Make(name),
		Transactions: make(transaction.Transactions, 0),
	}
	change := bson.M{
		"$set": bson.M{"updatedat": time.Now()},
		"$push": bson.M{"sectors": sector},
	}
	if err := server.App.Database.C("budget").Update(bson.M{"slug": budgetSlug}, change); err != nil {
		panic(exception.New(404, "Budget not found", err))
	}
	return sector
}
