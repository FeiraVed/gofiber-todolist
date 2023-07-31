package app

import (
	"github.com/FeiraVed/todolist/controller"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/template/html/v2"
)

func NewRouter(controller controller.TodolistController) *fiber.App {
	engine := html.New("./view", ".html")
	app := fiber.New(fiber.Config{
		Views:       engine,
		ViewsLayout: "layouts/main",
		ErrorHandler: func(c *fiber.Ctx, err error) error {

			if err != nil {
				return c.Status(fiber.StatusInternalServerError).Render("500", fiber.Map{
					"Title":   fiber.StatusInternalServerError,
					"Message": err.Error(),
				})
			}
			return nil
		},
	})

	app.Use(recover.New())
	app.Static("/", "./public")

	app.Get("/", controller.FindAll)

	app.Post("/todolist", controller.Create)
	app.Post("/todolist/:id", controller.Delete)

	app.Get("/todolist/:id", controller.FindById)
	app.Post("/todolist/:id/update", controller.Update)

	return app
}
