package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/vitorsiqueirarecife/students/internal/modules/student/domain"
)

type StudentRepository interface {
	GetAll(page int, limit int) (*domain.GetAllStudentsResponse, error)
}

type studentRepository struct {
	db *gorm.DB
}

type StudentRepositoryOptions struct {
	DB *gorm.DB
}

func NewStudentRepository(opt StudentRepositoryOptions) StudentRepository {
	return &studentRepository{
		db: opt.DB,
	}
}

func (r *studentRepository) GetAll(page int, limit int) (*domain.GetAllStudentsResponse, error) {
	var students []domain.Student
	var total int

	err := r.db.Model(&domain.Student{}).Count(&total).Error
	if err != nil {
		return nil, err
	}

	err = r.db.Offset((page - 1) * limit).Limit(limit).Find(&students).Error
	if err != nil {
		return nil, err
	}
	return &domain.GetAllStudentsResponse{
		Students: students,
		Total:    total,
	}, nil
}
