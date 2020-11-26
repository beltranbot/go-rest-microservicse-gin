package controllers

import (
	"net/http"

	"github.com/beltranbot/go-rest-microservice-gin/testable-code/services"
	"github.com/gin-gonic/gin"
)

// Ping func
func Ping(c *gin.Context) {
	result, err := services.PingService.HandlePing()
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	}
	c.String(http.StatusOK, result)
}
