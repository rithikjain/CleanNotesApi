package user

import "context"

type Repository interface {
	FindByID(ctx context.Context, id uint) (*User, error)

	FindByEmail(ctx context.Context, email string) (*User, error)

	Register(ctx context.Context, user *User) (*User, error)

	DoesEmailExist(ctx context.Context, email string) (bool, error)
}
