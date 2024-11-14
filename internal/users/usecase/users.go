package usecase

import (
	"context"
	"social/pkg/constants"
	"social/pkg/models"
)

type Repository interface {
	GetIsExists(ctx context.Context, username string) (bool, error)
	GetMany(ctx context.Context, page int32, limit int32) ([]*models.User, error)
	GetById(ctx context.Context, id int64) (*models.User, error)
	Create(ctx context.Context, dto *models.CreateUserDto) (*models.User, error)
	Update(ctx context.Context, dto *models.UpdateUserDto) (*models.User, error)
	Delete(ctx context.Context, id int64) (*models.User, error)
}

type usersUsecase struct {
	repo Repository
}

func (u *usersUsecase) GetUsers(ctx context.Context, page int32, limit int32) ([]*models.User, error) {
	if page < 0 {
		page = 1
	}
	if limit < 0 {
		limit = 10
	}
	return u.repo.GetMany(ctx, page, limit)
}

func (u *usersUsecase) GetUser(ctx context.Context, id int64) (*models.User, error) {
	return u.repo.GetById(ctx, id)
}

func (u *usersUsecase) CreateUser(ctx context.Context, dto *models.CreateUserDto) (*models.User, error) {
	isExists, err := u.repo.GetIsExists(ctx, dto.Username)
	if err != nil {
		return nil, err
	}
	if isExists {
		return nil, constants.ErrUserAlreadyExists
	}
	return u.repo.Create(ctx, dto)
}

func (u *usersUsecase) UpdateUser(ctx context.Context, dto *models.UpdateUserDto) (*models.User, error) {
	currUser, err := u.repo.GetById(ctx, dto.ID)
	if err != nil {
		return nil, err
	}
	if currUser.Username == dto.Username {
		return currUser, nil
	}
	isExists, err := u.repo.GetIsExists(ctx, dto.Username)
	if err != nil {
		return nil, err
	}
	if isExists {
		return nil, constants.ErrUserAlreadyExists
	}
	return u.repo.Update(ctx, dto)
}

func (u *usersUsecase) DeleteUser(ctx context.Context, id int64) (*models.User, error) {
	return u.repo.Delete(ctx, id)
}

func NewUsersUsecase(repo Repository) *usersUsecase {
	return &usersUsecase{
		repo: repo,
	}
}
