package manager

import(
	"fmt"
	"gopkg.in/mgo.v2"
	"os"
)

var MongoDBConnection *mgo.Session

func InitMongo() {
	var err error

	fmt.Println("mongodb://" + os.Getenv("MONGO_HOST") + ":" + os.Getenv("MONGO_PORT"))

	MongoDBConnection, err = mgo.Dial("mongodb://" + os.Getenv("MONGO_HOST") + ":" + os.Getenv("MONGO_PORT"))

	CheckError(err)
}
