package repositories

import (
	"github.com/Ckala62rus/go/domain"
	"gorm.io/gorm"
)

type ImageGorm struct {
	db *gorm.DB
}

func NewImageGorm(db *gorm.DB) *ImageGorm {
	return &ImageGorm{db: db}
}

func (u *ImageGorm) SaveImage(image domain.Image) (domain.Image, error) {
	result := u.db.Create(&image)
	return image, result.Error
}
