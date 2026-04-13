package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/newrelic/go-agent/v3/integrations/nrgin"
	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/soolame/student-mgmt-be/internal/config"
	database "github.com/soolame/student-mgmt-be/internal/database/gorm"
	"github.com/soolame/student-mgmt-be/internal/logger"
	"github.com/soolame/student-mgmt-be/internal/routes"
)

func main() {

	router := gin.Default()
	config := config.Load()
	logger.Init(logger.INFO)

	if config.APMConfig.ApmEnabled {
		fmt.Println("License Key", config.APMConfig.NewRelicLicenseKey)
		app, err := newrelic.NewApplication(
			newrelic.ConfigAppName(config.APMConfig.NewRelicAppName),
			newrelic.ConfigLicense(config.APMConfig.NewRelicLicenseKey),
			newrelic.ConfigAppLogForwardingEnabled(true),
		)
		if err != nil {
			panic(err)
		}

		router.Use(nrgin.Middleware(app))
	}

	_, err := database.Load(config)
	if err != nil {
		logger.Error("failed to load db ")
		log.Fatal("failed to load db")
	}
	logger.Info("Starting the server [ENV]", config.Environment)
	routes.SetUpRoutes(*router)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", config.Port),
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Fatalf("listen %s\n", err)
		}
	}()

	<-quit
	log.Println("Shutting Down Gracefully..")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	srv.Shutdown(ctx)

}

func init() {
	time.Local = time.UTC
}
