package models

import "time"

type Task struct {
	Id          string `bson:"_id,omitempty" json:"id,omitempty"`
	Description string `json:"description"`
	IsComplete  bool   `json:"isComplete"`
	DateCreated string `json:"dateCreated"`
}

// Constructor
func NewTask() Task {
	instance := Task{}
	instance.Description = ""
	instance.IsComplete = false
	instance.DateCreated = time.Now().Format("2006-01-02")
	return instance
}
