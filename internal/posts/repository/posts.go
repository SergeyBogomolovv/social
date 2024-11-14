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
	query := `SELECT * FROM posts LIMIT $1 OFFSET $2`
	posts := make([]*models.Post, 0)
	err := r.db.SelectContext(ctx, &posts, query, limit, page)
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (r *postsRepository) GetById(ctx context.Context, id int64) (*models.Post, error) {
	query := `SELECT * FROM users WHERE post_id = $1`
	post := new(models.Post)
	err := r.db.GetContext(ctx, post, query, id)
	if err != nil {
		return nil, err
	}
	return post, nil
}

func (r *postsRepository) Create(ctx context.Context, dto models.CreatePostDto) (*models.Post, error) {
	query := `INSERT INTO posts (author_id, title, content) VALUES ($1, $2, $3) RETURNING *`
	post := new(models.Post)
	err := r.db.GetContext(ctx, post, query, dto.AuthorID, dto.Title, dto.Content)
	if err != nil {
		return nil, err
	}
	return post, nil
}

func (r *postsRepository) Update(ctx context.Context, dto models.UpdatePostDto) (*models.Post, error) {
	query := `UPDATE posts SET title = $1, content = $2 WHERE post_id = $3 RETURNING *`
	post := new(models.Post)
	err := r.db.GetContext(ctx, post, query, dto.Title, dto.Content, dto.ID)
	if err != nil {
		return nil, err
	}
	return post, nil
}

func (r *postsRepository) Delete(ctx context.Context, id int64) (*models.Post, error) {
	query := `DELETE FROM posts WHERE post_id = $1 RETURNING *`
	post := new(models.Post)
	err := r.db.GetContext(ctx, post, query, id)
	if err != nil {
		return nil, err
	}
	return post, nil
}

func NewPostsRepository(db *sqlx.DB) *postsRepository {
	return &postsRepository{db: db}
}
