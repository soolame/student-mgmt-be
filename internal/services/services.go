package services

import (
	"errors"

	"github.com/soolame/student-mgmt-be/internal/config"
	"github.com/soolame/student-mgmt-be/internal/dto"
	"github.com/soolame/student-mgmt-be/internal/logger"
	"github.com/soolame/student-mgmt-be/internal/repositories"
	"github.com/soolame/student-mgmt-be/internal/utils"
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
