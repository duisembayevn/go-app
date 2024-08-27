package users

import (
	"go_app/models"
)

type UserStore interface {
	CreateUser(username string, firstName string, lastName, password string) error
	GetUserByUsername(username string) (*models.User, error)
}
