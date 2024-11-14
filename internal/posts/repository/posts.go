package repository

import (
	"context"
	"social/pkg/models"

	"github.com/jmoiron/sqlx"
)

type postsRepository struct {
	db *sqlx.DB
}

func (r *postsRepository) GetMany(ctx context.Context, page int32, limit int32) ([]*models.Post, error) {
	return nil, nil
}

func (r *postsRepository) GetById(ctx context.Context, id int64) (*models.Post, error) {
	return nil, nil
}

func (r *postsRepository) Create(ctx context.Context, dto models.CreatePostDto) (*models.Post, error) {
	return nil, nil
}

func (r *postsRepository) Update(ctx context.Context, dto models.UpdatePostDto) (*models.Post, error) {
	return nil, nil
}

func (r *postsRepository) Delete(ctx context.Context, id int64) (*models.Post, error) {
	return nil, nil
}

func NewPostsRepository(db *sqlx.DB) *postsRepository {
	return &postsRepository{
		db: db,
	}
}
