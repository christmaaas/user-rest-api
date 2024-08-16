package postgres

import (
	"context"
	"errors"
	"fmt"
	"user-rest-api/internal/apperror"
	"user-rest-api/internal/domain"
	"user-rest-api/internal/repository"
	"user-rest-api/pkg/dbclient"
	"user-rest-api/pkg/logger"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
)

type usersRepo struct {
	client dbclient.Client
	logger logger.Logger
}

func NewUsersRepo(client dbclient.Client, logger logger.Logger) repository.Users {
	return &usersRepo{
		client: client,
		logger: logger,
	}
}

func (u *usersRepo) Create(ctx context.Context, user domain.User) (string, error) {
	sql := `
		INSERT INTO users (first_name, last_name, email, phone, login, password) 
		VALUES ($1, $2, $3, $4, $5, $6) 
		RETURNING id
	`

	u.logger.Trace("creating user")
	var userId string
	err := u.client.QueryRow(ctx, sql,
		user.FirstName,
		user.LastName,
		user.Email,
		user.Phone,
		user.Login,
		user.Password).Scan(&userId)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			pgErr = err.(*pgconn.PgError)
			u.logger.Error(fmt.Errorf(fmt.Sprintf("SQL Error: %s, Detail: %s, Where: %s, Code: %s, SQLState: %s",
				pgErr.Message,
				pgErr.Detail,
				pgErr.Where,
				pgErr.Code,
				pgErr.SQLState())))
			return "", apperror.ErrConflict
		}
		return "", err
	}
	return userId, nil
}

func (u *usersRepo) FindAll(ctx context.Context) ([]domain.User, error) {
	sql := `
		SELECT id, first_name, last_name, email, phone, login, password 
		FROM public.users;
	`

	u.logger.Trace("finding all users")
	rows, err := u.client.Query(ctx, sql)
	if err != nil {
		return nil, err
	}

	var (
		users []domain.User
		user  domain.User
	)

	for rows.Next() {
		err = rows.Scan(
			&user.UUID,
			&user.FirstName,
			&user.LastName,
			&user.Email,
			&user.Phone,
			&user.Login,
			&user.Password)
		if err != nil {
			if errors.Is(err, pgx.ErrNoRows) {
				return nil, apperror.ErrNotFound
			}
			var pgErr *pgconn.PgError
			if errors.As(err, &pgErr) {
				pgErr = err.(*pgconn.PgError)

				newErr := fmt.Errorf(fmt.Sprintf("SQL Error: %s, Detail: %s, Where: %s, Code: %s, SQLState: %s",
					pgErr.Message,
					pgErr.Detail,
					pgErr.Where,
					pgErr.Code,
					pgErr.SQLState()))
				u.logger.Error(newErr)

				return nil, newErr
			}
			return nil, err
		}
		users = append(users, user)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return users, nil
}

func (u *usersRepo) FindOne(ctx context.Context, uuid string) (domain.User, error) {
	sql := `
		SELECT id, first_name, last_name, email, phone, login, password
		FROM public.users
		WHERE id = $1
	`

	u.logger.Tracef("finding user with ID: %s", uuid)
	var user domain.User
	err := u.client.QueryRow(ctx, sql, uuid).Scan(
		&user.UUID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Phone,
		&user.Login,
		&user.Password)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return domain.User{}, apperror.ErrNotFound
		}
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			pgErr = err.(*pgconn.PgError)
			newErr := fmt.Errorf(fmt.Sprintf("SQL Error: %s, Detail: %s, Where: %s, Code: %s, SQLState: %s",
				pgErr.Message,
				pgErr.Detail,
				pgErr.Where,
				pgErr.Code,
				pgErr.SQLState()))
			u.logger.Error(newErr)
			return domain.User{}, newErr
		}
		return domain.User{}, err
	}
	return user, nil
}

func (u *usersRepo) FindByEmail(ctx context.Context, email string) (domain.User, error) {
	sql := `
		SELECT id, first_name, last_name, email, phone, login, password
		FROM public.users
		WHERE email = $1
	`

	u.logger.Tracef("finding user with email: %s", email)
	var user domain.User
	err := u.client.QueryRow(ctx, sql, email).Scan(
		&user.UUID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Phone,
		&user.Login,
		&user.Password)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return domain.User{}, apperror.ErrNotFound
		}
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			pgErr = err.(*pgconn.PgError)
			newErr := fmt.Errorf(fmt.Sprintf("SQL Error: %s, Detail: %s, Where: %s, Code: %s, SQLState: %s",
				pgErr.Message,
				pgErr.Detail,
				pgErr.Where,
				pgErr.Code,
				pgErr.SQLState()))
			u.logger.Error(newErr)
			return domain.User{}, newErr
		}
		return domain.User{}, err
	}
	return user, nil
}

func (u *usersRepo) FindByPhone(ctx context.Context, phone string) (domain.User, error) {
	sql := `
		SELECT id, first_name, last_name, email, phone, login, password
		FROM public.users
		WHERE phone = $1
	`

	u.logger.Tracef("finding user with phone: %s", phone)
	var user domain.User
	err := u.client.QueryRow(ctx, sql, phone).Scan(
		&user.UUID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Phone,
		&user.Login,
		&user.Password)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return domain.User{}, apperror.ErrNotFound
		}
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			pgErr = err.(*pgconn.PgError)
			newErr := fmt.Errorf(fmt.Sprintf("SQL Error: %s, Detail: %s, Where: %s, Code: %s, SQLState: %s",
				pgErr.Message,
				pgErr.Detail,
				pgErr.Where,
				pgErr.Code,
				pgErr.SQLState()))
			u.logger.Error(newErr)
			return domain.User{}, newErr
		}
		return domain.User{}, err
	}
	return user, nil
}

func (u *usersRepo) Update(ctx context.Context, uuid string, user domain.User) error {
	sql := `
        UPDATE users
        SET first_name = $1, last_name = $2, email = $3, phone = $4, login = $5, password = $6
        WHERE id = $7
		RETURNING id
    `

	u.logger.Tracef("updating user with ID: %s", uuid)
	var id string
	err := u.client.QueryRow(ctx, sql,
		user.FirstName,
		user.LastName,
		user.Email,
		user.Phone,
		user.Login,
		user.Password,
		uuid).Scan(&id)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			pgErr = err.(*pgconn.PgError)
			u.logger.Error(fmt.Errorf(fmt.Sprintf("SQL Error: %s, Detail: %s, Where: %s, Code: %s, SQLState: %s",
				pgErr.Message,
				pgErr.Detail,
				pgErr.Where,
				pgErr.Code,
				pgErr.SQLState())))
			return apperror.ErrConflict
		}
		return err
	}
	return nil
}

func (u *usersRepo) Delete(ctx context.Context, uuid string) error {
	sql := `
		DELETE FROM users
		WHERE id = $1
		RETURNING id
	`

	u.logger.Tracef("deleting user with ID: %s", uuid)
	var deletedID string
	err := u.client.QueryRow(ctx, sql, uuid).Scan(&deletedID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return apperror.ErrNotFound
		}
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			pgErr = err.(*pgconn.PgError)
			newErr := fmt.Errorf(fmt.Sprintf("SQL Error: %s, Detail: %s, Where: %s, Code: %s, SQLState: %s",
				pgErr.Message,
				pgErr.Detail,
				pgErr.Where,
				pgErr.Code,
				pgErr.SQLState()))
			u.logger.Error(newErr)
			return newErr
		}
		return err
	}
	return nil
}
