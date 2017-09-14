package controller

import(
	"ct-budget-manager/manager"
	"net/http"
	"encoding/json"
	"io"
	"io/ioutil"
)

/*
* POST request to create a new Budget object
*
* @param http.ResponseWriter w
* @param http.Request r
* @param int amount
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
