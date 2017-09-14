package manager

import(
	"ct-budget-manager/model"
	"net/http"
	"encoding/json"
	"io"
	"io/ioutil"
	"time"
	"gopkg.in/mgo.v2/bson"
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
	var t model.Transaction
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	CheckError(err)
	err = r.Body.Close()
	CheckError(err)
	if err := json.Unmarshal(body, &t); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusUnprocessableEntity)
		if err := json.NewEncoder(w).Encode(err); err != nil {
    			panic(err)
		}
	}

	err = json.Unmarshal(body, &t)

	t.Id = bson.NewObjectId()
	t.Date = time.Now()

  	err = MongoDBConnection.DB(MongoDBName).C("budget").Insert(t)
	CheckError(err)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(t)
	CheckError(err)
}

func GetTransaction(w http.ResponseWriter, r *http.Request) {

	var t model.Transaction

	vars := mux.Vars(r)

	if !bson.IsObjectIdHex(vars["id"]) {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	err := MongoDBConnection.DB(MongoDBName).C("budget").FindId(bson.ObjectIdHex(vars["id"])).One(&t)
	CheckError(err)

	err = json.NewEncoder(w).Encode(t)
	CheckError(err)

}

func GetTransactions(w http.ResponseWriter, r *http.Request) {

	t := make(model.Transactions, 0)
	err := MongoDBConnection.DB(MongoDBName).C("budget").Find(nil).All(&t)
	CheckError(err)

	w.Header().Set("Access-Control-Allow-Origin", "*")

	err = json.NewEncoder(w).Encode(t)
	CheckError(err)

}
