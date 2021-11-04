package models

import "time"

type Task struct {
	Id          string `bson:"_id,omitempty" json:"id,omitempty"`
	Description string `bson:"description,omitempty" json:"description"`
	IsComplete  bool   `bson:"isComplete,omitempty" json:"isComplete"`
	DateCreated string `bson:"dateCreated,omitempty" json:"dateCreated"`
}

// Constructor
func NewTask() Task {
	instance := Task{}
	instance.Description = ""
	instance.IsComplete = false
	instance.DateCreated = time.Now().Format("2006-01-02")
	return instance
}
