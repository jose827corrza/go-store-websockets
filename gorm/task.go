package orm

import (
	"context"
	"log"

	"github.com/jose827corrza/go-store-websockets/models"
)

func (repo *PostgresRepository) CreateTask(ctx context.Context, task *models.Task) error {
	result := repo.DB.Create(&models.Task{
		Title:       task.Title,
		Description: task.Description,
		IsCompleted: task.IsCompleted,
		Priority:    task.Priority,
		UserID:      task.UserID,
		// SubTasks:    task.SubTasks,
		Id: task.Id,
	})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repo *PostgresRepository) GetAllTasksByUserId(ctx context.Context, userId string) ([]*models.Task, error) {
	var tasks []*models.Task

	result := repo.DB.Where("user_id=?", userId).Find(&tasks)
	if result.Error != nil {
		return nil, result.Error
	}
	return tasks, nil
}

func (repo *PostgresRepository) EditATaskByTaskId(ctx context.Context, taskId string, task *models.Task) (*models.Task, error) {
	result := repo.DB.Where("id = ?", taskId).Updates(&task).First(&task) // .First is used to be able to Return an Error of type ErrRecordNotFound
	if result.Error != nil {
		log.Print("entro al error")
		return nil, result.Error
	}
	return task, result.Error
}

func (repo *PostgresRepository) DeleteATaskByTaskId(ctx context.Context, taskId string) error {
	taskToDelete := repo.DB.Where("id=?", taskId).First(&models.Task{})
	if taskToDelete.Error != nil {
		log.Print("error does not exist")
		return taskToDelete.Error
	}
	result := repo.DB.Where("id=?", taskId).Delete(&models.Task{})
	if result.Error != nil {
		log.Print("error cannot be deleted")
		return taskToDelete.Error
	}
	return nil
}

func (repo *PostgresRepository) CreateASubTask(ctx context.Context, subTask *models.SubTask, taskId string) (*models.Task, error) {
	var task *models.Task
	var updatedTask *models.Task

	result := repo.DB.Where("id=?", taskId).Model(&models.Task{}).First(&task)
	if result.Error != nil {
		return nil, result.Error
	}
	task.SubTasks = append(task.SubTasks, models.SubTask{
		Name:      subTask.Name,
		Completed: subTask.Completed,
		ID:        subTask.ID,
		TaskID:    taskId,
	})

	repo.DB.Save(&models.Task{
		Id:          task.Id,
		Title:       task.Title,
		IsCompleted: task.IsCompleted,
		Priority:    task.Priority,
		UserID:      task.UserID,
		SubTasks:    task.SubTasks,
		Description: task.Description,
	})
	updatedResult := repo.DB.Where("id=?", taskId).Model(&models.Task{}).Preload("SubTasks").First(&updatedTask)
	if updatedResult.Error != nil {
		return nil, result.Error
	}
	return updatedTask, nil
	// err := repo.DB.Model(&models.Task{}).Preload("SubTasks").Where("id=?", taskId).First(&task).Error
	// if err != nil {
	// 	return nil, result.Error
	// }
	// return task, nil
}
