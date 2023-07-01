package routes

import (
	"github.com/gin-gonic/gin"

	"todo-backend/controller"
)

func UseRoute(router *gin.Engine) {
	router.GET("/tasks", controller.GetTasks)
	router.GET("/task", controller.GetTask)
	router.PUT("/tasks", controller.CreateTask)
	router.DELETE("/tasks", controller.DeleteTask)
}
