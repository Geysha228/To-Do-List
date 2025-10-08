package service

import (
	"context"

	"github.com/Geysha228/To-Do-List/internal/models"
	"github.com/Geysha228/To-Do-List/internal/repository"
)

type TaskServiceInterface interface {
	GetAllTasksTodayByUser(ctx context.Context, userID int) ([]models.Task, error)
}

type TaskService struct {
	repo repository.TaskRepositoryInterface
}

func NewTaskService(repository repository.TaskRepositoryInterface) *TaskService {
	return &TaskService{repo: repository}
}

func (s *TaskService) GetAllTasksTodayByUser(ctx context.Context, userID int) ([]models.Task, error) {
	return s.repo.GetTodayTasksByUser(ctx, userID)
}
