package request

import (
	"errors"

	"github.com/randijulio13/gogin/model"
)

type CreateUserRequest struct {
	Password    string   `json:"password" binding:"required"`
	Email       string   `json:"email,omitempty"`
	Nip         string   `json:"nip" binding:"required"`
	Name        string   `json:"name" binding:"required"`
	Position    string   `json:"position" binding:"required"`
	Permissions []string `json:"permissions" binding:"required"`
}

func (request *CreateUserRequest) Validate() error {
	var user model.User
	user.GetByEmailOrNip(request.Email, request.Nip)

	if len(request.Email) != 0 {
		if request.Email == user.Email && request.Nip == user.Nip {
			return errors.New("user already registered")
		} else if request.Email == user.Email {
			return errors.New("email already registered")
		}
	} else {
		if request.Nip == user.Nip {
			return errors.New("nip already registered")
		}
	}

	return nil
}
