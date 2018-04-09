package budget

import(
	"ct-budget-manager/server"
	"net/http"
	"github.com/gorilla/mux"
)

func GetBudgetsAction(w http.ResponseWriter, r *http.Request) {
	defer server.CatchException(w)

  	w.Header().Set("Access-Control-Allow-Origin", "*")

  	server.SendJsonResponse(w, 200, GetBudgets())
}

func GetBudgetAction(w http.ResponseWriter, r *http.Request) {
	defer server.CatchException(w)

	server.SendJsonResponse(w, 200, GetBudget(mux.Vars(r)["slug"]))
}

func CreateBudgetAction(w http.ResponseWriter, r *http.Request) {
	defer server.CatchException(w)

	data := server.DecodeJsonRequest(r)

	server.SendJsonResponse(w, 201, CreateBudget(
		data["name"].(string),
		data["description"].(string),
	))
}
