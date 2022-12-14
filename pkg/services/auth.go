package services

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"time"

	"github.com/Ckala62rus/go/domain"
	"github.com/Ckala62rus/go/pkg/repositories"
	"github.com/dgrijalva/jwt-go"
)

const (
	salt       = "qweasdzxc" // хорошая практика добавлять соль к паролю
	signingKey = "qweasadasfasfasfsfafasf"
	// tokenTTL   = 12 * time.Hour
	tokenTTL = 10 * time.Minute
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

type AuthService struct {
	repo repositories.Authorization
}

func NewAuthService(repo repositories.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user domain.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func (s *AuthService) GenerateToken(email string, password string) (string, error) {
	user, err := s.repo.GetUser(email, generatePasswordHash(password))

	if user.Id == 0 {
		return "User not found", err
	}

	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
	})

	return token.SignedString([]byte(signingKey))
}

func (s *AuthService) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(signingKey), nil
	})
	if err != nil {
		return 0, nil
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of type *tokenClaims")
	}

	return claims.UserId, nil
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
