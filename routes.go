package main

import (
	"net/http"

	"github.com/gorilla/mux"

	"ct-budget-manager/manager"
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
        "Transactions",
        "GET",
        "/transactions",
        manager.GetTransactions,
    },
    Route{
        "Create Transaction",
        "POST",
        "/transactions",
        manager.CreateTransaction,
    },
    Route{
        "Transaction",
        "GET",
        "/transactions/{id}",
        manager.GetTransaction,
    },
}
