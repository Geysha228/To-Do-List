package mocks

import (
	"context"
	"fmt"

	"github.com/Geysha228/To-Do-List/internal/models"
)

type MockTaskRepository struct{}

func (m *MockTaskRepository) GetTodayTasksByUser(ctx context.Context, userID int) ([]models.Task, error) {
	return []models.Task{
		{ID: 1, Name: "Task 1", CreatorID: uint(userID)},
		{ID: 2, Name: "Task 2", CreatorID: uint(userID)},
	}, nil
}

type ErrorTaskRepository struct{}

func (r *ErrorTaskRepository) GetTodayTasksByUser(ctx context.Context, userID int) ([]models.Task, error) {
	return nil, fmt.Errorf("db err")
}
