package service

import (
	"crypto/sha1"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	restJournal "rest_journal"
	"rest_journal/pkg/repository"
	"time"
)

const (
	salt       = "2jhgubvyqwk"
	signingKey = "12-i0j9hgouervwyb"
	tokenTTL   = time.Hour * 12
)

type tokenClaims struct {
	jwt.RegisteredClaims
	UserId   int    `json:"user_id"`
	UserRole string `json:"user_role"`
}

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user restJournal.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func (s *AuthService) GenerateToken(email, password string) (string, error) {
	user, err := s.repo.GetUser(email, generatePasswordHash(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaims{
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(tokenTTL)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		user.Id,
		user.Role,
	})

	return token.SignedString([]byte(signingKey))
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
