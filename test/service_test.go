package test

import (
	"context"
	"testing"

	"github.com/FeiraVed/todolist/helper"
	"github.com/FeiraVed/todolist/injector"
	"github.com/FeiraVed/todolist/model/web"
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
)

func TestServiceCreateSuccess(t *testing.T) {
	db := SetupNewDb()
	TruncateTodolist(db)
	defer db.Close()

	validate := validator.New()
	service := injector.InitializedService(db, validate)
	ctx := context.Background()

	request := web.TodolistCreateRequest{
		Name: "Belajar Golang",
	}

	response := service.Create(ctx, request)
	assert.Equal(t, request.Name, response.Name)

}

func TestServiceCreateFailed(t *testing.T) {
	defer func() {
		err := recover()
		assert.NotNil(t, err)
	}()

	db := SetupNewDb()
	TruncateTodolist(db)
	defer db.Close()

	validate := validator.New()
	service := injector.InitializedService(db, validate)
	ctx := context.Background()

	request := web.TodolistCreateRequest{
		Name: "",
	}

	response := service.Create(ctx, request)
	assert.NotEqual(t, request.Name, response.Name)

}

func TestServiceUpdateSuccess(t *testing.T) {
	db := SetupNewDb()
	TruncateTodolist(db)
	defer db.Close()
	validate := validator.New()
	ctx := context.Background()
	service := injector.InitializedService(db, validate)
	service.Create(ctx, web.TodolistCreateRequest{Name: "Belajar Golang"})

	request := web.TodolistUpdateRequest{
		Id:   1,
		Name: "Belajar NodeJS",
	}

	response := service.Update(ctx, request)
	assert.Equal(t, request.Name, response.Name)

}

func TestServiceUpdateFailedName(t *testing.T) {
	defer func() {
		error := recover()
		assert.NotNil(t, error)
	}()

	db := SetupNewDb()
	TruncateTodolist(db)
	defer db.Close()
	validate := validator.New()
	ctx := context.Background()
	service := injector.InitializedService(db, validate)
	service.Create(ctx, web.TodolistCreateRequest{Name: "Belajar Golang"})

	request := web.TodolistUpdateRequest{
		Id:   1,
		Name: "",
	}

	response := service.Update(ctx, request)
	assert.Equal(t, request.Name, response.Name)

}
func TestServiceUpdateFailedId(t *testing.T) {
	defer func() {
		error := recover()
		assert.NotNil(t, error)
	}()

	db := SetupNewDb()
	TruncateTodolist(db)
	defer db.Close()
	validate := validator.New()
	ctx := context.Background()
	service := injector.InitializedService(db, validate)
	service.Create(ctx, web.TodolistCreateRequest{Name: "Belajar Golang"})

	request := web.TodolistUpdateRequest{
		Id:   4,
		Name: "Belajar NodeJS",
	}

	responseUpdate := service.Update(ctx, request)
	assert.Equal(t, request.Name, responseUpdate.Name)

}

func TestServiceDeleteSucess(t *testing.T) {
	db := SetupNewDb()
	TruncateTodolist(db)
	defer db.Close()
	validate := validator.New()
	ctx := context.Background()

	tx, err := db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)
	service := injector.InitializedService(db, validate)
	service.Create(ctx, web.TodolistCreateRequest{Name: "Belajar Golang"})

	service.Delete(ctx, 1)
	_, err2 := todolistRepository.FindById(ctx, tx, 1)
	assert.NotNil(t, err2)

}

func TestServiceDeleteFailed(t *testing.T) {
	defer func() {
		error := recover()
		assert.NotNil(t, error)
	}()

	db := SetupNewDb()
	TruncateTodolist(db)
	defer db.Close()
	validate := validator.New()
	ctx := context.Background()

	tx, err := db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)
	service := injector.InitializedService(db, validate)
	service.Create(ctx, web.TodolistCreateRequest{Name: "Belajar Golang"})

	service.Delete(ctx, 2)

}

func TestFindByIdSuccess(t *testing.T) {
	db := SetupNewDb()
	TruncateTodolist(db)
	defer db.Close()
	validate := validator.New()
	ctx := context.Background()

	tx, err := db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)
	service := injector.InitializedService(db, validate)
	responseCreate := service.Create(ctx, web.TodolistCreateRequest{Name: "Belajar Golang"})

	result := service.FindById(ctx, 1)
	assert.Equal(t, responseCreate.Id, result.Id)
	assert.Equal(t, responseCreate.Name, result.Name)

}

func TestFindByIdFailed(t *testing.T) {
	defer func() {
		error := recover()
		assert.NotNil(t, error)
	}()

	db := SetupNewDb()
	TruncateTodolist(db)
	defer db.Close()
	validate := validator.New()
	ctx := context.Background()

	tx, err := db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)
	service := injector.InitializedService(db, validate)
	responseCreate := service.Create(ctx, web.TodolistCreateRequest{Name: "Belajar Golang"})

	result := service.FindById(ctx, 2)
	assert.NotEqual(t, responseCreate.Id, result.Id)
	assert.NotEqual(t, responseCreate.Name, result.Name)

}

func TestFindAll(t *testing.T) {
	db := SetupNewDb()
	TruncateTodolist(db)
	defer db.Close()
	validate := validator.New()
	ctx := context.Background()

	tx, err := db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)
	service := injector.InitializedService(db, validate)

	requests := []web.TodolistCreateRequest{
		{
			Name: "Belajar Golang",
		},
		{
			Name: "Belajar PHP",
		},
	}

	for _, request := range requests {
		service.Create(ctx, request)
	}

	responses := service.FindAll(ctx)
	assert.Equal(t, requests[0].Name, responses[0].Name)
}
