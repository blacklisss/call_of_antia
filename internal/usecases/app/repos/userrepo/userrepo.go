package userrepo

import (
	"antia/internal/entities/userentity"
	"context"
	"database/sql"
	"fmt"
)

type CreateUserParams struct {
	ID   uint64         `json:"id"`
	Name sql.NullString `json:"name"`
}

type UserStore interface {
	CreateUser(ctx context.Context, args *CreateUserParams) (*userentity.User, error)
	GetUserByID(ctx context.Context, id uint64) (*userentity.User, error)
	DeleteUser(ctx context.Context, id uint64) error
}

type Users struct {
	ustore UserStore
}

func NewUsers(ustore UserStore) *Users {
	return &Users{
		ustore,
	}
}

func (us *Users) CreateUser(ctx context.Context, args *CreateUserParams) (*userentity.User, error) {
	user, err := us.ustore.CreateUser(ctx, args)
	if err != nil {
		return nil, fmt.Errorf("create user error: %w", err)
	}

	return user, nil
}

func (us *Users) GetUserByID(ctx context.Context, id uint64) (*userentity.User, error) {
	user, err := us.ustore.GetUserByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("create user error: %w", err)
	}

	return user, nil
}

func (us *Users) DeleteUser(ctx context.Context, id uint64) (*userentity.User, error) {
	user, err := us.ustore.GetUserByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("delete country error: %w", err)
	}
	return user, us.ustore.DeleteUser(ctx, id)
}
