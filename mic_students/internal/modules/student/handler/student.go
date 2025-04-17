package handler

import (
	"net/http"

	"github.com/vitorsiqueirarecife/students/internal/modules/student/domain"
	"github.com/vitorsiqueirarecife/students/internal/modules/student/service"
)

type StudentHandler struct {
	service service.StudentService
}

func NewStudentHandler(service service.StudentService) *StudentHandler {
	return &StudentHandler{service}
}

func (h *StudentHandler) GetAll(w http.ResponseWriter, r *http.Request) []domain.Student {
	students, err := h.service.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return nil
	}

	return students
}
