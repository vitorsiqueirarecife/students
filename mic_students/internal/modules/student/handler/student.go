package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/vitorsiqueirarecife/students/internal/modules/student/domain"
	"github.com/vitorsiqueirarecife/students/internal/modules/student/service"
)

type StudentHandler struct {
	service service.StudentService
}

func NewStudentHandler(service service.StudentService) *StudentHandler {
	return &StudentHandler{service}
}

func (h *StudentHandler) GetAll(w http.ResponseWriter, r *http.Request) *domain.GetAllStudentsResponse {

	page := r.URL.Query().Get("page")
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		http.Error(w, "Invalid page number", http.StatusBadRequest)
		return nil
	}

	limit := r.URL.Query().Get("limit")
	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		http.Error(w, "Invalid limit number", http.StatusBadRequest)
		return nil
	}

	res, err := h.service.GetAll(pageInt, limitInt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return nil
	}

	return res
}

func (h *StudentHandler) Create(w http.ResponseWriter, r *http.Request) {
	var student domain.Student
	if err := json.NewDecoder(r.Body).Decode(&student); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	fmt.Println(student)

	w.WriteHeader(http.StatusCreated)
}
