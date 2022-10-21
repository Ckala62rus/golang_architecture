package repositories

import (
	"github.com/Ckala62rus/go/domain"
	"gorm.io/gorm"
)

type Authorization interface {
	CreateUser(user domain.User) (int, error)
	GetUser(email, password string) (domain.User, error)
}

type Users interface {
	CreateUser(user domain.User) (domain.User, error)
	GetUserByName(name string) (domain.User, error)
	GetById(id int) (domain.User, error)
	GetAllUsers() []domain.User
	DeleteUserById(id int) (bool, error)
	UpdateUser(user domain.User) (domain.User, error)
}

type Repository struct {
	Users
	Authorization
}

func NewUserRepository(db *gorm.DB) *Repository {
	return &Repository{
		Users:         NewUsersMysql(db),
		Authorization: NewAuthGorm(db),
	}
}
