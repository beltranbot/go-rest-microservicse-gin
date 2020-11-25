package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var (
	users           = map[int64]*User{}
	currentID int64 = 1
	router          = gin.Default()
)

type httepError struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

// User struct
type User struct {
	ID        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

func main() {
	mapUrls()

	router.Run(":8080")
}

func mapUrls() {
	router.POST("/users", Create)
	router.GET("/users/:id", Get)
}

// Get func
func Get(c *gin.Context) {
	userID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		httpErr := httepError{
			Message: "invalid user id",
			Code:    http.StatusBadRequest,
		}
		c.JSON(httpErr.Code, httpErr)
		return
	}
	user := users[userID]
	if user == nil {
		httpErr := httepError{
			Message: "not found",
			Code:    http.StatusNotFound,
		}
		c.JSON(httpErr.Code, httpErr)
		return
	}

	c.JSON(http.StatusOK, user)
}

// Create func
func Create(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		httpErr := httepError{
			Message: "invalid json body",
			Code:    http.StatusBadRequest,
		}
		c.JSON(httpErr.Code, httpErr)
		return
	}

	user.ID = currentID
	currentID++
	users[user.ID] = &user

	// return created user
	c.JSON(http.StatusCreated, user)
}
