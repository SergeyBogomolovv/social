package service

import (
	"context"
	"social/pkg/models"
)

type Repository interface {
	GetMany(ctx context.Context, page int32, limit int32) ([]*models.Post, error)
	GetById(ctx context.Context, id int64) (*models.Post, error)
	Create(ctx context.Context, dto models.CreatePostDto) (*models.Post, error)
	Update(ctx context.Context, dto models.UpdatePostDto) (*models.Post, error)
	Delete(ctx context.Context, id int64) (*models.Post, error)
}

type postsUsecase struct {
	repo Repository
}

func (u *postsUsecase) GetManyPosts(ctx context.Context, page int32, limit int32) ([]*models.Post, error) {
	if page < 0 {
		page = 1
	}
	if limit < 0 {
		limit = 10
	}
	return u.repo.GetMany(ctx, page, limit)
}

func (u *postsUsecase) GetPost(ctx context.Context, id int64) (*models.Post, error) {
	return u.repo.GetById(ctx, id)
}

func (u *postsUsecase) CreatePost(ctx context.Context, dto models.CreatePostDto) (*models.Post, error) {
	return u.repo.Create(ctx, dto)
}

func (u *postsUsecase) UpdatePost(ctx context.Context, dto models.UpdatePostDto) (*models.Post, error) {
	post, err := u.repo.GetById(ctx, dto.ID)
	if err != nil {
		return nil, err
	}
	if dto.Content == "" {
		dto.Content = post.Content
	}
	if dto.Title == "" {
		dto.Title = post.Title
	}
	return u.repo.Update(ctx, dto)
}

func (u *postsUsecase) DeletePost(ctx context.Context, id int64) (*models.Post, error) {
	return u.repo.Delete(ctx, id)
}

func NewPostsUsecase(repo Repository) *postsUsecase {
	return &postsUsecase{
		repo: repo,
	}
}
