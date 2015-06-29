package main

import(
	"budget-manager/manager"
	"log"
	"net/http"
	"fmt"
	"os"
)

func main() {

    port := os.Args[1]

    fmt.Printf("%s\n", port)

	fmt.Println("MongoDB server initialization")

	manager.InitMongo()
	defer manager.MongoDBConnection.Close()

	fmt.Println("MongoDB is ready")
	fmt.Println("Router initialization")

	router := NewRouter()

	fmt.Println("Server is listening on port " + port)

	log.Fatal(http.ListenAndServe(":" + port, router))
}
