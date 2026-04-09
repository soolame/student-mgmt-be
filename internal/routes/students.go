package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/soolame/student-mgmt-be/internal/config"
	"github.com/soolame/student-mgmt-be/internal/handlers"
	"github.com/soolame/student-mgmt-be/internal/middleware"
)

func SetUpStudentRoutes(v1router *gin.RouterGroup) {
	students := v1router.Group("/students")
	handler := handlers.NewStudentHandler(*config.GetConfig())

	students.GET("/", handler.GetAllStudent)

}
