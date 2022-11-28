package usecase

import (
	"backend/config"
	"backend/handler/user"
	"backend/handler/user/entity"
)

type useCase struct {
	config     config.Configuration
	userEntity entity.UserEntity
}

// NewUserUseCase is a function for override interface
func NewUserUseCase(
	config config.Configuration,
	userEntity entity.UserEntity) user.IUserUseCase {
	return &useCase{
		config:     config,
		userEntity: userEntity}
}
