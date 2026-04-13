package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/soolame/student-mgmt-be/internal/config"
	database "github.com/soolame/student-mgmt-be/internal/database/gorm"
	"github.com/soolame/student-mgmt-be/internal/logger"
	"github.com/soolame/student-mgmt-be/internal/repositories"
	"github.com/soolame/student-mgmt-be/internal/utils"
)

func main() {

	fmt.Println("Create Admin User")
	config := config.Load()
	logger.Init(logger.INFO)
	database.Load(config)
	fmt.Println("Length of Args", len(os.Args))
	if len(os.Args) < 3 || len(os.Args) > 3 {
		log.Fatal("can only expect exactly two parameter exe [email] [password]")
	}

	email := os.Args[1]
	passwordRow := os.Args[2]

	hashedPassword, err := utils.HashPassword(passwordRow)
	if err != nil {
		log.Fatal("failed to hash password")
		return

	}

	repo := repositories.NewRepository("db")
	admin, Cerr := repo.CreateAdmin(email, hashedPassword)
	if Cerr != nil {
		log.Fatal("Failed to create admin")
		return

	}
	fmt.Printf("\nCreate User with Email: %s ", admin.Email)

}

func init() {
	time.Local = time.UTC
}
