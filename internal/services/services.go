package services

import "github.com/soolame/student-mgmt-be/internal/repositories"

type Service struct {
}

func NewServices(repo repositories.Repository) *Service {
	return &Service{}

}
