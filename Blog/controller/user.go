package controller

import (
	"narie/monarie-project/config"
	"narie/monarie-project/models"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	users := []models.User{}
	config.DB.Find(&users)
	c.JSON(200, &users)
}

func GetPosts(c *gin.Context) {
	posts := []models.Post{}
	config.DB.Find(&posts)
	c.JSON(200, &posts)
}

func DeleteUser(c *gin.Context) {
	var user models.User
	config.DB.Where("id=?", c.Param("id")).Delete(&user)
	c.JSON(202, &user)
}

func DeletePost(c *gin.Context) {
	var post models.Post
	config.DB.Where("id=?", c.Param("id")).Delete(&post)
	c.JSON(202, &post)
}

func CreateUser(c *gin.Context) {
	var user models.User
	c.BindJSON(&user)
	config.DB.Create(&user)
	c.JSON(200, &user)
}

func CreatePost(c *gin.Context) {
	var post models.Post
	c.BindJSON(&post)
	config.DB.Create(&post)
	c.JSON(200, &post)
}
