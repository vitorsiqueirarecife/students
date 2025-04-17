package service

import "github.com/vitorsiqueirarecife/students/internal/modules/healthcheck/domain"

type HealthUseCase struct{}

func NewHealthUseCase() domain.HealthService {
	return &HealthUseCase{}
}

func (h *HealthUseCase) Check() string {
	return "ok"
}
