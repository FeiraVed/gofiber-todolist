package main

import (
	"github.com/FeiraVed/todolist/helper"
	"github.com/FeiraVed/todolist/injector"
)

func main() {

	app := injector.InitializedServer()
	err := app.Listen(":3000")
	helper.PanicIfError(err)
}
