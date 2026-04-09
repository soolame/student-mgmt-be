package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/soolame/student-mgmt-be/internal/config"
	"github.com/soolame/student-mgmt-be/internal/handlers"
)

func SetUpAdminRoutes(v1router *gin.RouterGroup) {

	admin := v1router.Group("/admin")
	handler := handlers.NewAdminHandler(*config.GetConfig())
	{

		admin.POST("/login", handler.Login)

	}

}
