package services

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/soolame/student-mgmt-be/internal/config"
	database "github.com/soolame/student-mgmt-be/internal/database/gorm"
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

func (s *Service) CreateStudent(req dto.CreateStudent) (*models.Student, error) {
	student, err := s.repo.CreateStudent(&req)
	if err != nil {
		logger.Error("failed to create student error", err)
		return nil, fmt.Errorf("failed to create student")
	}
	return student, nil
}

func (s *Service) UpdateStudent(id int, req dto.UpdateStudent) (models.Student, error) {
	tx := database.GetAppDB().Begin()

	student, err := s.repo.GetStudentDetails(id)
	if err != nil {
		tx.Rollback()
		return models.Student{}, err
	}

	updates := make(map[string]interface{})

	if req.FirstName != nil {
		updates["first_name"] = *req.FirstName
	}
	if req.MiddleName != nil {
		updates["middle_name"] = *req.MiddleName
	}
	if req.LastName != nil {
		updates["last_name"] = *req.LastName
	}
	if req.Email != nil {
		updates["email"] = *req.Email
	}
	if req.Class != nil {
		updates["class"] = strconv.FormatInt(*req.Class, 10)
	}
	if req.Address != nil {
		updates["address"] = *req.Address
	}
	if req.Gender != nil {
		updates["gender"] = *req.Gender
	}
	if req.Phone != nil {
		updates["phone"] = *req.Phone
	}
	if req.DOB != nil {
		updates["dob"] = *req.DOB
	}
	if req.GuardianContact != nil {
		updates["guardian_contact"] = *req.GuardianContact
	}
	if req.GuardianName != nil {
		updates["guardian_name"] = *req.GuardianName
	}
	if req.GuardianRelation != nil {
		updates["guardian_relation"] = *req.GuardianRelation
	}

	if len(updates) > 0 {
		if err := s.repo.UpdateStudent(&student, updates); err != nil {
			tx.Rollback()
			return models.Student{}, err
		}
	}

	if req.Rank != nil {
		rank := models.RankHistory{
			StudentID:     student.ID,
			Rank:          req.Rank.Rank,
			MarksAttained: req.Rank.MarksAttained,
			Grade:         req.Rank.Grade,
			Term:          req.Rank.Term,
			Year:          req.Rank.Year,
		}

		if err := s.repo.CreateRankHistory(&rank); err != nil {
			tx.Rollback()
			return models.Student{}, err
		}
	}

	if err := tx.Commit().Error; err != nil {
		return models.Student{}, err
	}

	return student, nil
}
