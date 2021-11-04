package services

import (
	"context"
	"time"

	"github.com/bymi15/go-mongo-serverless-crud-boilerplate/db/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TaskService struct {
	Collection *mongo.Collection
}

func NewTaskService(db *mongo.Database, collectionName string) TaskService {
	return TaskService{
		Collection: db.Collection(collectionName),
	}

}

func (service TaskService) GetTasks() ([]models.Task, error) {
	var Tasks []models.Task

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := service.Collection.Find(ctx, bson.D{})
	if err != nil {
		defer cursor.Close(ctx)
		return Tasks, err
	}

	for cursor.Next(ctx) {
		Task := models.NewTask()
		err := cursor.Decode(&Task)
		if err != nil {
			return Tasks, err
		}
		Tasks = append(Tasks, Task)
	}

	return Tasks, nil
}

func (service TaskService) GetTaskById(id string) (models.Task, error) {
	Task := models.NewTask()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return Task, err
	}

	err = service.Collection.FindOne(ctx, bson.D{{"_id", objectId}}).Decode(&Task)
	if err != nil {
		return Task, err
	}
	return Task, nil

}

func (service TaskService) CreateTask(Task models.Task) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err := service.Collection.InsertOne(ctx, Task)
	if err != nil {
		return err
	}
	return nil
}

func (service TaskService) UpdateTask(id string, Task models.Task) error {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var data bson.M
	bytes, err := bson.Marshal(Task)
	if err != nil {
		return err
	}
	err = bson.Unmarshal(bytes, &data)
	if err != nil {
		return err
	}
	_, err = service.Collection.UpdateOne(
		ctx,
		bson.D{{"_id", objectId}},
		bson.D{{"$set", data}},
	)
	return err
}

func (service TaskService) DeleteTask(id string) error {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err = service.Collection.DeleteOne(ctx, bson.D{{"_id", objectId}})
	if err != nil {
		return err
	}
	return nil
}
