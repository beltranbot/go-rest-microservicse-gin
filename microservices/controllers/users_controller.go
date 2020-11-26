package controllers

import (
	"net/http"
	"strconv"

	"github.com/beltranbot/go-rest-microservice-gin/microservices/services"

	"github.com/beltranbot/go-rest-microservice-gin/microservices/domain/httperrors"
	"github.com/beltranbot/go-rest-microservice-gin/microservices/domain/users"
	"github.com/gin-gonic/gin"
)

var (
	// UsersController var
	UsersController = usersController{}
)

type usersController struct{}

// Get func
func (controller *usersController) Get(c *gin.Context) {
	userID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		httpErr := httperrors.NewBadRequestError("invalid user id")
		respond(c, http.StatusBadRequest, httpErr)
		return
	}
	user, getErr := services.UsersService.Get(userID)
	if getErr != nil {
		respond(c, getErr.Code, getErr)
		return
	}

	respond(c, http.StatusOK, user)
}

func respond(c *gin.Context, httpCode int, body interface{}) {
	isXML := (c.GetHeader("Accept") == "application/xml")
	if isXML {
		c.XML(httpCode, body)
		return
	}

	c.JSON(httpCode, body)
}

// Create func
func (controller *usersController) Create(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		httpErr := httperrors.NewBadRequestError("invalid request body")
		c.JSON(httpErr.Code, httpErr)
		return
	}

	createdUser, err := services.UsersService.Create(user)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusCreated, createdUser)
}
