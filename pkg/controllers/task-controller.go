package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/JoaNMiFTW/go-todo-list/pkg/models"
	"github.com/JoaNMiFTW/go-todo-list/pkg/utils"
	"github.com/gorilla/mux"
)

func GetTask(w http.ResponseWriter, r *http.Request) {
	newTasks := models.GetAllTasks()
	res, _ := json.Marshal(newTasks)

	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetTaskById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskId := vars["taskId"]

	ID, err := strconv.ParseInt(taskId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}

	taskDetails, _ := models.GetTaskById(ID)
	res, _ := json.Marshal(taskDetails)

	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	CreateTask := &models.Task{}
	utils.ParseBody(r, CreateTask)

	t := CreateTask.CreateTask()
	res, _ := json.Marshal(t)

	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskId := vars["taskId"]

	ID, err := strconv.ParseInt(taskId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}

	task := models.DeleteTask(ID)
	res, _ := json.Marshal(task)

	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	var updateTask = &models.Task{}
	utils.ParseBody(r, updateTask)
	vars := mux.Vars(r)
	taskId := vars["taskId"]
	ID, err := strconv.ParseInt(taskId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	taskDetails, db := models.GetTaskById(ID)
	if updateTask.Title != "" {
		taskDetails.Title = updateTask.Title
	}
	if updateTask.Description != "" {
		taskDetails.Description = updateTask.Description
	}

	db.Save(&taskDetails)

	res, _ := json.Marshal(taskDetails)

	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
