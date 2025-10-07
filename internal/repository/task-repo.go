package repository

import (
	"github.com/Geysha228/To-Do-List/internal/models"
	"gorm.io/gorm"
)

type TaskRepositoryInterface interface {
	GetTasks() ([]models.Task, error)
}

type TaskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(database *gorm.DB) *TaskRepository {
	return &TaskRepository{db: database}
}

func (r *TaskRepository) GetTasks() ([]models.Task, error) {
	var tasks []models.Task
	err := r.db.Find(&tasks).Error
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func (r *TaskRepository) CreateTask(task models.Task) error {

}
