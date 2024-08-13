package postgres

import (
	"context"
	"user-rest-api/internal/domain"
	"user-rest-api/internal/repository"
	"user-rest-api/pkg/dbclient"
	"user-rest-api/pkg/logger"
)

type usersRepo struct {
	client dbclient.Client
	logger logger.Logger
}

func (u *usersRepo) Create(ctx context.Context, dto domain.CreateUserDTO) (string, error) {
	panic("unimplemented")
}

func (u *usersRepo) Delete(ctx context.Context, uuid string) error {
	panic("unimplemented")
}

func (u *usersRepo) FindAll(ctx context.Context) ([]domain.User, error) {
	panic("unimplemented")
}

func (u *usersRepo) FindByEmail(ctx context.Context, email string) (domain.User, error) {
	panic("unimplemented")
}

func (u *usersRepo) FindByFirstName(ctx context.Context, email string) (domain.User, error) {
	panic("unimplemented")
}

func (u *usersRepo) FindByLastName(ctx context.Context, email string) (domain.User, error) {
	panic("unimplemented")
}

func (u *usersRepo) FindByPhone(ctx context.Context, email string) (domain.User, error) {
	panic("unimplemented")
}

func (u *usersRepo) FindOne(ctx context.Context, uuid string) (domain.User, error) {
	panic("unimplemented")
}

func (u *usersRepo) Update(ctx context.Context, dto domain.UpdateUserDTO) error {
	panic("unimplemented")
}

func NewUsersRepo(client dbclient.Client, logger logger.Logger) repository.Users {
	return &usersRepo{
		client: client,
		logger: logger,
	}
}
