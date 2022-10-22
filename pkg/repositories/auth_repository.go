package repositories

import (
	"github.com/Ckala62rus/go/domain"
	"gorm.io/gorm"
)

type AuthGorm struct {
	db *gorm.DB
}

func NewAuthGorm(db *gorm.DB) *AuthGorm {
	return &AuthGorm{db: db}
}

func (a *AuthGorm) CreateUser(user domain.User) (int, error) {
	result := a.db.Create(&user)
	if user.Id == 0 || result.Error != nil {
		return 0, result.Error
	}

	return user.Id, nil
}

func (a *AuthGorm) GetUser(email, password string) (domain.User, error) {
	var user domain.User
	a.db.Where(map[string]interface{}{"email": email, "password": password}).Find(&user)
	return user, nil
}
