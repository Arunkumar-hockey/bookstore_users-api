package services

import (
	"userapi/domain/users"
)

func CreateUser(user users.User) (*users.User, error) {
	return &user, nil
}