package services

import (
	"fmt"

	"github.com/beltranbot/go-rest-microservice-gin/microservices/domain/httperrors"
	"github.com/beltranbot/go-rest-microservice-gin/microservices/domain/users"
)

var (
	// UsersService var
	UsersService          = usersService{}
	registeredUsers       = map[int64]*users.User{}
	currentUserID   int64 = 1
)

type usersService struct{}

// Create func
func (service *usersService) Create(user users.User) (*users.User, *httperrors.HTTPError) {
	if err := user.Validate(); err != nil {
		return nil, err
	}
	user.ID = currentUserID
	currentUserID++
	registeredUsers[user.ID] = &user

	return &user, nil
}

// Get func
func (service *usersService) Get(userID int64) (*users.User, *httperrors.HTTPError) {
	if user := registeredUsers[userID]; user != nil {
		return user, nil
	}

	return nil, httperrors.NewNotFoundError(fmt.Sprintf("user %d not found", userID))
}
