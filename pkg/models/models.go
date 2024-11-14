package models

import (
	proto "social/pkg/proto/generated"
	"time"
)

type User struct {
	ID        int64     `json:"id" db:"user_id"`
	Username  string    `json:"username" db:"username"`
	Password  string    `json:"-" db:"password"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

func (u *User) ToProto() *proto.User {
	return &proto.User{
		Id:        u.ID,
		Username:  u.Username,
		CreatedAt: u.CreatedAt.Unix(),
	}
}

type UserPayload struct {
	ID        int64     `json:"id"`
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"created_at"`
}

type Post struct {
	ID        int64     `json:"id" db:"post_id"`
	AuthorID  int64     `json:"author_id" db:"author_id"`
	Title     string    `json:"title" db:"title"`
	Content   string    `json:"content" db:"content"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

func (p *Post) ToProto() *proto.Post {
	return &proto.Post{
		Id:        p.ID,
		AuthorId:  p.AuthorID,
		Title:     p.Title,
		Content:   p.Content,
		CreatedAt: p.CreatedAt.Unix(),
	}
}

type CreatePostDto struct {
	AuthorID int64  `json:"author_id" validate:"required"`
	Title    string `json:"title" validate:"required"`
	Content  string `json:"content" validate:"required"`
}

type UpdatePostDto struct {
	ID      int64  `json:"id" validate:"required"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type CreateUserDto struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required,min=6"`
}

type UpdateUserDto struct {
	ID       int64  `json:"id" validate:"required"`
	Username string `json:"username"`
}
