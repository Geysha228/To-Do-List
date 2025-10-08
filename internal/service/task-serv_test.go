package service

import (
	"context"
	"testing"

	"github.com/Geysha228/To-Do-List/internal/repository/mocks"
)

func TestTaskService_GetAllTasksTodayByUser(t *testing.T) {
	mock := &mocks.MockTaskRepository{}
	svc := NewTaskService(mock)

	ctx := context.Background()
	userID := 42

	tasks, err := svc.GetAllTasksTodayByUser(ctx, userID)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(tasks) != 2 {
		t.Fatalf("expected 2 tasks, but got %d", len(tasks))
	}

	for _, task := range tasks {
		if task.CreatorID != uint(userID) {
			t.Fatalf("in task %d unxecpected creatorID: %d, but expected %d", task.ID, task.CreatorID, userID)
		}
	}
}

func TestTaskService_GetAllTaskTodayByUser_Error(t *testing.T) {
	mock := &mocks.ErrorTaskRepository{}
	service := NewTaskService(mock)

	ctx := context.Background()
	userID := 1

	_, err := service.GetAllTasksTodayByUser(ctx, userID)
	if err == nil {
		t.Fatalf("expected error, got nil")
	}

}
