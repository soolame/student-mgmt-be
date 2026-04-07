package database

import (
	"fmt"

	"github.com/soolame/student-mgmt-be/internal/config"
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

func LoadPostgresDB(host, user, pwd, port, dbname string, enforceSsl bool) *gorm.DB {
	sslMode := "disable"
	if enforceSsl {
		sslMode = "require"
	}
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=Asia/Shanghai", host, user, pwd, dbname, port, sslMode)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil
	}
	return db

}

func Load(config *config.Config) *DbStore {
	enforceSSL := false
	if config.Environment == "production" {
		enforceSSL = true
	}
	db := LoadPostgresDB(config.DBConfig.Host, config.DBConfig.User, config.DBConfig.Password, fmt.Sprint(config.DBConfig.Port), config.DBConfig.Name, enforceSSL)

	dbStore := DbStore{appdb: db}
	DB = dbStore
	return &DbStore{
		appdb: db,
	}
}
