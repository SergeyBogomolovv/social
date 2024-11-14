package usecase

import (
	"context"
	"social/pkg/models"
	proto "social/pkg/proto/generated"
	"time"
)

type postsUsecase struct {
	client proto.PostServiceClient
}

func (u *postsUsecase) CreatePost(ctx context.Context, dto *models.CreatePostDto) (*models.Post, error) {
	res, err := u.client.CreatePost(ctx, &proto.CreatePostRequest{AuthorId: dto.AuthorID, Title: dto.Title, Content: dto.Content})
	if err != nil {
		return nil, err
	}

	post := &models.Post{
		ID:        res.Post.Id,
		AuthorID:  res.Post.AuthorId,
		Title:     res.Post.Title,
		Content:   res.Post.Content,
		CreatedAt: time.Unix(res.Post.CreatedAt, 0),
	}

	return post, nil
}

func (u *postsUsecase) GetPost(ctx context.Context, id int64) (*models.Post, error) {
	res, err := u.client.GetPostById(ctx, &proto.GetPostByIdRequest{Id: id})
	if err != nil {
		return nil, err
	}

	post := &models.Post{
		ID:        res.Post.Id,
		AuthorID:  res.Post.AuthorId,
		Title:     res.Post.Title,
		Content:   res.Post.Content,
		CreatedAt: time.Unix(res.Post.CreatedAt, 0),
	}

	return post, nil
}

func (u *postsUsecase) GetManyPosts(ctx context.Context, page int32, limit int32) ([]*models.Post, error) {
	res, err := u.client.GetPosts(ctx, &proto.GetPostsRequest{Page: page, Limit: limit})
	if err != nil {
		return nil, err
	}

	posts := make([]*models.Post, len(res.Posts))

	for i, post := range res.Posts {
		posts[i] = &models.Post{
			ID:        post.Id,
			AuthorID:  post.AuthorId,
			Title:     post.Title,
			Content:   post.Content,
			CreatedAt: time.Unix(post.CreatedAt, 0),
		}
	}

	return posts, nil
}

func (u *postsUsecase) UpdatePost(ctx context.Context, dto *models.UpdatePostDto) (*models.Post, error) {
	res, err := u.client.UpdatePost(ctx, &proto.UpdatePostRequest{Id: dto.ID, Title: dto.Title, Content: dto.Content})
	if err != nil {
		return nil, err
	}

	post := &models.Post{
		ID:        res.Post.Id,
		AuthorID:  res.Post.AuthorId,
		Title:     res.Post.Title,
		Content:   res.Post.Content,
		CreatedAt: time.Unix(res.Post.CreatedAt, 0),
	}

	return post, nil
}

func (u *postsUsecase) DeletePost(ctx context.Context, id int64) (*models.Post, error) {
	res, err := u.client.DeletePost(ctx, &proto.DeletePostRequest{Id: id})
	if err != nil {
		return nil, err
	}

	post := &models.Post{
		ID:        res.Post.Id,
		AuthorID:  res.Post.AuthorId,
		Title:     res.Post.Title,
		Content:   res.Post.Content,
		CreatedAt: time.Unix(res.Post.CreatedAt, 0),
	}

	return post, nil
}

func NewPostsUsecase(client proto.PostServiceClient) *postsUsecase {
	return &postsUsecase{
		client: client,
	}
}
