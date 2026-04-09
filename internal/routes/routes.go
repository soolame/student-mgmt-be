package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/soolame/student-mgmt-be/internal/handlers"
)

func SetUpRoutes(router gin.Engine) {

	miscHandler := handlers.NewMiscHandler()
	router.GET("/healthcheck", miscHandler.HealthCheck)

	v1 := router.Group("/v1")
	SetUpStudentRoutes(v1)
	SetUpAdminRoutes(v1)

}
