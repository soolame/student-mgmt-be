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
	"github.com/soolame/student-mgmt-be/internal/config"
	"github.com/soolame/student-mgmt-be/internal/routes"
)

func main() {

	router := gin.Default()
	config := config.Load()
	log.Println("Starting server in Environemnt", config.Environment)
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
