package controller

import (
	"context"
	"social/pkg/constants"
	"social/pkg/models"
	proto "social/pkg/proto/generated"

	"github.com/go-playground/validator/v10"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type PostUsecase interface {
	GetManyPosts(ctx context.Context, page int32, limit int32) ([]*models.Post, error)
	GetPost(ctx context.Context, id int64) (*models.Post, error)
	CreatePost(ctx context.Context, dto models.CreatePostDto) (*models.Post, error)
	UpdatePost(ctx context.Context, dto models.UpdatePostDto) (*models.Post, error)
	DeletePost(ctx context.Context, id int64) (*models.Post, error)
}

type postsContoller struct {
	usecase  PostUsecase
	validate *validator.Validate
	proto.UnimplementedPostServiceServer
}

func RegisterPostsController(grpc *grpc.Server, usecase PostUsecase) {
	gRPCHandler := &postsContoller{
		usecase:  usecase,
		validate: validator.New(validator.WithRequiredStructEnabled()),
	}

	proto.RegisterPostServiceServer(grpc, gRPCHandler)
}

func (c *postsContoller) CreatePost(ctx context.Context, req *proto.CreatePostRequest) (*proto.CreatePostResponse, error) {
	dto := models.CreatePostDto{
		AuthorID: req.AuthorId,
		Content:  req.Content,
		Title:    req.Title,
	}
	err := c.validate.Struct(dto)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "error: %v", err)
	}

	post, err := c.usecase.CreatePost(ctx, dto)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "error: %v", err)
	}

	return &proto.CreatePostResponse{
		Post: post.ToProto(),
	}, nil
}

func (c *postsContoller) GetPosts(ctx context.Context, req *proto.GetPostsRequest) (*proto.GetPostsResponse, error) {
	posts, err := c.usecase.GetManyPosts(ctx, req.Page, req.Limit)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error: %v", err)
	}

	mapped := make([]*proto.Post, len(posts))

	for _, post := range posts {
		mapped = append(mapped, post.ToProto())
	}

	return &proto.GetPostsResponse{
		Posts: mapped,
	}, nil
}

func (c *postsContoller) GetPostById(ctx context.Context, req *proto.GetPostByIdRequest) (*proto.GetPostByIdResponse, error) {
	post, err := c.usecase.GetPost(ctx, req.Id)
	switch err {
	case constants.ErrPostNotFound:
		return nil, status.Errorf(codes.NotFound, "error: %v", err)
	case nil:
		return &proto.GetPostByIdResponse{
			Post: post.ToProto(),
		}, nil
	default:
		return nil, status.Errorf(codes.Internal, "error: %v", err)
	}
}

func (c *postsContoller) UpdatePost(ctx context.Context, req *proto.UpdatePostRequest) (*proto.UpdatePostResponse, error) {
	dto := models.UpdatePostDto{
		ID:      req.Id,
		Title:   req.Title,
		Content: req.Content,
	}

	err := c.validate.Struct(dto)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "error: %v", err)
	}

	post, err := c.usecase.UpdatePost(ctx, dto)

	switch err {
	case constants.ErrPostNotFound:
		return nil, status.Errorf(codes.NotFound, "error: %v", err)
	case nil:
		return &proto.UpdatePostResponse{
			Post: post.ToProto(),
		}, nil
	default:
		return nil, status.Errorf(codes.Internal, "error: %v", err)
	}
}

func (c *postsContoller) DeletePost(ctx context.Context, req *proto.DeletePostRequest) (*proto.DeletePostResponse, error) {
	post, err := c.usecase.DeletePost(ctx, req.Id)
	switch err {
	case constants.ErrPostNotFound:
		return nil, status.Errorf(codes.NotFound, "error: %v", err)
	case nil:
		return &proto.DeletePostResponse{
			Post: post.ToProto(),
		}, nil
	default:
		return nil, status.Errorf(codes.Internal, "error: %v", err)
	}
}
