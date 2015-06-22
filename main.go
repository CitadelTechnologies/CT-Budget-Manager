package main

import(
	"budget-manager/manager"
	"log"
	"net/http"
)

func main() {

	manager.InitMongo()
	defer manager.MongoDBConnection.Close()

	router := NewRouter()

	log.Fatal(http.ListenAndServe(":28515", router))
}
