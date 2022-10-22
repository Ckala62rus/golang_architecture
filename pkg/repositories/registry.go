package repositories

import "github.com/Ckala62rus/go/domain"

type Model struct {
	Model interface{}
}

func RegisterModel() []Model {
	return []Model{
		{Model: domain.User{}},
	}
}
