package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"todo-backend/config"
	"todo-backend/models"
)

func GetTasks(c *gin.Context) {
	tasks := []models.Task{}
	err := config.DB.Find(&tasks)
	if err != nil {
		log.Fatal(err)
	}
	if len(tasks) == 0 {
		c.JSON(http.StatusNotFound,
			gin.H{
				"message": "tasks not found",
			},
		)
	}

	c.JSON(http.StatusOK, &tasks)
}

func GetTask(c *gin.Context) {}

func DeleteTask(c *gin.Context) {}

func UpdateTask(c *gin.Context) {}

func CreateTask(c *gin.Context) {}
