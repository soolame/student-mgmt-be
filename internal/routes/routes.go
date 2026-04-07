package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/soolame/student-mgmt-be/internal/handlers"
)

func SetUpRoutes(router gin.Engine) {

	miscHandler := handlers.NewMiscHandler()
	router.GET("/healthcheck", miscHandler.HealthCheck)

	_ = router.Group("/v1")

}
