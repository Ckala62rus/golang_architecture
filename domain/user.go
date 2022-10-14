package domain

import (
	"fmt"
)

type User struct {
	Id int
	Name string
	Age  int
}

func (u *User) GetInfo() {
	fmt.Printf("Username : %s, Age: %d", u.Name, u.Age)
}