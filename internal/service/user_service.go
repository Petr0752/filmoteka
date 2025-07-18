package service

import (
	"errors"
	"filmoteka/pkg/jwtutil"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo UserRepository
}

func NewUserService(r UserRepository) *UserService {
	return &UserService{repo: r}
}

func (s *UserService) Login(username, password string) (string, error) {
	user, err := s.repo.GetByUsername(username)
	if err != nil {
		return "", errors.New("invalid username")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", errors.New("invalid username or password")
	}

	token, err := jwtutil.GenerateToken(user.ID, user.Role)
	if err != nil {
		return "", err
	}

	return token, nil
}
