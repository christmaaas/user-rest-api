package service

import (
	"context"
	"errors"
	"fmt"
	"user-rest-api/internal/apperror"
	"user-rest-api/internal/domain"
	"user-rest-api/internal/repository"
	"user-rest-api/pkg/logger"

	"golang.org/x/crypto/bcrypt"
)

type UsersService struct {
	logger     logger.Logger
	repository repository.Users
}

func NewUsersService(logger logger.Logger, repository repository.Users) *UsersService {
	return &UsersService{
		logger:     logger,
		repository: repository,
	}
}

func generatePasswordHash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return "", fmt.Errorf("failed to hash password due to error %w", err)
	}

	return string(hash), nil
}

func (s *UsersService) CreateUser(ctx context.Context, dto domain.CreateUserDTO) (string, error) {
	s.logger.Debug("check password and repeat password")
	if dto.Password != dto.RepeatPassword {
		return "", apperror.ErrBadRequest
	}

	user := domain.NewUser(dto)

	s.logger.Debug("generate password hash")
	pwd, err := generatePasswordHash(user.Password)
	if err != nil {
		return "", fmt.Errorf("failed to generate hash due to error: %s", err)
	}
	user.Password = pwd

	userUUID, err := s.repository.Create(ctx, user)

	if err != nil {
		return userUUID, fmt.Errorf("failed to create user. error: %w", err)
	}

	return userUUID, nil
}

func (s *UsersService) GetOneUser(ctx context.Context, uuid string) (domain.User, error) {
	user, err := s.repository.FindOne(ctx, uuid)

	if err != nil {
		if errors.Is(err, apperror.ErrNotFound) {
			return user, err
		}
		return user, fmt.Errorf("failed to find user by uuid. error: %w", err)
	}
	return user, nil
}

func (s *UsersService) GetAllUsers(ctx context.Context) ([]domain.User, error) {
	users, err := s.repository.FindAll(ctx)

	if err != nil {
		if errors.Is(err, apperror.ErrNotFound) {
			return users, err
		}
		return users, fmt.Errorf("failed to find all users. error: %w", err)
	}
	return users, nil
}

func (s *UsersService) GetUserByEmail(ctx context.Context, email string) (domain.User, error) {
	user, err := s.repository.FindByEmail(ctx, email)

	if err != nil {
		if errors.Is(err, apperror.ErrNotFound) {
			return user, err
		}
		return user, fmt.Errorf("failed to find user by email. error: %w", err)
	}
	return user, nil
}

func (s *UsersService) GetUserByPhone(ctx context.Context, phone string) (domain.User, error) {
	user, err := s.repository.FindByPhone(ctx, phone)

	if err != nil {
		if errors.Is(err, apperror.ErrNotFound) {
			return user, err
		}
		return user, fmt.Errorf("failed to find user by phone. error: %w", err)
	}
	return user, nil
}

func (s *UsersService) UpdateUser(ctx context.Context, uuid string, dto domain.UpdateUserDTO) error {
	s.logger.Debug("compare old and new passwords")
	if dto.OldPassword != dto.NewPassword {
		user, err := s.GetOneUser(ctx, uuid)
		if err != nil {
			return err
		}

		s.logger.Debug("compare current and old password hash")
		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(dto.OldPassword))
		if err != nil {
			return apperror.ErrBadRequest
		}

		dto.OldPassword = dto.NewPassword
	}

	updatedUser := domain.UpdatedUser(dto)

	s.logger.Debug("generate password hash")
	pwd, err := generatePasswordHash(updatedUser.Password)
	if err != nil {
		return fmt.Errorf("failed to generate hash due to error: %s", err)
	}
	updatedUser.Password = pwd

	if err = s.repository.Update(ctx, uuid, updatedUser); err != nil {
		if errors.Is(err, apperror.ErrNotFound) {
			return err
		}
		return fmt.Errorf("failed to update user. error: %w", err)
	}
	return nil
}

func (s *UsersService) DeleteUser(ctx context.Context, uuid string) error {
	err := s.repository.Delete(ctx, uuid)

	if err != nil {
		if errors.Is(err, apperror.ErrNotFound) {
			return err
		}
		return fmt.Errorf("failed to delete user. error: %w", err)
	}
	return nil
}
