package repositories

import (
	"errors"
	"fmt"

	"github.com/Ckala62rus/go/domain"
	"gorm.io/gorm"
)

type UserMysql struct {
	db *gorm.DB
}

func NewUsersMysql(db *gorm.DB) *UserMysql {
	return &UserMysql{db: db}
}

func (u *UserMysql) GetUserByName(name string) (domain.User, error) {
	user := domain.User{}
	u.db.Debug().Where("name = ?", name).First(&user)
	if user.Id == 0 {
		mistake := fmt.Sprintf("user with name:%s not found", name)
		return user, errors.New(mistake)
	}
	return user, nil
}

func (u *UserMysql) GetById(id int) (domain.User, error) {
	var user domain.User
	u.db.First(&user, id)
	if user.Id == 0 {
		mistake := fmt.Sprintf("user with id:%d not found", id)
		return user, errors.New(mistake)
	}
	return user, nil
}

func (u *UserMysql) GetAllUsers() []domain.User {
	var users []domain.User
	u.db.Debug().Order("id desc").Find(&users)
	return users
}

func (u *UserMysql) CreateUser(user domain.User) (domain.User, error) {
	result := u.db.Create(&user)
	return user, result.Error
}

func (u *UserMysql) DeleteUserById(id int) (bool, error) {
	res := u.db.Delete(&domain.User{}, id)
	intDelete := res.RowsAffected
	err := res.Error

	if err != nil || intDelete == 0 {
		mistake := fmt.Sprintf("can't delete user with id:%d", id)
		return false, errors.New(mistake)
	}

	return true, nil
}

func (u *UserMysql) UpdateUser(userRequest domain.User) (domain.User, error) {
	var user domain.User
	u.db.Debug().First(&user, userRequest.Id)
	if user.Id == 0 {
		mistake := fmt.Sprintf("user not found with id:%d", userRequest.Id)
		return user, errors.New(mistake)
	}

	if user.Name != userRequest.Name && len(userRequest.Name) > 0 {
		user.Name = userRequest.Name
	}

	if user.Age != userRequest.Age && userRequest.Age != 0 {
		user.Age = userRequest.Age
	}

	u.db.Save(user)
	return user, nil
}
