package main

import (
	"backend/config"
	healthHttp "backend/handler/health/delivery/http"
	healthEntity "backend/handler/health/entity"
	healthUseCase "backend/handler/health/usecase"
	securityHttp "backend/handler/security/delivery/http"
	securityEntity "backend/handler/security/entity"
	securityUseCase "backend/handler/security/usecase"
	userHttp "backend/handler/user/delivery/http"
	userEntity "backend/handler/user/entity"
	userUseCase "backend/handler/user/usecase"

	"github.com/gin-gonic/gin"
	goutil "github.com/muhammadrivaldy/go-util"
)

func service(
	route *gin.Engine,
	config config.Configuration,
	validate goutil.Validation) {

	// call the function of method entity
	securityEntity, err := securityEntity.NewSecurityEntity(config)
	if err != nil {
		panic(err)
	}

	healthEntity, err := healthEntity.NewHealthEntity(config)
	if err != nil {
		panic(err)
	}

	userEntity, err := userEntity.NewUserEntity(config)
	if err != nil {
		panic(err)
	}

	// call the function of method useCase
	healthUseCase := healthUseCase.NewUseCase(healthEntity)
	securityUseCase := securityUseCase.NewSecurityUseCase(config, securityEntity, userEntity)
	userUseCase := userUseCase.NewUserUseCase(config, userEntity)

	// call the function of method endpoint
	healthHttp.NewEndpoint(route, securityUseCase, healthUseCase, validate)
	userHttp.NewEndpoint(config, route, securityUseCase, userUseCase, validate)
	securityHttp.NewEndpoint(config, route, securityUseCase, userUseCase, validate)
}
