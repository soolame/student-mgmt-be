package repositories

import (
	"fmt"

	database "github.com/soolame/student-mgmt-be/internal/database/gorm"
	"github.com/soolame/student-mgmt-be/internal/models"
)

type Repository struct {
}

func NewRepository(db string) *Repository {

	return &Repository{}

}
func (r *Repository) GetAdminByEmail(email string) (*models.Admin, error) {
	var admin models.Admin
	err := database.GetAppDB().Where("email = ?", email).First(&admin).Error
	if err != nil {
		return nil, err
	}
	return &admin, nil
}

func (r *Repository) CreateAdmin(email, password string) (*models.Admin, error) {
	admin := models.Admin{
		Email:    email,
		Password: password,
	}
	err := database.GetAppDB().Create(&admin).Error
	if err != nil {
		return nil, err
	}
	return &admin, nil
}

