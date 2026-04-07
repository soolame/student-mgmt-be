package handlers

import (
	"time"

	"github.com/gin-gonic/gin"
)

type MiscHandler struct{}

func NewMiscHandler() MiscHandler {
	return MiscHandler{}
}

func (h *MiscHandler) HealthCheck(ctx *gin.Context) {
	ctx.JSON(200,
		gin.H{"status": "ok", "time": time.Now()},
	)

}
