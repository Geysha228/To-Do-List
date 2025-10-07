package handler

import (
	"net/http"

	"github.com/Geysha228/To-Do-List/internal/service"
	"github.com/gin-gonic/gin"
)

type TaskHandler struct {
	TaskService service.TaskServiceInterface
}

func NewTaskHandler(taskService service.TaskServiceInterface) *TaskHandler {
	return &TaskHandler{TaskService: taskService}
}

func (h *TaskHandler) GetTasks(c *gin.Context) {

	tasks, err := h.TaskService.GetAllTasks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, tasks)
}
