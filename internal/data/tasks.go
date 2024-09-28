package data

import (
	"database/sql"
	"time"
)

type Task struct {
	ID 							int					`json:"id"`
	Name						string			`json:"name"`
	Description			string			`json:"description,omitempty"`
	Status 					string			`json:"status"`
	CreatedAt       time.Time 	`json:"-"`
}

type TaskModel struct {
	db 	*sql.DB
}

func (t TaskModel) GetAll() ([]Task, error){
	// TODO collect list of tasks from a real database 
	tasks := []Task{
		{
			ID: 1,
			Name: "Task in Progress",
			Description: "",
			Status: "inprogress",
		},
		{
			ID: 2,
			Name: "Task Completed",
			Description: "",
			Status: "completed",
		},		
		{
			ID: 3,
			Name: "Task Won't Do",
			Description: "",
			Status: "won't do",
		},
		{
			ID: 4,
			Name: "Task To do",
			Description: "Work on a Challenge on devChanllengies.io, learn Go With React (TypeScript)",
			Status: "won't do",
		},
	}

	return tasks, nil
}