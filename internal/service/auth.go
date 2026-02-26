package service

import (
	"crypto/rand"
	"encoding/hex"

	"desent/internal/domain"
	"desent/internal/port"
)

type AuthService struct {
	repo port.TokenRepository
}

func NewAuthService(repo port.TokenRepository) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) GenerateToken(username, password string) (string, error) {
	if username == "" || password == "" {
		return "", domain.ErrCredentials
	}
	bytes := make([]byte, 32)
	rand.Read(bytes)
	token := hex.EncodeToString(bytes)
	s.repo.Store(token)
	return token, nil
}

func (s *AuthService) ValidateToken(token string) error {
	if !s.repo.Exists(token) {
		return domain.ErrUnauthorized
	}
	return nil
}
