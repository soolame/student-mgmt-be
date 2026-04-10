package repositories

import (
	"errors"
	"fmt"

	database "github.com/soolame/student-mgmt-be/internal/database/gorm"
	"github.com/soolame/student-mgmt-be/internal/dto"
	"github.com/soolame/student-mgmt-be/internal/models"
	"gorm.io/gorm"
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

func (r *Repository) GetAllStudents() ([]models.Student, error) {
	var students []models.Student

	err := database.GetAppDB().Where("is_active = ?", true).Find(&students).Error
	if err != nil {
		return []models.Student{}, fmt.Errorf("failed to get students")
	}
	return students, nil
}

func (r *Repository) GetStudentDetails(id int) (models.Student, error) {
	var student models.Student

	err := database.GetAppDB().Where("id = ?", id).First(&student).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.Student{}, fmt.Errorf("no student with the id")
		}
		return models.Student{}, fmt.Errorf("failed to get student: %w", err)
	}
	return student, nil
}

func (r *Repository) GetLatestRankOfStudent(studentID uint) (models.RankHistory, error) {
	var rank models.RankHistory

	err := database.GetAppDB().
		Where("student_id = ?", studentID).
		Order("created_at DESC").
		First(&rank).Error

	if err != nil {
		return models.RankHistory{}, err
	}

	return rank, nil
}

func (r *Repository) CreateStudent(req *dto.CreateStudent) (*models.Student, error) {
	student := &models.Student{
		FirstName:        req.FirstName,
		MiddleName:       req.MiddleName,
		LastName:         req.LastName,
		Email:            req.Email,
		GuardianName:     req.GuardianName,
		GuardianRelation: req.GuardianRelation,
		GuardianContact:  req.GuardianContact,
		Class:            string(req.Class),
		Address:          req.Address,
		Phone:            req.Phone,
		DOB:              req.DOB,
		Gender:           req.Gender,
	}

	if err := database.GetAppDB().Create(student).Error; err != nil {
		return nil, err
	}

	return student, nil
}

func (r *Repository) UpdateStudent(student *models.Student, updates map[string]interface{}) error {
	return database.GetAppDB().Model(student).Updates(updates).Error
}

func (r *Repository) CreateRankHistory(rank *models.RankHistory) error {
	return database.GetAppDB().Create(rank).Error
}
