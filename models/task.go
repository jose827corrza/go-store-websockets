package models

import "gorm.io/gorm"

type SubTask struct {
	ID        string
	Name      string `json:"name"`
	Completed bool   `json:"completed"`
	TaskID    string
}

type Task struct {
	gorm.Model
	Id          string
	Title       string `json:"title"`
	Description string `json:"description"`
	IsCompleted bool   `json:"isCompleted"`
	Priority    int    `json:"priority"`
	SubTasks    []SubTask
	UserID      string // 1:N
}

// type Priority byte

// const (
// 	Low    Priority = 1
// 	Medium          = 2
// 	High            = 3
// )
