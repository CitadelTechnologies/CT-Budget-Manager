package main

import (
	"net/http"

	"github.com/gorilla/mux"

	"ct-budget-manager/controller"
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
				"Create Budget",
				"POST",
				"/budgets",
				controller.CreateBudget,
		},
		Route{
				"Get Budgets",
				"GET",
				"/budgets",
				controller.GetBudgets,
		},
		Route{
				"Get Budget",
				"GET",
				"/budgets/{id}",
				controller.GetBudget,
		},
    Route{
        "Transactions",
        "GET",
        "/transactions",
        controller.GetTransactions,
    },
    Route{
        "Create Transaction",
        "POST",
        "/budgets/{id}/transactions",
        controller.CreateTransaction,
    },
    Route{
        "Transaction",
        "GET",
        "/transactions/{id}",
        controller.GetTransaction,
    },
}
