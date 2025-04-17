package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/vitorsiqueirarecife/students/internal/modules/student/domain"
)

type StudentRepository interface {
	GetAll() ([]domain.Student, error)
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

func (r *studentRepository) GetAll() ([]domain.Student, error) {
	var students []domain.Student
	if err := r.db.Find(&students).Error; err != nil {
		return nil, err
	}
	return students, nil
}
