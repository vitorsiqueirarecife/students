package service

import (
	"github.com/vitorsiqueirarecife/students/internal/modules/student/domain"
	"github.com/vitorsiqueirarecife/students/internal/modules/student/repository"
)

type StudentService interface {
	GetAll() ([]domain.Student, error)
}

type studentService struct {
	repo repository.StudentRepository
}

func NewStudentService(repo repository.StudentRepository) StudentService {
	return &studentService{repo}
}

func (s *studentService) GetAll() ([]domain.Student, error) {
	return s.repo.GetAll()
}
