package controller

import (
	"context"
	"net/http"
	"social/pkg/models"
	"social/pkg/utils"

	"github.com/go-playground/validator/v10"
)

type UserUsecase interface {
	GetUsers(ctx context.Context, page int32, limit int32) ([]*models.UserPayload, error)
	CreateUser(ctx context.Context, dto *models.CreateUserDto) (*models.UserPayload, error)
	GetUser(ctx context.Context, id int64) (*models.UserPayload, error)
	UpdateUser(ctx context.Context, dto *models.UpdateUserDto) (*models.UserPayload, error)
	DeleteUser(ctx context.Context, id int64) (*models.UserPayload, error)
}

type usersController struct {
	usecase  UserUsecase
	validate *validator.Validate
}

func (c *usersController) GetUsers(w http.ResponseWriter, r *http.Request) {
	page := utils.GetIntQuery("page", 1, r)
	limit := utils.GetIntQuery("limit", 10, r)

	users, err := c.usecase.GetUsers(r.Context(), int32(page), int32(limit))
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, users)
}

func (c *usersController) CreateUser(w http.ResponseWriter, r *http.Request) {
	dto := new(models.CreateUserDto)

	if err := utils.ParseJSON(r, dto); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := c.validate.Struct(dto); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	user, err := c.usecase.CreateUser(r.Context(), dto)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, user)
}

func (c *usersController) GetUser(w http.ResponseWriter, r *http.Request) {
	id, err := utils.GetIntParam("id", r)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	user, err := c.usecase.GetUser(r.Context(), int64(id))
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, user)
}

func (c *usersController) UpdateUser(w http.ResponseWriter, r *http.Request) {
	dto := new(models.UpdateUserDto)
	if err := utils.ParseJSON(r, dto); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := c.validate.Struct(dto); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	user, err := c.usecase.UpdateUser(r.Context(), dto)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, user)
}

func (c *usersController) DeleteUser(w http.ResponseWriter, r *http.Request) {
	id, err := utils.GetIntParam("id", r)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	user, err := c.usecase.DeleteUser(r.Context(), int64(id))
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, user)
}

func RegisterUsersController(router *http.ServeMux, usecase UserUsecase) {
	controller := &usersController{
		usecase:  usecase,
		validate: validator.New(validator.WithRequiredStructEnabled()),
	}
	users := http.NewServeMux()

	users.HandleFunc("GET /", controller.GetUsers)
	users.HandleFunc("POST /", controller.CreateUser)
	users.HandleFunc("GET /{id}", controller.GetUser)
	users.HandleFunc("PUT /update", controller.UpdateUser)
	users.HandleFunc("DELETE /{id}", controller.DeleteUser)

	router.Handle("/users/", http.StripPrefix("/users", users))
}
