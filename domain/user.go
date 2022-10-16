package domain

import (
	"fmt"
	"time"
)

type User struct {
	Id        int    `gorm:"not null;uniqueIndex;primary_key"`
	Name      string `gorm:"size:100;not null"`
	Age       int    `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u *User) GetInfo() {
	fmt.Printf("Username : %s, Age: %d", u.Name, u.Age)
}
