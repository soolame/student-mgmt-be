package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/soolame/student-mgmt-be/internal/config"
	"github.com/soolame/student-mgmt-be/internal/services"
)

type StudentHandler struct {
	service services.Service
}

func NewStudentHandler(config config.Config) StudentHandler {
	return StudentHandler{
		service: *services.NewServices(&config),
	}
}

func (h *StudentHandler) GetAllStudent(ctx *gin.Context) {

	students, err := h.service.GetAllStudents()
	if err != nil {
		ctx.JSON(400, gin.H{"error": err})
		return
	}
	ctx.JSON(200, gin.H{"students": students})
}

