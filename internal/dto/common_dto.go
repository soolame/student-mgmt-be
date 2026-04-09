package dto

import (
	"fmt"
	"net/mail"
	"strings"

	"github.com/soolame/student-mgmt-be/internal/models"
)

type AdminLogin struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type StudentWithRank struct {
	Student    models.Student      `json:"student"`
	LatestRank *models.RankHistory `json:"latest_rank"`
}

type CreateStudent struct {
	FirstName  string `json:"first_name" binding:"required,max=100"`
	MiddleName string `json:"middle_name,omitempty" binding:"max=100"`
	LastName   string `json:"last_name" binding:"required,max=100"`

	Email string `json:"email" binding:"required,email,max=150"`

	GuardianName     string `json:"guardian_name,omitempty" binding:"max=150"`
	GuardianRelation string `json:"guardian_relation,omitempty" binding:"max=50"`
	GuardianContact  string `json:"guardian_contact,omitempty" binding:"required,numeric,min=10,max=15"`

	Class   int64          `json:"class" binding:"required,min=9,max=12"`
	Address models.Address `json:"address" binding:"required"`

	Gender string `json:"gender" binding:"required,oneof=male female other"`
	DOB    string `json:"dob" binding:"required,datetime=2006-01-02"`
	Phone  string `json:"phone" binding:"required,numeric,min=10,max=15"`
}

func (c *CreateStudent) Validate() error {
	var errorList []string

	if c.Class < 9 || c.Class > 12 {
		errorList = append(errorList, "Class should be between 9 and 12")
	}

	if _, err := mail.ParseAddress(c.Email); err != nil {
		errorList = append(errorList, "Invalid email format")
	}

	if len(errorList) > 0 {
		return fmt.Errorf(strings.Join(errorList, ", "))
	}

	return nil
}
