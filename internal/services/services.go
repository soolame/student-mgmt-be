package services

import (
	"errors"
	"fmt"

	"github.com/soolame/student-mgmt-be/internal/config"
	"github.com/soolame/student-mgmt-be/internal/dto"
	"github.com/soolame/student-mgmt-be/internal/logger"
	"github.com/soolame/student-mgmt-be/internal/models"
	"github.com/soolame/student-mgmt-be/internal/repositories"
	"github.com/soolame/student-mgmt-be/internal/utils"
	"gorm.io/gorm"
)

type Service struct {
	repo   repositories.Repository
	config *config.Config
}

func NewServices(cfg *config.Config) *Service {
	return &Service{
		repo:   *repositories.NewRepository(""),
		config: cfg,
	}
}

func (s *Service) AdminLogin(req dto.AdminLogin) (string, error) {

	admin, err := s.repo.GetAdminByEmail(req.Email)
	if err != nil {
		logger.Error("failed to get user with email", req.Email)
		return "", errors.New("invalid email or password")
	}

	err = utils.CheckPassword(req.Password, admin.Password)
	if err != nil {
		logger.Error("failed to check password")
		return "", errors.New("invalid email or password")
	}

	token, err := utils.GenerateJWT(admin.ID, admin.Email, *s.config)
	if err != nil {
		logger.Error("failed to generate token")
		return "", err
	}

	return token, nil

}

func (s *Service) GetAllStudents() ([]models.Student, error) {
	students, err := s.repo.GetAllStudents()
	if err != nil {
		logger.Error("failed to fetch all students error", err.Error())
		return nil, fmt.Errorf("failed to fetch all students")
	}
	return students, nil

}

func (s *Service) GetStudentDetails(id int) (*dto.StudentWithRank, error) {
	var result dto.StudentWithRank
	student, err := s.repo.GetStudentDetails(id)
	if err != nil {
		logger.Error("failed to fetch student details", err.Error())
		return nil, fmt.Errorf("failed to fetch student details %s", err.Error())
	}
	result.Student = student

	latestRank, Rerr := s.repo.GetLatestRankOfStudent(student.ID)
	if Rerr != nil && errors.Is(Rerr, gorm.ErrRecordNotFound) {
		logger.Error("failed to fetch latest rank", Rerr.Error())
		return &result, nil

	} else {
		result.LatestRank = &latestRank
	}

	return &result, nil
}
