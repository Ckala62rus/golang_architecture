package services

import (
	"github.com/Ckala62rus/go/domain"
	"github.com/Ckala62rus/go/pkg/repositories"
)

type UserImageService struct {
	repo repositories.UserImage
}

func NewUserImageService(repo repositories.UserImage) *UserImageService {
	return &UserImageService{repo: repo}
}

func (img *UserImageService) SaveImage(image domain.Image) (domain.Image, error) {
	image, err := img.repo.SaveImage(image)
	if err != nil {
		return image, err
	}
	return domain.Image{}, nil
}
