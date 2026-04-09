package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/soolame/student-mgmt-be/internal/config"
	"github.com/soolame/student-mgmt-be/internal/dto"
	"github.com/soolame/student-mgmt-be/internal/logger"
	"github.com/soolame/student-mgmt-be/internal/services"
)

type AdminHandler struct {
	service services.Service
}

func NewAdminHandler(config config.Config) AdminHandler {
	return AdminHandler{
		service: *services.NewServices(&config),
	}
}

func (h *AdminHandler) Login(ctx *gin.Context) {
	req := dto.AdminLogin{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		logger.Error("failed to bind the body ,validation failed", err.Error())
		ctx.JSON(400, gin.H{"error": "failed to validate either email or password is invalid or missing"})
		return
	}

	token, err := h.service.AdminLogin(req)
	if err != nil {
		logger.Error("failed to get token", err.Error())
		ctx.JSON(401, gin.H{"error": "user doesnt exist"})
		return
	}

	ctx.JSON(200, gin.H{"token": token})

}
