package admin

import (
	e "ukk-smkn2/entities"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func GetUser(db *gorm.DB, id string) (*e.Admin, error) {
	var admin e.Admin
	result := db.Raw("SELECT * FROM admins WHERE id = ? LIMIT 1", id).Scan(&admin)
	if result.Error != nil {
		return nil, result.Error
	}
	return &admin, nil
}

func CreateAdmin(db *gorm.DB, admin *e.Admin) error {
	return db.Create(admin).Error
}
