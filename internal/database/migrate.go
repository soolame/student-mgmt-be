package database

import (
	"errors"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/soolame/student-mgmt-be/internal/logger"
	"gorm.io/gorm"
)

func newMigrate(gormDB *gorm.DB, path string) (*migrate.Migrate, error) {
	sqlDB, err := gormDB.DB()
	if err != nil {
		return nil, err
	}

	driver, err := postgres.WithInstance(sqlDB, &postgres.Config{})
	if err != nil {
		return nil, err
	}

	return migrate.NewWithDatabaseInstance(
		"file://"+path,
		"postgres",
		driver,
	)
}

func RunMigrations(gormDB *gorm.DB, path string) error {
	m, err := newMigrate(gormDB, path)
	if err != nil {
		return err
	}

	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return err
	}

	logger.Info("Migrations Applied")
	return nil
}

func RollbackLast(gormDB *gorm.DB, path string) error {
	m, err := newMigrate(gormDB, path)
	if err != nil {
		return err
	}

	if err := m.Steps(-1); err != nil {
		return err
	}

	logger.Info(" rolled back 1 migration")
	return nil
}
