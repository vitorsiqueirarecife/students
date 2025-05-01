package service

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/vitorsiqueirarecife/students/internal/modules/student/domain"
	"github.com/vitorsiqueirarecife/students/internal/modules/student/repository"
)

type StudentService interface {
	GetAll(page int, limit int) (*domain.GetAllStudentsResponse, error)
	Create(student *domain.Student) error
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

	return res, nil
}

func (s *studentService) Create(student *domain.Student) error {
	fmt.Println(1)
	if student.Name == "" {
		return errors.New("Name is required")
	}
	if student.Grade < 1 || student.Grade > 12 {
		return errors.New("Grade must be between 1 and 12")
	}

	student.ID = uuid.New().String()

	err := s.repo.Create(student)
	if err != nil {
		return err
	}

	return nil
}
