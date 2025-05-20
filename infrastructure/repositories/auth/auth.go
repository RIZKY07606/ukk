package auth

import (
	"os"
	"time"
	e "ukk-smkn2/entities"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func CheckPasswordHash(password, hash string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}

func GenerateJWT(adminID uuid.UUID) (string, error) {
	claims := jwt.MapClaims{
		"sub": adminID,
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

func FindUserByEmail(db *gorm.DB, email string) (*e.Admin, error) {
	var admin e.Admin
	result := db.Raw("SELECT * FROM admins WHERE email = ? LIMIT 1", email).Scan(&admin)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, nil
	}
	return &admin, nil
}
