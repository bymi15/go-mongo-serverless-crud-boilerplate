package main

import (
	"log"
	"net/http"

	"github.com/bymi15/go-mongo-serverless-crud-boilerplate/db"
	"github.com/bymi15/go-mongo-serverless-crud-boilerplate/db/models"
	"github.com/bymi15/go-mongo-serverless-crud-boilerplate/functions/src/utils"
)

func handler(w http.ResponseWriter, r *http.Request) {
	client := db.InitMongoClient()
	id := r.URL.Query().Get("id")

	utils.SetDefaultHeaders(w)
	var response []byte

	switch r.Method {
	case "GET":
		if id != "" {
			// Get by id
			task, err := client.TaskService.GetTaskById(id)
			if err != nil {
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				return
			}
			w.WriteHeader(http.StatusOK)
			response = utils.CreateApiResponse(task)
		} else {
			// Get all
			tasks, err := client.TaskService.GetTasks()
			if err != nil {
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				return
			}
			w.WriteHeader(http.StatusOK)
			response = utils.CreateApiResponse(tasks)
		}
	case "POST":
		task := models.NewTask()
		err := utils.ParseRequestBody(r, &task)
		if err != nil {
			log.Printf("Error: %v", err)
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		err = client.TaskService.CreateTask(task)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
		response = utils.CreateApiResponse(task)
	case "PUT":
		var task models.Task
		err := utils.ParseRequestBody(r, &task)
		if err != nil {
			log.Printf("Error: %v", err)
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		err = client.TaskService.UpdateTask(id, task)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		response = utils.CreateApiResponse(task)
	case "DELETE":
		err := client.TaskService.DeleteTask(id)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		response = utils.CreateApiResponse("")
	}

	w.Write(response)
}

func main() {
	utils.ServeFunction("/api/tasks", handler)
}
