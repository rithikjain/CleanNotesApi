package user

import (
	"context"
	"errors"
	"github.com/rithikjain/CleanNotesApi/pkg"
	"golang.org/x/crypto/bcrypt"
	"strings"
)

type Service interface {
	Register(ctx context.Context, user *User) (*User, error)

	Login(ctx context.Context, email, password string) (*User, error)

	GetUserByID(ctx context.Context, id uint) (*User, error)

	GetRepo() Repository
}

type service struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &service{
		repo: r,
	}
}

func (user *User) Validate() (bool, error) {
	if !strings.Contains(user.Email, "@") {
		return false, pkg.ErrEmail
	}

	if len(user.Password) < 6 || len(user.Password) > 60 {
		return false, pkg.ErrPassword
	}
	return true, nil
}

func (s *service) Register(ctx context.Context, user *User) (*User, error) {
	// Validation
	validate, err := user.Validate()
	if !validate {
		return nil, err
	}

	exists, err := s.repo.DoesEmailExist(ctx, user.Email)
	if err != nil {
		return nil, err
	}
	if exists {
		//noinspection GoErrorStringFormat
		return nil, errors.New("User already exists")
	}
	pass, err := HashPassword(user.Password)
	if err != nil {
		return nil, err
	}
	user.Password = pass
	return s.repo.Register(ctx, user)
}

func (s *service) Login(ctx context.Context, email, password string) (*User, error) {
	user := &User{}
	user, err := s.repo.FindByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	if CheckPasswordHash(password, user.Password) {
		return user, nil
	}
	return nil, pkg.ErrNotFound
}

func (s *service) GetUserByID(ctx context.Context, id uint) (*User, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *service) GetRepo() Repository {
	return s.repo
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
