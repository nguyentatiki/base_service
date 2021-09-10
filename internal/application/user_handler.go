//Handle business logic
package app

import (
	"base_service/internal/domain/entities"
	"base_service/internal/domain/interfaces"
)

type UserHandler struct {
	interfaces.UserRepository
}

func NewUserHandler(repo interfaces.UserRepository) *UserHandler {
	return &UserHandler{repo}
}

func (handler *UserHandler) GetUser(user *entities.User) (*entities.User, error) {
	return handler.UserRepository.GetUser(user.Username)
}
