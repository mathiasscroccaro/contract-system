package mongo_adapter

import(
	"os"
	"log"
	"fmt"
	"context"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo"
)

type DB struct {
	MongoInstance *mongo.Client
}

var dBInstance DB

func GetDBInstance() DB {
	return dBInstance
}

func BuildUrlFromEnv() string {
	username := os.Getenv("MONGO_USERNAME")
	password := os.Getenv("MONGO_PASSWORD")
	database := os.Getenv("MONGO_DATABASE")
	host := os.Getenv("MONGO_HOST")
	port := os.Getenv("MONGO_PORT")

	return fmt.Sprintf("mongodb://%s:%s@%s:%s/%s?authSource=admin", username, password, host, port, database)
}

func (db *DB) Init(dbUrl string) {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(dbUrl))
	if err != nil {
		log.Fatal(err)
	}
	db.MongoInstance = client
}