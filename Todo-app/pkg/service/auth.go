package service

import (
	"crypto/sha1"
	"fmt"
	todo "github.com/Leo-tumo/learngo/Todo-app"
	"github.com/Leo-tumo/learngo/Todo-app/pkg/repository"
	"github.com/dgrijalva/jwt-go"
	"github.com/fatih/color"
	"time"
)

const (
	salt = "hdjakosoasdhfhiedhehrtjhdflgp"

	signingKey = "hjksdklgaiudgaisdgb"
	tokenTTL   = 12 * time.Hour
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	color.Green("\t\t NEW AUTH SERVICE // from service")
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user todo.User) (int, error) {
	color.Green("\t\t CREATE USER // from service")
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func (s *AuthService) GenerateToken(username, password string) (string, error) {
	color.Green("\t\t GENERATE TOKEN // from service")
	user, err := s.repo.GetUser(username, generatePasswordHash(password))
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

func generatePasswordHash(password string) string {
	color.Green("\t\t GENERATE PASSWORD HASH // from service")
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
