//go:build wireinject
// +build wireinject

package injector

import (
	"database/sql"

	"github.com/FeiraVed/todolist/app"
	"github.com/FeiraVed/todolist/controller"
	"github.com/FeiraVed/todolist/repository"
	"github.com/FeiraVed/todolist/service"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/wire"
)

var todolistSet = wire.NewSet(
	repository.New,
	wire.Bind(new(repository.TodolistRepository), new(*repository.TodolistRepositoryImpl)),
	service.NewTodolistServiceImpl,
	wire.Bind(new(service.TodolistService), new(*service.TodolistServiceImpl)),
	controller.NewTodolistControllerImpl,
	wire.Bind(new(controller.TodolistController), new(*controller.TodolistControllerImpl)),
)

var repositorySet = wire.NewSet(
	repository.New,
	wire.Bind(new(repository.TodolistRepository), new(*repository.TodolistRepositoryImpl)),
)

func InitializedService(db *sql.DB, validate *validator.Validate) *service.TodolistServiceImpl {
	wire.Build(repositorySet, service.NewTodolistServiceImpl)
	return nil
}

func InitializedServer() *fiber.App {
	wire.Build(
		app.NewDb,
		validator.New,
		todolistSet,
		app.NewRouter,
	)
	return nil
}
