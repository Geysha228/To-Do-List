package repository

import (
	"context"

	"github.com/Geysha228/To-Do-List/internal/models"
	"gorm.io/gorm"
)

type TaskRepositoryInterface interface {
	GetTodayTasksByUser(ctx context.Context, UserID int) ([]models.Task, error)
}

type TaskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(database *gorm.DB) *TaskRepository {
	return &TaskRepository{db: database}
}

func (r *TaskRepository) GetTodayTasksByUser(ctx context.Context, UserID int) ([]models.Task, error) {
	var tasks []models.Task
	err := r.db.WithContext(ctx).
		Where("creator_id = ? AND DATE(date_of_remind) = CURRENT_DATE", UserID).
		Find(&tasks).Error
	if err != nil {
		return nil, err
	}
	return tasks, nil
}
