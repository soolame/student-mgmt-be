package dto

import "github.com/soolame/student-mgmt-be/internal/models"

type AdminLogin struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type StudentWithRank struct {
	Student    models.Student      `json:"student"`
	LatestRank *models.RankHistory `json:"latest_rank"`
}
