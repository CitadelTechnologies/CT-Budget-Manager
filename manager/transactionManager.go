package manager

import(
	"budget-manager/model"
	"net/http"
	"encoding/json"
	"io"
	"io/ioutil"
	"time"
	"gopkg.in/mgo.v2/bson"
)

func CreateTransaction(w http.ResponseWriter, r *http.Request) {
	var t model.Transaction
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	CheckError(err)
	err = r.Body.Close()
	CheckError(err)
	if err := json.Unmarshal(body, &t); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
    			panic(err)
		}
	}

	err = json.Unmarshal(body, &t)

	t.Date = time.Now()

  	err = MongoDBConnection.DB("test").C("budget").Insert(t)
	CheckError(err)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(t)
	CheckError(err)
}

func GetTransaction(w http.ResponseWriter, r *http.Request) {

	var t model.Transaction

	err := MongoDBConnection.DB("test").C("budget").Find(bson.M{}).One(&t)
	CheckError(err)

	err = json.NewEncoder(w).Encode(t)
	CheckError(err)

}

func GetTransactions(w http.ResponseWriter, r *http.Request) {

	var t model.Transactions
	err := MongoDBConnection.DB("test").C("budget").Find(nil).All(&t)
	CheckError(err)

	err = json.NewEncoder(w).Encode(t)
	CheckError(err)

}
