package repository

import (
	"context"
	"database/sql"

	"github.com/FeiraVed/todolist/model"
)

type TodolistRepository interface {
	Save(ctx context.Context, tx *sql.Tx, todolist model.Todolist) model.Todolist
	Update(ctx context.Context, tx *sql.Tx, todolist model.Todolist) model.Todolist
	Delete(ctx context.Context, tx *sql.Tx, todolistId int)
	FindById(ctx context.Context, tx *sql.Tx, todolistId int) (model.Todolist, error)
	FindAll(ctx context.Context, tx *sql.Tx) []model.Todolist
}
