package controller

import (
	"context"
	"net/http"
	"social/pkg/models"
	"social/pkg/utils"

	"github.com/go-playground/validator/v10"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type PostUsecase interface {
	CreatePost(ctx context.Context, dto *models.CreatePostDto) (*models.Post, error)
	GetPost(ctx context.Context, id int64) (*models.Post, error)
	GetManyPosts(ctx context.Context, page, limit int32) ([]*models.Post, error)
	UpdatePost(ctx context.Context, dto *models.UpdatePostDto) (*models.Post, error)
	DeletePost(ctx context.Context, id int64) (*models.Post, error)
}

type postsController struct {
	usecase  PostUsecase
	validate *validator.Validate
}

func (c *postsController) CreatePost(w http.ResponseWriter, r *http.Request) {
	dto := new(models.CreatePostDto)

	if err := utils.ParseJSON(r, dto); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := c.validate.Struct(dto); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	post, err := c.usecase.CreatePost(r.Context(), dto)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, post)
}

func (c *postsController) GetPost(w http.ResponseWriter, r *http.Request) {
	id, err := utils.GetIntParam("id", r)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	post, err := c.usecase.GetPost(r.Context(), int64(id))
	if err != nil {
		switch status.Code(err) {
		case codes.NotFound:
			utils.WriteError(w, http.StatusNotFound, err)
		default:
			utils.WriteError(w, http.StatusInternalServerError, err)
		}
		return
	}

	utils.WriteJSON(w, http.StatusOK, post)
}

func (c *postsController) GetPosts(w http.ResponseWriter, r *http.Request) {
	page := utils.GetIntQuery("page", 1, r)
	limit := utils.GetIntQuery("limit", 10, r)

	posts, err := c.usecase.GetManyPosts(r.Context(), int32(page), int32(limit))
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, posts)
}

func (c *postsController) UpdatePost(w http.ResponseWriter, r *http.Request) {
	dto := new(models.UpdatePostDto)
	if err := utils.ParseJSON(r, dto); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := c.validate.Struct(dto); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	post, err := c.usecase.UpdatePost(r.Context(), dto)
	if err != nil {
		switch status.Code(err) {
		case codes.NotFound:
			utils.WriteError(w, http.StatusNotFound, err)
		default:
			utils.WriteError(w, http.StatusInternalServerError, err)
		}
		return
	}

	utils.WriteJSON(w, http.StatusOK, post)
}

func (c *postsController) DeletePost(w http.ResponseWriter, r *http.Request) {
	id, err := utils.GetIntParam("id", r)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	post, err := c.usecase.DeletePost(r.Context(), int64(id))
	if err != nil {
		switch status.Code(err) {
		case codes.NotFound:
			utils.WriteError(w, http.StatusNotFound, err)
		default:
			utils.WriteError(w, http.StatusInternalServerError, err)
		}
		return
	}

	utils.WriteJSON(w, http.StatusOK, post)
}

func RegisterPostsController(router *http.ServeMux, usecase PostUsecase) {
	controller := &postsController{
		usecase:  usecase,
		validate: validator.New(validator.WithRequiredStructEnabled()),
	}

	posts := http.NewServeMux()

	posts.HandleFunc("POST /", controller.CreatePost)
	posts.HandleFunc("GET /{id}", controller.GetPost)
	posts.HandleFunc("GET /", controller.GetPosts)
	posts.HandleFunc("PUT /update", controller.UpdatePost)
	posts.HandleFunc("DELETE /{id}", controller.DeletePost)

	router.Handle("/posts/", http.StripPrefix("/posts", posts))
}
