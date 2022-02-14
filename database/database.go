package database

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
)

func MongoDBInitialization() {
	if dotErr := godotenv.Load(); dotErr != nil {
		log.Fatalln(dotErr.Error())
	}
	url := os.Getenv("MONGO_STRING_CONNECTION")
	dbName := os.Getenv("MONGO_DB_NAME")
	err := mgm.SetDefaultConfig(nil, dbName, options.Client().ApplyURI(url))
	if err != nil {
		log.Fatalln(err.Error())
	}
	fmt.Println("MONGODB CONNECTED")
}
