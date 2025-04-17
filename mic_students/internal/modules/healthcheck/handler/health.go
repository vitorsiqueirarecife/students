package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vitorsiqueirarecife/students/internal/modules/healthcheck/domain"
)

type HealthHandler struct {
	service domain.HealthService
}

func NewHealthHandler(service domain.HealthService) *HealthHandler {
	return &HealthHandler{service}
}

func (h *HealthHandler) Check(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": h.service.Check()})
}
