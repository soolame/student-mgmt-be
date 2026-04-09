package handlers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/soolame/student-mgmt-be/internal/config"
	"github.com/soolame/student-mgmt-be/internal/dto"
	"github.com/soolame/student-mgmt-be/internal/logger"
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

func (h *StudentHandler) GetStudentDetails(ctx *gin.Context) {
	id := ctx.Param("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		logger.Error("invalid student id", err)
		ctx.JSON(400, gin.H{"error": "invalid student id "})
		return
	}
	studentDetails, err := h.service.GetStudentDetails(intID)
	if err != nil {
		logger.Error("failed to get student details", err)
		ctx.JSON(404, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, studentDetails)
}

func (h *StudentHandler) CreateStudent(ctx *gin.Context) {
	req := dto.CreateStudent{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		logger.Error("failed to bind the body", err.Error())
		ctx.JSON(400, gin.H{"message": "Invalid request body", "error": err.Error()})
		return
	}

	student, err := h.service.CreateStudent(req)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(201, gin.H{"student": student})

}
