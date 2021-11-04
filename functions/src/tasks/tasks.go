package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/bymi15/go-mongo-serverless-crud-boilerplate/db"
	"github.com/bymi15/go-mongo-serverless-crud-boilerplate/db/models"
	"github.com/bymi15/go-mongo-serverless-crud-boilerplate/functions/src/utils"
)

func handler(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	client := db.InitMongoClient()
	id := request.QueryStringParameters["id"]
	var response *events.APIGatewayProxyResponse

	switch request.HTTPMethod {
	case "GET":
		if id != "" {
			// Get by id
			task, err := client.TaskService.GetTaskById(id)
			if err != nil {
				return &events.APIGatewayProxyResponse{
					StatusCode: 500,
				}, err
			}
			response = utils.CreateApiResponse(task, 200)
		} else {
			// Get all
			tasks, err := client.TaskService.GetTasks()
			if err != nil {
				return &events.APIGatewayProxyResponse{
					StatusCode: 500,
				}, err
			}
			response = utils.CreateApiResponse(tasks, 200)
		}
	case "POST":
		var task models.Task
		err := utils.ParseBody(request.Body, &task)
		if err != nil {
			return &events.APIGatewayProxyResponse{
				StatusCode: 400,
			}, err
		}
		err = client.TaskService.CreateTask(task)
		if err != nil {
			return &events.APIGatewayProxyResponse{
				StatusCode: 500,
			}, err
		}
		response = utils.CreateApiResponse(task, 201)
	case "PUT":
		var task models.Task
		err := utils.ParseBody(request.Body, &task)
		if err != nil {
			return &events.APIGatewayProxyResponse{
				StatusCode: 400,
			}, err
		}
		err = client.TaskService.UpdateTask(id, task)
		if err != nil {
			return &events.APIGatewayProxyResponse{
				StatusCode: 500,
			}, err
		}
		response = utils.CreateApiResponse(task, 200)
	case "DELETE":
		err := client.TaskService.DeleteTask(id)
		if err != nil {
			return &events.APIGatewayProxyResponse{
				StatusCode: 500,
			}, err
		}
		response = utils.CreateApiResponse("", 200)
	}

	return response, nil
}

func main() {
	// Make the handler available for Remote Procedure Call by AWS Lambda
	lambda.Start(handler)
}
