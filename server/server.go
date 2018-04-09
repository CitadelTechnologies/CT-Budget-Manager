package server

import(
	"gopkg.in/mgo.v2"
	"os"
)

type Application struct {
    Database *mgo.Database
}

var App Application

func init() {
	connection, err := mgo.Dial("mongodb://" + os.Getenv("MONGO_HOST") + ":" + os.Getenv("MONGO_PORT"))
	if err != nil {
		panic(err)
	}

	App.Database = connection.DB(os.Getenv("MONGO_DBNAME"))
}
