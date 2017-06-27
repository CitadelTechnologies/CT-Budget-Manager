package main

import(
    "ct-budget-manager/manager"
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

    fmt.Println("Server is listening on port 80")

    log.Fatal(http.ListenAndServe(":80", router))
}
