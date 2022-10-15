package repositories

import (
	"github.com/Ckala62rus/go/domain"
	"gorm.io/gorm"
)

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
}

func NewUserRepository(db *gorm.DB) *Repository {
	return &Repository{Users: NewUsersMysql(db)}
}
