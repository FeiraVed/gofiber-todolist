package service

import (
	"context"
	"database/sql"

	"github.com/FeiraVed/todolist/helper"
	"github.com/FeiraVed/todolist/model"
	"github.com/FeiraVed/todolist/model/web"
	"github.com/FeiraVed/todolist/repository"
	"github.com/go-playground/validator/v10"
)

type TodolistServiceImpl struct {
	repository repository.TodolistRepository
	db         *sql.DB
	validate   *validator.Validate
}

func (service *TodolistServiceImpl) Create(
	ctx context.Context,
	request web.TodolistCreateRequest,
) (_ web.TodolistResponse) {
	tx, err := service.db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	errValidate := service.validate.Struct(request)
	helper.PanicIfError(errValidate)

	todolist := model.Todolist{
		Name: request.Name,
	}
	result := service.repository.Save(ctx, tx, todolist)
	response := web.TodolistResponse{
		Id:   result.Id,
		Name: result.Name,
	}

	return response
}

func (service *TodolistServiceImpl) Update(
	ctx context.Context,
	request web.TodolistUpdateRequest,
) (_ web.TodolistResponse) {
	tx, err := service.db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	errValidate := service.validate.Struct(request)
	helper.PanicIfError(errValidate)

	todolist := model.Todolist{
		Id:   request.Id,
		Name: request.Name,
	}

	_, errNotFound := service.repository.FindById(ctx, tx, request.Id)

	helper.PanicIfError(errNotFound)

	result := service.repository.Update(ctx, tx, todolist)
	response := web.TodolistResponse{
		Id:   todolist.Id,
		Name: result.Name,
	}

	return response
}
func (service *TodolistServiceImpl) Delete(ctx context.Context, todolistId int) {
	tx, err := service.db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	_, err2 := service.repository.FindById(ctx, tx, todolistId)
	helper.PanicIfError(err2)

	service.repository.Delete(ctx, tx, todolistId)
}

func (service *TodolistServiceImpl) FindById(
	ctx context.Context,
	todolistId int,
) (_ web.TodolistResponse) {
	tx, err := service.db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	result, err2 := service.repository.FindById(ctx, tx, todolistId)
	helper.PanicIfError(err2)

	return web.TodolistResponse{
		Id:   result.Id,
		Name: result.Name,
	}
}

func (service *TodolistServiceImpl) FindAll(ctx context.Context) (_ []web.TodolistResponse) {
	todolists := []web.TodolistResponse{}
	tx, err := service.db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	results := service.repository.FindAll(ctx, tx)
	for _, t := range results {
		todolist := web.TodolistResponse{Id: t.Id, Name: t.Name}
		todolists = append(todolists, todolist)

	}

	return todolists
}

func NewTodolistServiceImpl(
	repository repository.TodolistRepository,
	db *sql.DB,
	validate *validator.Validate,
) *TodolistServiceImpl {

	return &TodolistServiceImpl{
		repository: repository,
		db:         db,
		validate:   validate,
	}
}
