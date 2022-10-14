package services

import (
	"github.com/Ckala62rus/go/domain"
	"github.com/Ckala62rus/go/internal/repositories"
)

type Users interface {
	CreateUser(user domain.User) (domain.User, error)
	GetUserByName(name string) (domain.User, error)
	GetById(id int) (domain.User, error)
	GetAllUsers() []domain.User
	DeleteUserById(id int) (bool, error)
	UpdateUser(userRequest domain.User) (domain.User, error)
}

type Service struct {
	Users
}

func NewService(repo *repositories.Repository) *Service {
	return &Service{Users: NewUserService(repo.Users)}
}
