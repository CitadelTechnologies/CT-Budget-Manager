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
* POST request to create a new Transaction object
*
* @param string wording
* @param string description
* @param int amount
*/
func CreateTransaction(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
  var body []byte
  var err error
  if body, err = ioutil.ReadAll(io.LimitReader(r.Body, 1048576)); err != nil {
    panic(err)
  }
	if err = r.Body.Close(); err != nil {
    panic(err)
  }
  var data map[string]interface{}
	if err = json.Unmarshal(body, &data); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusUnprocessableEntity)
		if err = json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}
  transaction := manager.CreateTransaction(
    vars["id"],
    data["wording"],
    data["description"],
    data["type"],
    data["sector"],
    data["amount"],
  )

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err = json.NewEncoder(w).Encode(transaction); err != nil {
    panic(err)
  }
}

func GetTransaction(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

  transaction := manager.GetTransaction(vars["id"])
	if transaction == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
  if err := json.NewEncoder(w).Encode(transaction); err != nil {
    panic(err)
  }
}

func GetTransactions(w http.ResponseWriter, r *http.Request) {
  transactions := manager.GetTransactions()

	w.Header().Set("Access-Control-Allow-Origin", "*")

	if err := json.NewEncoder(w).Encode(transactions); err != nil {
    panic(err)
  }
}
