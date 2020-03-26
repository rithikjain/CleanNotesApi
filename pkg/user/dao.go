package user

import (
	"context"
	"github.com/jinzhu/gorm"
	"github.com/rithikjain/CleanNotesApi/pkg"
)

type repo struct {
	DB *gorm.DB
}

func NewRepo(db *gorm.DB) Repository {
	return &repo{
		DB: db,
	}
}

func (r *repo) FindByID(ctx context.Context, id uint) (*User, error) {
	user := &User{}
	r.DB.Where("id = ?", id).First(user)
	if user.Email == "" {
		return nil, pkg.ErrNotFound
	}
	return user, nil
}

func (r *repo) Register(ctx context.Context, user *User) (*User, error) {
	result := r.DB.Create(user)
	if result.Error != nil {
		return nil, pkg.ErrDatabase
	}
	return user, nil
}

func (r *repo) DoesEmailExist(ctx context.Context, email string) (bool, error) {
	user := &User{}
	if r.DB.Where("email = ?", email).First(user).RecordNotFound() {
		return false, nil
	}
	return true, nil
}

func (r *repo) FindByEmail(ctx context.Context, email string) (*User, error) {
	user := &User{}
	result := r.DB.Where("email = ?", email).First(user)

	if result.Error == gorm.ErrRecordNotFound {
		return nil, pkg.ErrNotFound
	}
	return user, nil
}
