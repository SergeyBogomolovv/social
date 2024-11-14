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

type UserUsecase interface {
	GetUser(ctx context.Context, id int64) (*models.User, error)
	GetUsers(ctx context.Context, page int32, limit int32) ([]*models.User, error)
	CreateUser(ctx context.Context, dto *models.CreateUserDto) (*models.User, error)
	UpdateUser(ctx context.Context, dto *models.UpdateUserDto) (*models.User, error)
	DeleteUser(ctx context.Context, id int64) (*models.User, error)
}

type usersController struct {
	usecase  UserUsecase
	validate *validator.Validate
	proto.UnimplementedUserServiceServer
}

func RegisterUsersController(grpcServer *grpc.Server, usecase UserUsecase) {
	gRPCHandler := &usersController{
		usecase:  usecase,
		validate: validator.New(),
	}
	proto.RegisterUserServiceServer(grpcServer, gRPCHandler)
}

func (c *usersController) GetUserById(ctx context.Context, req *proto.GetUserByIdRequest) (*proto.GetUserByIdResponse, error) {
	user, err := c.usecase.GetUser(ctx, req.Id)
	if err != nil {
		if err == constants.ErrUserNotFound {
			return nil, status.Errorf(codes.NotFound, "user not found")
		}
		return nil, status.Errorf(codes.Internal, "failed to get user: %v", err)
	}
	return &proto.GetUserByIdResponse{User: user.ToProto()}, nil
}

func (c *usersController) GetUsers(ctx context.Context, req *proto.GetUsersRequest) (*proto.GetUsersResponse, error) {
	users, err := c.usecase.GetUsers(ctx, req.Page, req.Limit)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get users: %v", err)
	}

	response := &proto.GetUsersResponse{}
	for _, user := range users {
		response.Users = append(response.Users, user.ToProto())
	}
	return response, nil
}

func (c *usersController) CreateUser(ctx context.Context, req *proto.CreateUserRequest) (*proto.CreateUserResponse, error) {
	dto := &models.CreateUserDto{
		Username: req.Username,
		Password: req.Password,
	}

	if err := c.validate.Struct(dto); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid request data: %v", err)
	}

	user, err := c.usecase.CreateUser(ctx, dto)
	if err != nil {
		if err == constants.ErrUserAlreadyExists {
			return nil, status.Errorf(codes.AlreadyExists, "user already exists")
		}
		return nil, status.Errorf(codes.Internal, "failed to create user: %v", err)
	}

	return &proto.CreateUserResponse{User: user.ToProto()}, nil
}

func (c *usersController) UpdateUser(ctx context.Context, req *proto.UpdateUserRequest) (*proto.UpdateUserResponse, error) {
	dto := &models.UpdateUserDto{
		ID:       req.Id,
		Username: req.Username,
	}

	if err := c.validate.Struct(dto); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid request data: %v", err)
	}

	user, err := c.usecase.UpdateUser(ctx, dto)
	if err != nil {
		if err == constants.ErrUserNotFound {
			return nil, status.Errorf(codes.NotFound, "user not found")
		}
		if err == constants.ErrUserAlreadyExists {
			return nil, status.Errorf(codes.AlreadyExists, "user already exists")
		}
		return nil, status.Errorf(codes.Internal, "failed to update user: %v", err)
	}

	return &proto.UpdateUserResponse{User: user.ToProto()}, nil
}

func (c *usersController) DeleteUser(ctx context.Context, req *proto.DeleteUserRequest) (*proto.DeleteUserResponse, error) {
	user, err := c.usecase.DeleteUser(ctx, req.Id)
	if err != nil {
		if err == constants.ErrUserNotFound {
			return nil, status.Errorf(codes.NotFound, "user not found")
		}
		return nil, status.Errorf(codes.Internal, "failed to delete user: %v", err)
	}

	return &proto.DeleteUserResponse{User: user.ToProto()}, nil
}
