package repository

import (
	"context"
	"social/pkg/models"

	"github.com/jmoiron/sqlx"
)

type usersRepository struct {
	db *sqlx.DB
}

func (r *usersRepository) GetMany(ctx context.Context, page int32, limit int32) ([]*models.User, error) {
	return nil, nil
}

func (r *usersRepository) GetById(ctx context.Context, id int64) (*models.User, error) {
	return nil, nil
}

func (r *usersRepository) Create(ctx context.Context, dto models.CreateUserDto) (*models.User, error) {
	return nil, nil
}

func (r *usersRepository) Update(ctx context.Context, dto models.UpdateUserDto) (*models.User, error) {
	return nil, nil
}

func (r *usersRepository) Delete(ctx context.Context, id int64) (*models.User, error) {
	return nil, nil
}

func NewUsersRepository(db *sqlx.DB) *usersRepository {
	return &usersRepository{
		db: db,
	}
}
