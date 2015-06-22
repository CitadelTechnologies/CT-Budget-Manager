package manager

import(
	"gopkg.in/mgo.v2"
)

var MongoDBConnection *mgo.Session

func InitMongo() {
	var err error
	MongoDBConnection, err = mgo.Dial("mongodb://localhost:27017")

	CheckError(err)
}
