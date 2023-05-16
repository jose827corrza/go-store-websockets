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
	Title       string `json:"title"`
	Description string `json:"description"`
	IsCompleted bool   `json:"isCompleted"`
	Priority
	SubTasks []SubTask
	UserID   string // 1:N
}

type Priority byte

const (
	Low Priority = iota
	Medium
	High
)
