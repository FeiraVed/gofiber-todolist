package web

type TodolistResponse struct {
	Id   int    `validate:"required"`
	Name string `validate:"required"`
}
