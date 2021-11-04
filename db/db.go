package db

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/bymi15/go-mongo-serverless-crud-boilerplate/db/services"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDbClient struct {
	db          *mongo.Database
	TaskService services.TaskService
}

func InitMongoClient() MongoDbClient {
	uri := os.Getenv("CONNECTION_URI")
	clientOptions := options.Client().ApplyURI(uri)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	db := client.Database(os.Getenv("DB_NAME"))
	return MongoDbClient{
		db:          db,
		TaskService: services.NewTaskService(db, "tasks"),
		// New services can be added here
	}

}
