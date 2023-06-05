package routes

import (
	"narie/monarie-project/controller"

	"github.com/gin-gonic/gin"
)

func UseRoute(router *gin.Engine) {
	router.GET("/users", controller.GetUsers)
	router.GET("/posts", controller.GetPosts)
	router.POST("/user", controller.CreateUser)
	router.POST("/post", controller.CreatePost)
	router.DELETE("/user", controller.DeleteUser)
	router.DELETE("/post", controller.DeletePost)
}
