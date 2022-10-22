package services

import (
	"github.com/Ckala62rus/go/domain"
	"github.com/Ckala62rus/go/pkg/repositories"
)

type UserService struct {
	repo repositories.Users
}

func NewUserService(repo repositories.Users) *UserService {
	return &UserService{repo: repo}
}

func (u *UserService) GetUserByName(name string) (domain.User, error) {
	return u.repo.GetUserByName(name)
}

func (u *UserService) GetById(id int) (domain.User, error) {
	return u.repo.GetById(id)
}

func (u *UserService) GetAllUsers() []domain.User {
	return u.repo.GetAllUsers()
}

func (u *UserService) CreateUser(user domain.User) (domain.User, error) {
	return u.repo.CreateUser(user)
}

func (u *UserService) DeleteUserById(id int) (bool, error) {
	return u.repo.DeleteUserById(id)
}

func (u *UserService) UpdateUser(userRequest domain.User) (domain.User, error) {
	return u.repo.UpdateUser(userRequest)
}
