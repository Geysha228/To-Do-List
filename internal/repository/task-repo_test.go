package repository

import (
	"context"
	"testing"
	"time"

	"github.com/Geysha228/To-Do-List/internal/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestGetTodayTasksByUser(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to open db: %v", err)
	}

	if err := db.AutoMigrate(&models.Task{}, &models.User{}); err != nil {
		t.Fatalf("failed to migrate: %v", err)
	}

	user1 := models.User{ID: 1, TelegramID: 1, Username: "alice"}
	user2 := models.User{ID: 2, TelegramID: 2, Username: "bob"}
	db.Create(&user1)
	db.Create(&user2)

	now := time.Now()
	db.Create(&models.Task{Name: "Task1", CreatorID: 1, DateOfRemind: &now})
	db.Create(&models.Task{Name: "Task2", CreatorID: 1, DateOfRemind: &now})
	db.Create(&models.Task{Name: "Task3", CreatorID: 2, DateOfRemind: &now})

	repo := NewTaskRepository(db)

	ctx := context.Background()

	tasks, err := repo.GetTodayTasksByUser(ctx, 1)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(tasks) != 2 {
		t.Fatalf("expected 2 tasks, got %d", len(tasks))
	}

	for _, task := range tasks {
		if task.CreatorID != 1 {
			t.Errorf("unexpected creator_id: got %d", task.CreatorID)
		}
	}
}
