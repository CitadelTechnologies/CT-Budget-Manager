package main

import (
	"net/http"

	"github.com/gorilla/mux"

	"ct-budget-manager/budget"
	"ct-budget-manager/transaction"
)

type(
	Route struct {
		Name        string
		Method      string
		Pattern     string
		HandlerFunc http.HandlerFunc
	}
	Routes []Route
)

func NewRouter() *mux.Router {
    router := mux.NewRouter().StrictSlash(true)
    for _, route := range routes {
		go func(router *mux.Router, route Route){
			router.
				Methods(route.Method).
				Path(route.Pattern).
				Name(route.Name).
				Handler(route.HandlerFunc)
		}(router, route)
    }
    return router
}

var routes = Routes{
	Route{
		"Get Budgets",
		"GET",
		"/budgets",
		budget.GetBudgetsAction,
	},
	Route{
		"Create Budget",
		"POST",
		"/budgets",
		budget.CreateBudgetAction,
	},
	Route{
		"Create Sector",
		"POST",
		"/budgets/{slug}/sectors",
		budget.CreateSectorAction,
	},
	Route{
		"Get Budget",
		"GET",
		"/budgets/{slug}",
		budget.GetBudgetAction,
	},
    Route{
        "Transactions",
        "GET",
        "/budgets/{budget}/sectors/{sector}/transactions",
        transaction.GetTransactionsAction,
    },
    Route{
        "Create Transaction",
        "POST",
        "/budgets/{budget}/sectors/{sector}/transactions",
        transaction.CreateTransactionAction,
    },
    Route{
        "Transaction",
        "GET",
        "/budgets/{budget}/sectors/{sector}/transactions/{id}",
        transaction.GetTransactionAction,
    },
}
