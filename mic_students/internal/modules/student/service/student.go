package service

import (
	"errors"

	"github.com/vitorsiqueirarecife/students/internal/modules/student/domain"
	"github.com/vitorsiqueirarecife/students/internal/modules/student/repository"
)

type StudentService interface {
	GetAll(page int, limit int) (*domain.GetAllStudentsResponse, error)
}

type studentService struct {
	repo repository.StudentRepository
}

func NewStudentService(repo repository.StudentRepository) StudentService {
	return &studentService{repo}
}

func (s *studentService) GetAll(page int, limit int) (*domain.GetAllStudentsResponse, error) {
	if page < 1 || limit < 1 {
		return nil, errors.New("Page and limit must be greater than 0")
	}

	if limit > 100 {
		return nil, errors.New("Limit must be less than or equal to 100")
	}

	res, err := s.repo.GetAll(page, limit)
	if err != nil {
		return nil, err
	}

	return &domain.GetAllStudentsResponse{
		Students: res.Students,
		Total:    res.Total,
	}, nil
}
