package main

import(
	"budget-manager/manager"
	"log"
	"net/http"
	"fmt"
)

func main() {

	fmt.Println("MongoDB server initialization")

	manager.InitMongo()
	defer manager.MongoDBConnection.Close()

	fmt.Println("MongoDB is ready")
	fmt.Println("Router initialization")

	router := NewRouter()

	fmt.Println("Server is ready")

	log.Fatal(http.ListenAndServe(":28515", router))
}
