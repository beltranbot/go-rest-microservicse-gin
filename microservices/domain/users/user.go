package users

import (
	"github.com/beltranbot/go-rest-microservice-gin/microservices/domain/httperrors"
)

// User struct
type User struct {
	ID        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

// Validate func
func (user User) Validate() *httperrors.HTTPError {
	if user.FirstName == "" {
		return httperrors.NewBadRequestError("invalid user first name")
	}
	if user.LastName == "" {
		return httperrors.NewBadRequestError("invalid user last name")
	}
	if user.Email == "" {
		return httperrors.NewBadRequestError("invalid user email")
	}

	return nil
}
