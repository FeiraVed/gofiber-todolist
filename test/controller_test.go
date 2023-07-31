package test

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"regexp"
	"strings"
	"testing"

	"github.com/FeiraVed/todolist/controller"
	"github.com/FeiraVed/todolist/injector"
	"github.com/FeiraVed/todolist/model/web"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/template/html/v2"
	"github.com/stretchr/testify/assert"
)

func TestControllerCreateSucces(t *testing.T) {
	db := SetupNewDb()
	TruncateTodolist(db)
	defer db.Close()

	var todolistService = injector.InitializedService(db, validator.New())
	var todolistController = controller.NewTodolistControllerImpl(todolistService)

	app := fiber.New(fiber.Config{
		Views: html.New("../view", ".html"),
	})

	app.Post("/todolist", todolistController.Create)
	app.Get("/", todolistController.FindAll)
	body := strings.NewReader("name=Belajar+Golang")
	req := httptest.NewRequest("POST", "http://localhost:300/todolist", body)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r, err := app.Test(req, -1)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusMovedPermanently, r.StatusCode)
}

func TestControllerCreateFailed(t *testing.T) {
	db := SetupNewDb()
	TruncateTodolist(db)
	defer db.Close()

	var todolistService = injector.InitializedService(db, validator.New())
	var todolistController = controller.NewTodolistControllerImpl(todolistService)

	app := fiber.New(fiber.Config{
		Views: html.New("../view", ".html"),
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			if err != nil {
				return c.Status(500).SendString("Internal Server Error")
			}
			return nil
		},
	})

	app.Use(recover.New())
	app.Post("/todolist", todolistController.Create)
	req := httptest.NewRequest("POST", "http://localhost:300/todolist", nil)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r, err := app.Test(req, -1)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusInternalServerError, r.StatusCode)

}

func TestControllerUpdateSuccess(t *testing.T) {
	db := SetupNewDb()
	TruncateTodolist(db)
	defer db.Close()

	var todolistService = injector.InitializedService(db, validator.New())
	var todolistController = controller.NewTodolistControllerImpl(todolistService)

	app := fiber.New(fiber.Config{
		Views: html.New("../view", ".html"),
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			if err != nil {
				return c.Status(500).SendString(err.Error())
			}
			return nil
		},
	})

	app.Use(recover.New())

	todolistService.Create(context.Background(), web.TodolistCreateRequest{Name: "Belajar Golang"})
	app.Post("/todolist/:id/update", todolistController.Update)
	app.Get("/", todolistController.FindAll)

	request := httptest.NewRequest(
		"POST",
		"http://localhost:3000/todolist/1/update",
		strings.NewReader("id=1&name=Belajar PHP"),
	)
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	response, err := app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, 301, response.StatusCode)

	request = httptest.NewRequest("GET", "http://localhost:3000/", nil)
	response, err2 := app.Test(request)
	assert.Nil(t, err2)
	bodyByte, err3 := io.ReadAll(response.Body)
	assert.Nil(t, err3)
	assert.Regexp(t, regexp.MustCompile("Belajar PHP"), string(bodyByte))
}

func TestControllerFindByIdSuccess(t *testing.T) {
	db := SetupNewDb()
	TruncateTodolist(db)
	defer db.Close()

	var todolistService = injector.InitializedService(db, validator.New())
	var todolistController = controller.NewTodolistControllerImpl(todolistService)

	app := fiber.New(fiber.Config{
		Views: html.New("../view", ".html"),
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			if err != nil {
				return c.Status(500).SendString(err.Error())
			}
			return nil
		},
	})

	app.Use(recover.New())

	todolistService.Create(context.Background(), web.TodolistCreateRequest{Name: "Belajar Golang"})

	app.Get("/todolist/:id", todolistController.FindById)
	request := httptest.NewRequest("GET", "http://localhost:3000/todolist/1", nil)
	response, err := app.Test(request)
	assert.Nil(t, err)
	respByte, _ := io.ReadAll(response.Body)
	assert.Regexp(t, regexp.MustCompile("Belajar Golang"), string(respByte))

}

func TestControllerFindByIdFailed(t *testing.T) {
	db := SetupNewDb()
	TruncateTodolist(db)
	defer db.Close()

	var todolistService = injector.InitializedService(db, validator.New())
	var todolistController = controller.NewTodolistControllerImpl(todolistService)

	app := fiber.New(fiber.Config{
		Views: html.New("../view", ".html"),
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			if err != nil {
				return c.Status(500).SendString(err.Error())
			}
			return nil
		},
	})

	app.Use(recover.New())

	todolistService.Create(context.Background(), web.TodolistCreateRequest{Name: "Belajar Golang"})

	app.Get("/todolist/:id", todolistController.FindById)
	request := httptest.NewRequest("GET", "http://localhost:3000/todolist/2", nil)
	response, _ := app.Test(request)
	assert.Equal(t, 500, response.StatusCode)

}

func TestControllerFindAll(t *testing.T) {
	db := SetupNewDb()
	TruncateTodolist(db)
	defer db.Close()

	var todolistService = injector.InitializedService(db, validator.New())
	var todolistController = controller.NewTodolistControllerImpl(todolistService)

	app := fiber.New(fiber.Config{
		Views: html.New("../view", ".html"),
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			if err != nil {
				return c.Status(500).SendString(err.Error())
			}
			return nil
		},
	})

	app.Use(recover.New())

	todolistService.Create(context.Background(), web.TodolistCreateRequest{Name: "Belajar Golang"})

	app.Get("/", todolistController.FindAll)
	request := httptest.NewRequest("GET", "http://localhost:3000/", nil)
	response, _ := app.Test(request)
	responseByte, _ := io.ReadAll(response.Body)
	assert.Regexp(t, regexp.MustCompile("Belajar Golang"), string(responseByte))

}
