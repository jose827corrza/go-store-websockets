package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/jose827corrza/go-store-websockets/dtos"
	"github.com/jose827corrza/go-store-websockets/models"
	"github.com/jose827corrza/go-store-websockets/repository"
	"github.com/jose827corrza/go-store-websockets/server"
	"github.com/jose827corrza/go-store-websockets/utils"
	"github.com/jose827corrza/go-store-websockets/validators"
)

func CreateTaskHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var newTask dtos.TaskRequest

		err := validators.ValidateTask(&newTask, w, r)
		if err != nil {
			utils.ErrorResponse(400, err.Error(), w)
			return
		}

		tokenString := strings.TrimSpace(r.Header.Get("Authorization"))
		userInfo, err, status := utils.GetUserInfoByJWTToken(tokenString, s, r)
		if err != nil {
			utils.ErrorResponse(status, err.Error(), w)
			return
		}

		id := uuid.New()

		var task = models.Task{
			Id:          id.String(),
			Title:       newTask.Title,
			Description: newTask.Description,
			IsCompleted: false,
			UserID:      userInfo.Id,
			Priority:    newTask.Priority,
		}

		err = repository.CreateTask(r.Context(), &task)
		if err != nil {
			utils.ErrorResponse(500, err.Error(), w)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(task)
	}
}

func GetAllTaskByUserIdHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// path:=mux.Vars(r)

		tokenString := strings.TrimSpace(r.Header.Get("Authorization"))
		userInfo, err, status := utils.GetUserInfoByJWTToken(tokenString, s, r)
		if err != nil {
			utils.ErrorResponse(status, err.Error(), w)
			return
		}
		log.Print(userInfo.Id)
		tasks, err := repository.GetAllTasksByUserId(r.Context(), userInfo.Id)
		if err != nil {
			utils.ErrorResponse(500, err.Error(), w)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(tasks)
	}
}

func EditATaskByTaskId(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := mux.Vars(r)
		var editTask dtos.TaskUpdate

		err := validators.ValidateUpdateTask(&editTask, w, r)
		if err != nil {
			utils.ErrorResponse(400, err.Error(), w)
			return
		}
		var task = models.Task{
			Title:       editTask.Title,
			Description: editTask.Description,
			IsCompleted: editTask.IsComplete,
		}
		editedTask, err := repository.EditATaskByTaskId(r.Context(), path["taskId"], &task)
		if err != nil {
			utils.ErrorResponse(500, err.Error(), w)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(editedTask)
	}
}

func DeleteATaskByTaskId(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := mux.Vars(r)

		err := repository.DeleteATaskByTaskId(r.Context(), path["taskId"])
		if err != nil {
			utils.ErrorResponse(500, err.Error(), w)
			return
		}

		message := fmt.Sprintf(`{"message": task with ID: %s has been deleted}`, path["taskId"])
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(message)
	}
}

func CreateASubTaskForATaskHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := mux.Vars(r)
		var newSubTask dtos.SubTask

		err := validators.ValidateSubTask(&newSubTask, w, r)
		if err != nil {
			utils.ErrorResponse(400, err.Error(), w)
			return
		}
		id := uuid.New()
		var subTask = models.SubTask{
			Name:      newSubTask.Name,
			Completed: false,
			ID:        id.String(),
		}

		task, err := repository.CreateASubTask(r.Context(), &subTask, path["taskId"])
		if err != nil {
			utils.ErrorResponse(500, err.Error(), w)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(task)
	}
}
