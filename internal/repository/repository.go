package repository

import (
	"context"
	"user-rest-api/internal/domain"
)

type Users interface {
	Create(ctx context.Context, dto domain.CreateUserDTO) (string, error)
	FindOne(ctx context.Context, uuid string) (domain.User, error)
	FindAll(ctx context.Context) ([]domain.User, error)
	FindByEmail(ctx context.Context, email string) (domain.User, error)
	FindByPhone(ctx context.Context, email string) (domain.User, error)
	FindByFirstName(ctx context.Context, email string) (domain.User, error)
	FindByLastName(ctx context.Context, email string) (domain.User, error)
	Update(ctx context.Context, dto domain.UpdateUserDTO) error
	Delete(ctx context.Context, uuid string) error
}
