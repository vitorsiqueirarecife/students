package domain

type HealthService interface {
	Check() string
}
