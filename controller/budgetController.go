package controller

import(
	"ct-budget-manager/manager"
	"net/http"
	"encoding/json"
	"io"
	"io/ioutil"
	"github.com/gorilla/mux"
)

/*
* GET request to get all budgets
*
* @param http.ResponseWriter w
* @param http.Request r
*/
func GetBudgets(w http.ResponseWriter, r *http.Request) {
    budgets := manager.GetBudgets()

  	w.Header().Set("Access-Control-Allow-Origin", "*")

  	if err := json.NewEncoder(w).Encode(budgets); err != nil {
      panic(err)
    }
}

/*
* GET request to get a budget by its ID
*
* @param http.ResponseWriter w
* @param http.Request r
*/
func GetBudget(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

  budget := manager.GetBudget(vars["id"])
  if budget == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if err := json.NewEncoder(w).Encode(budget); err != nil {
    panic(err)
  }
}

/*
* POST request to create a new Budget object
*
* @param http.ResponseWriter w
* @param http.Request r
*/
func CreateBudget(w http.ResponseWriter, r *http.Request) {
  var body []byte
  var err error
	if body, err = ioutil.ReadAll(io.LimitReader(r.Body, 1048576)); err != nil {
    panic(err)
  }
	if err = r.Body.Close(); err != nil {
    panic(err)
  }
  var data map[string]string
	if err = json.Unmarshal(body, &data); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusUnprocessableEntity)
		if err = json.NewEncoder(w).Encode(err); err != nil {
  			panic(err)
		}
	}
  budget := manager.CreateBudget(data["name"], data["description"])

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err = json.NewEncoder(w).Encode(budget); err != nil {
    panic(err)
  }
}
