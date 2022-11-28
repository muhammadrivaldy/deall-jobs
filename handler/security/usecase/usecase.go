package usecase

import (
	"backend/config"
	"backend/handler/security"
	"backend/handler/security/entity"
	entityUser "backend/handler/user/entity"
)

type useCase struct {
	config         config.Configuration
	securityEntity entity.SecurityEntity
	userEntity     entityUser.UserEntity
}

// NewSecurityUseCase is a function for override interface
func NewSecurityUseCase(
	config config.Configuration,
	securityEntity entity.SecurityEntity,
	userEntity entityUser.UserEntity) security.ISecurityUseCase {
	return &useCase{
		config:         config,
		securityEntity: securityEntity,
		userEntity:     userEntity}
}
