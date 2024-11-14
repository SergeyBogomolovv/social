package repository

import (
	"context"
	"database/sql"
	"social/pkg/constants"
	"social/pkg/models"

	"github.com/jmoiron/sqlx"
)

type usersRepository struct {
	db *sqlx.DB
}

func (r *usersRepository) GetIsExists(ctx context.Context, username string) (bool, error) {
	query := `SELECT TRUE FROM users WHERE username = $1`
	var isExists bool
	err := r.db.GetContext(ctx, &isExists, query, username)
	switch err {
	case nil:
		return true, nil
	case sql.ErrNoRows:
		return false, nil
	default:
		return false, err
	}
}

func (r *usersRepository) GetMany(ctx context.Context, page int32, limit int32) ([]*models.User, error) {
	query := `SELECT * FROM users LIMIT $1 OFFSET $2`
	users := make([]*models.User, 0)
	err := r.db.SelectContext(ctx, &users, query, limit, page)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r *usersRepository) GetById(ctx context.Context, id int64) (*models.User, error) {
	query := `SELECT * FROM users WHERE user_id = $1`
	user := new(models.User)
	err := r.db.GetContext(ctx, user, query, id)
	switch err {
	case sql.ErrNoRows:
		return nil, constants.ErrUserNotFound
	case nil:
		return user, nil
	default:
		return nil, err
	}
}

func (r *usersRepository) Create(ctx context.Context, dto *models.CreateUserDto) (*models.User, error) {
	query := `INSERT INTO users (username, password) VALUES ($1, $2) RETURNING *`
	user := new(models.User)
	err := r.db.GetContext(ctx, user, query, dto.Username, dto.Password)
	switch err {
	case sql.ErrNoRows:
		return nil, constants.ErrUserAlreadyExists
	case nil:
		return user, nil
	default:
		return nil, err
	}
}

func (r *usersRepository) Update(ctx context.Context, dto *models.UpdateUserDto) (*models.User, error) {
	query := `UPDATE users SET username = $1 WHERE user_id = $2 RETURNING *`
	user := new(models.User)
	err := r.db.GetContext(ctx, user, query, dto.Username, dto.ID)
	switch err {
	case sql.ErrNoRows:
		return nil, constants.ErrUserNotFound
	case nil:
		return user, nil
	default:
		return nil, err
	}
}

func (r *usersRepository) Delete(ctx context.Context, id int64) (*models.User, error) {
	query := `DELETE FROM users WHERE user_id = $1 RETURNING *`
	user := new(models.User)
	err := r.db.GetContext(ctx, user, query, id)
	switch err {
	case sql.ErrNoRows:
		return nil, constants.ErrUserNotFound
	case nil:
		return user, nil
	default:
		return nil, err
	}
}

func NewUsersRepository(db *sqlx.DB) *usersRepository {
	return &usersRepository{
		db: db,
	}
}
