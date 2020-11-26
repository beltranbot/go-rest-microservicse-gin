package main

import (
	"github.com/beltranbot/go-rest-microservice-gin/testable-code/controllers"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func main() {
	router.GET("/ping", controllers.Ping)
	router.Run(":8080")
}
