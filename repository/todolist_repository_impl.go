package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/FeiraVed/todolist/helper"
	"github.com/FeiraVed/todolist/model"
)

type TodolistRepositoryImpl struct {
}

func New() *TodolistRepositoryImpl {
	return &TodolistRepositoryImpl{}
}

func (repository *TodolistRepositoryImpl) Save(
	ctx context.Context,
	tx *sql.Tx,
	todolist model.Todolist,
) (_ model.Todolist) {

	sql := `INSERT INTO todolist(name) VALUES (?)`
	result, err := tx.ExecContext(ctx, sql, todolist.Name)
	helper.PanicIfError(err)

	id, err2 := result.LastInsertId()

	todolist.Id = int(id)
	helper.PanicIfError(err2)
	return todolist
}

func (repository *TodolistRepositoryImpl) Update(
	ctx context.Context,
	tx *sql.Tx,
	todolist model.Todolist,
) (_ model.Todolist) {
	sql := `UPDATE todolist SET name = ? WHERE id = ?`
	_, err := tx.ExecContext(ctx, sql, todolist.Name, todolist.Id)
	helper.PanicIfError(err)

	return todolist
}

func (repository *TodolistRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, todolistId int) {
	sql := `DELETE FROM todolist WHERE id = ?`
	_, err := tx.ExecContext(ctx, sql, todolistId)
	helper.PanicIfError(err)

}

func (repository *TodolistRepositoryImpl) FindById(
	ctx context.Context,
	tx *sql.Tx,
	todolistId int,
) (_ model.Todolist, _ error) {
	sql := `SELECT id,name FROM todolist WHERE id = ?`
	rows, err := tx.QueryContext(ctx, sql, todolistId)
	defer rows.Close()
	helper.PanicIfError(err)

	todolist := model.Todolist{}
	if rows.Next() {
		rows.Scan(&todolist.Id, &todolist.Name)

		return todolist, nil
	}

	return todolist, errors.New("Todolist not found")
}

func (repository *TodolistRepositoryImpl) FindAll(
	ctx context.Context,
	tx *sql.Tx,
) (_ []model.Todolist) {
	todolists := []model.Todolist{}

	sql := `SELECT id,name FROM todolist`
	rows, err := tx.QueryContext(ctx, sql)
	defer rows.Close()
	helper.PanicIfError(err)

	for rows.Next() {
		todolist := model.Todolist{}
		rows.Scan(&todolist.Id, &todolist.Name)
		todolists = append(todolists, todolist)
	}

	return todolists

}
