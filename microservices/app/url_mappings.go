package app

import (
	"github.com/beltranbot/go-rest-microservice-gin/microservices/controllers"
)

func mapUrls() {
	router.POST("/users", controllers.UsersController.Create)
	router.GET("/users/:id", controllers.UsersController.Get)
}
