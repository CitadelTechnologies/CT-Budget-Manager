package transaction

import(
	"ct-budget-manager/server"
	"net/http"
	"time"
	"github.com/gorilla/mux"
)

func CreateTransactionAction(w http.ResponseWriter, r *http.Request) {
	defer server.CatchException(w)

	params := mux.Vars(r)
	data := server.DecodeJsonRequest(r)
	var isset bool
	processedAt := time.Now()
	if processedAt, isset = data["processed_at"]; isset {
		processedAt = time.Parse(time.RFC3339, data.(string))
	}
	transaction := CreateTransaction(
		params["budget"],
		params["sector"],
		data["wording"].(string),
		data["description"].(string),
		data["type"].(string),
		data["amount"].(float64),
		processedAt,
	)
	server.SendJsonResponse(w, 201, transaction)
}

func GetTransactionAction(w http.ResponseWriter, r *http.Request) {
	defer server.CatchException(w)

	params := mux.Vars(r)

	server.SendJsonResponse(w, 200, GetTransaction(
		params["budget"],
		params["sector"],
		params["id"],
	))
}

func GetTransactionsAction(w http.ResponseWriter, r *http.Request) {
	defer server.CatchException(w)

	w.Header().Set("Access-Control-Allow-Origin", "*")

	params := mux.Vars(r)
	server.SendJsonResponse(w, 200, GetTransactions(
		params["budget"],
		params["sector"],
	))
}
