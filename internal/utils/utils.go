package utils

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/soolame/student-mgmt-be/internal/config"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckPassword(password, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}

func GenerateJWT(adminID uint, email string, cfg config.Config) (string, error) {
	claims := jwt.MapClaims{
		"admin_id": adminID,
		"email":    email,
		"exp":      time.Now().Add(5 * time.Minute).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(cfg.Secret))
}
