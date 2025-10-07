package router

import (
	"github.com/Geysha228/To-Do-List/internal/api/handler"
	"github.com/gin-gonic/gin"
)

func SetupRouter(taskHandler *handler.TaskHandler) *gin.Engine {
	r := gin.Default()

	tasks := r.Group("/tasks")
	{
		tasks.GET("/", taskHandler.GetTasks)
	}
	return r
}
