package usecase

import (
	"context"
	"social/pkg/models"
	proto "social/pkg/proto/generated"
	"time"
)

type usersUsecase struct {
	client proto.UserServiceClient
}

func (u *usersUsecase) GetUsers(ctx context.Context, page int32, limit int32) ([]*models.UserPayload, error) {
	res, err := u.client.GetUsers(ctx, &proto.GetUsersRequest{Page: page, Limit: limit})
	if err != nil {
		return nil, err
	}

	users := make([]*models.UserPayload, len(res.Users))

	for i, user := range res.Users {
		users[i] = &models.UserPayload{
			ID:        user.Id,
			Username:  user.Username,
			CreatedAt: time.Unix(int64(user.CreatedAt), 0),
		}
	}

	return users, nil
}

func (u *usersUsecase) CreateUser(ctx context.Context, dto *models.CreateUserDto) (*models.UserPayload, error) {
	res, err := u.client.CreateUser(ctx, &proto.CreateUserRequest{Username: dto.Username, Password: dto.Password})
	if err != nil {
		return nil, err
	}

	user := &models.UserPayload{
		ID:        res.User.Id,
		Username:  res.User.Username,
		CreatedAt: time.Unix(res.User.CreatedAt, 0),
	}

	return user, nil
}

func (u *usersUsecase) GetUser(ctx context.Context, id int64) (*models.UserPayload, error) {
	res, err := u.client.GetUserById(ctx, &proto.GetUserByIdRequest{Id: id})
	if err != nil {
		return nil, err
	}

	user := &models.UserPayload{
		ID:        res.User.Id,
		Username:  res.User.Username,
		CreatedAt: time.Unix(res.User.CreatedAt, 0),
	}

	return user, nil
}

func (u *usersUsecase) UpdateUser(ctx context.Context, dto *models.UpdateUserDto) (*models.UserPayload, error) {
	res, err := u.client.UpdateUser(ctx, &proto.UpdateUserRequest{Id: dto.ID, Username: dto.Username})
	if err != nil {
		return nil, err
	}

	user := &models.UserPayload{
		ID:        res.User.Id,
		Username:  res.User.Username,
		CreatedAt: time.Unix(res.User.CreatedAt, 0),
	}

	return user, nil
}

func (u *usersUsecase) DeleteUser(ctx context.Context, id int64) (*models.UserPayload, error) {
	res, err := u.client.DeleteUser(ctx, &proto.DeleteUserRequest{Id: id})
	if err != nil {
		return nil, err
	}

	user := &models.UserPayload{
		ID:        res.User.Id,
		Username:  res.User.Username,
		CreatedAt: time.Unix(res.User.CreatedAt, 0),
	}

	return user, nil
}

func NewUsersUsecase(client proto.UserServiceClient) *usersUsecase {
	return &usersUsecase{
		client: client,
	}
}
