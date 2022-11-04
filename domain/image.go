package domain

import "time"

type Image struct {
	Id        int    `gorm:"not null;uniqueIndex;primary_key"`
	Filename  string `gorm:"size:255"`
	Path      string `gorm:"not null"`
	UserId    int
	CreatedAt time.Time
	UpdatedAt time.Time
}
