package web

type TodolistUpdateRequest struct {
	Id   int    `validate:"required"`
	Name string `validate:"required"`
}
