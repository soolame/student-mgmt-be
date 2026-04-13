package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/soolame/student-mgmt-be/internal/config"
	"github.com/soolame/student-mgmt-be/internal/database"
	gorm "github.com/soolame/student-mgmt-be/internal/database/gorm"
	"github.com/soolame/student-mgmt-be/internal/logger"
)

func main() {
	logger.Init(logger.INFO)
	cfg := config.Load()

	_, err := gorm.Load(cfg)
	if err != nil {
		log.Fatal("failed to load db:", err)
	}

	if len(os.Args) < 2 {
		log.Fatal("usage: migrate [up|down|generate]")
	}

	cmd := os.Args[1]
	fmt.Println("CMD", cmd)

	switch cmd {
	case "up":
		err = database.RunMigrations(gorm.GetAppDB(), cfg.MigrationPath)

	case "down":
		err = database.RollbackLast(gorm.GetAppDB(), cfg.MigrationPath)

	case "generate-schema":
		err = database.GenerateSchema(cfg)

	default:
		log.Fatalf("unknown command: %s", cmd)
	}

	if err != nil {
		log.Fatal(err)
	}
}

func init() {
	time.Local = time.UTC
}
