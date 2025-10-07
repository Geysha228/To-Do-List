package service

import (
	"github.com/Geysha228/To-Do-List/internal/models"
	"github.com/Geysha228/To-Do-List/internal/repository"
)

type TaskServiceInterface interface {
	GetAllTasks() ([]models.Task, error)
}

type TaskService struct {
	repo repository.TaskRepositoryInterface
}

func NewTaskService(repository repository.TaskRepositoryInterface) *TaskService {
	return &TaskService{repo: repository}
}

func (s *TaskService) GetAllTasks() ([]models.Task, error) {
	return s.repo.GetTasks()
}
