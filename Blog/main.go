package main

import (
	"narie/monarie-project/config"
	"narie/monarie-project/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	config.Connect()
	routes.UseRoute(router)
	router.Run(":8080")
}
