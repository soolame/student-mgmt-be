package database

import (
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	"github.com/soolame/student-mgmt-be/internal/config"
	"github.com/soolame/student-mgmt-be/internal/database"
	"github.com/soolame/student-mgmt-be/internal/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DbStore struct {
	appdb *gorm.DB
}

var DB DbStore

func GetAppDB() *gorm.DB {
	return DB.appdb
}

func LoadPostgresDB(host, user, pwd, port, dbname string, enforceSsl bool) (*gorm.DB, error) {
	sslMode := "disable"
	if enforceSsl {
		sslMode = "require"
	}
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=Asia/Shanghai", host, user, pwd, dbname, port, sslMode)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil

}

func Load(config *config.Config) (*DbStore, error) {
	enforceSSL := false
	if config.Environment == "production" {
		enforceSSL = true
	}

	db, err := LoadPostgresDB(config.DBConfig.Host, config.DBConfig.User, config.DBConfig.Password, fmt.Sprint(config.DBConfig.Port), config.DBConfig.Name, enforceSSL)
	logTag := "[LOAD_DB]"
	if err != nil {
		logger.Error(logTag, "Failed to load to app database error", err.Error())
		return nil, fmt.Errorf("failed to connect load app db %s", err.Error())
	}
	if config.IsEnvLocal() && config.MigrationEnabled {
		logger.Info("Migrating...")
		database.RunMigrations(db, config.MigrationPath)
	}
	dbStore := DbStore{appdb: db}
	DB = dbStore
	return &DbStore{
		appdb: db,
	}, nil
}
