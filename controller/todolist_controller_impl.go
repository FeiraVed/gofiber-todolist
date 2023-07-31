package controller

import (
	"context"
	"net/http"
	"strconv"
	"strings"

	"github.com/FeiraVed/todolist/model/web"
	"github.com/FeiraVed/todolist/service"
	"github.com/gofiber/fiber/v2"
)

type TodolistControllerImpl struct {
	service service.TodolistService
}

func NewTodolistControllerImpl(service service.TodolistService) *TodolistControllerImpl {

	return &TodolistControllerImpl{
		service: service,
	}
}

func (controller *TodolistControllerImpl) Create(c *fiber.Ctx) (_ error) {
	ctx := context.Background()
	name := c.FormValue("name")
	name = strings.TrimSpace(name)
	request := web.TodolistCreateRequest{
		Name: name,
	}

	controller.service.Create(ctx, request)
	return c.Redirect("/", http.StatusMovedPermanently)
}

func (controller *TodolistControllerImpl) Update(c *fiber.Ctx) (_ error) {

	id := c.FormValue("id")
	name := c.FormValue("name")
	name = strings.TrimSpace(name)
	ctx := context.Background()
	i, _ := strconv.Atoi(id)
	request := web.TodolistUpdateRequest{

		Id:   i,
		Name: name,
	}
	controller.service.Update(ctx, request)
	return c.Redirect("/", http.StatusMovedPermanently)
}

func (controller *TodolistControllerImpl) Delete(c *fiber.Ctx) (_ error) {

	idForm := c.FormValue("id")
	id, _ := strconv.Atoi(idForm)
	ctx := context.Background()
	controller.service.Delete(ctx, id)
	return c.Redirect("/")
}

func (controller *TodolistControllerImpl) FindById(c *fiber.Ctx) (_ error) {
	id := c.Params("id")
	i, _ := strconv.Atoi(id)
	ctx := context.Background()
	response := controller.service.FindById(ctx, i)
	return c.Render("update", fiber.Map{
		"Title": "Update Todolist",
		"Id":    i,
		"Name":  response.Name,
	})
}
func (controller *TodolistControllerImpl) FindAll(c *fiber.Ctx) (_ error) {
	ctx := context.Background()
	response := controller.service.FindAll(ctx)

	return c.Render("index", fiber.Map{
		"Todolist": response,
	})
}
