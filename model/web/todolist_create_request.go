package web

type TodolistCreateRequest struct {
	Name string `validate:"required"`
}
