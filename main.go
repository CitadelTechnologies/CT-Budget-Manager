package main

import(
    "budget-manager/manager"
    "log"
    "net/http"
    "fmt"
    gs "github.com/Kern046/GleipnirServer"
)

func main() {

    gs.Initialize()

    fmt.Println("MongoDB server initialization")

    manager.InitMongo()
    defer manager.MongoDBConnection.Close()

    fmt.Println("MongoDB is ready")
    fmt.Println("Router initialization")

    router := NewRouter()

    fmt.Println("Server is listening on port " + gs.Server.DedicatedPort)

    log.Fatal(http.ListenAndServe(":" + gs.Server.DedicatedPort, router))
}
