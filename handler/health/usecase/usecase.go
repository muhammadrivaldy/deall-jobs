package usecase

import (
	"backend/handler/health"
	healthEntity "backend/handler/health/entity"

	goutil "github.com/muhammadrivaldy/go-util"
)

type useCase struct {
	logs         goutil.Logs
	healthEntity healthEntity.HealthEntity
}

// NewUseCase is a function for override interface
func NewUseCase(
	healthEntity healthEntity.HealthEntity) health.IHealthUseCase {
	return &useCase{
		healthEntity: healthEntity}
}
