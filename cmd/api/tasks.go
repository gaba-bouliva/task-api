package main

import (
	"net/http"
)

func (app * application) listTasksHandler(w http.ResponseWriter, r *http.Request) {
	tasks,err := app.models.Task.GetAll()
	if err != nil {
		app.logger.Println(err)
		return
	}

	err = app.writeJSON(w, jsonPayload{"tasks": tasks}, http.StatusOK, nil)	
	if err != nil {
		app.logger.Println(err)
		return 
	}

}

