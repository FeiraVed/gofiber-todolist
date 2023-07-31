package service

import (
	"context"

	"github.com/FeiraVed/todolist/model/web"
)

type TodolistService interface {
	Create(ctx context.Context, request web.TodolistCreateRequest) web.TodolistResponse
	Update(ctx context.Context, request web.TodolistUpdateRequest) web.TodolistResponse
	Delete(ctx context.Context, todolistId int)
	FindById(ctx context.Context, todolistId int) web.TodolistResponse
	FindAll(ctx context.Context) []web.TodolistResponse
}
