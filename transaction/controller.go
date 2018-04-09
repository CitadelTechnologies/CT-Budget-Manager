package transaction

import(
	"ct-budget-manager/server"
	"net/http"
	"github.com/gorilla/mux"
)

func CreateTransactionAction(w http.ResponseWriter, r *http.Request) {
	defer server.CatchException(w)

	data := server.DecodeJsonRequest(r)
	transaction := CreateTransaction(
		mux.Vars(r)["id"],
		data["wording"].(string),
		data["description"].(string),
		data["type"].(string),
		data["amount"].(string),
	)
	server.SendJsonResponse(w, 201, transaction)
}

func GetTransactionAction(w http.ResponseWriter, r *http.Request) {
	defer server.CatchException(w)

	server.SendJsonResponse(w, 200, GetTransaction(mux.Vars(r)["id"]))
}

func GetTransactionsAction(w http.ResponseWriter, r *http.Request) {
	defer server.CatchException(w)

	w.Header().Set("Access-Control-Allow-Origin", "*")

	server.SendJsonResponse(w, 200, GetTransactions())
}
