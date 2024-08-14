package repository

import (
	"context"
	"user-rest-api/internal/domain"
)

type Users interface {
	Create(ctx context.Context, user domain.User) (string, error)
	FindOne(ctx context.Context, uuid string) (domain.User, error)
	FindAll(ctx context.Context) ([]domain.User, error)
	FindByEmail(ctx context.Context, email string) (domain.User, error)
	FindByPhone(ctx context.Context, phone string) (domain.User, error)
	Update(ctx context.Context, uuid string, user domain.User) error
	Delete(ctx context.Context, uuid string) error
}
