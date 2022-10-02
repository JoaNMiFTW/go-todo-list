package models

import (
	"github.com/JoaNMiFTW/go-todo-list/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Task struct {
	gorm.Model
	Title       string `json:"title"`
	Description string `json:"description"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Task{})
}

func (t *Task) CreateTask() *Task {
	db.NewRecord(t)
	db.Create(&t)
	return t
}

func GetAllTasks() []Task {
	var Tasks []Task
	db.Find(&Tasks)
	return Tasks
}

func GetTaskById(ID int64) (*Task, *gorm.DB) {
	var getTask Task
	db := db.Where("ID=?", ID).Find(&getTask)
	return &getTask, db
}

func DeleteTask(ID int64) Task {
	var task Task
	db.Where("ID=?", ID).Delete(task)
	return task
}
