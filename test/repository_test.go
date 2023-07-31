package test

import (
	"context"
	"database/sql"
	"os"
	"testing"
	"time"

	"github.com/FeiraVed/todolist/model"
	"github.com/FeiraVed/todolist/repository"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func SetupNewDb() *sql.DB {
	errEnv := godotenv.Load("../.env")

	if errEnv != nil {
		panic(errEnv)
	}

	username := os.Getenv("username")
	password := os.Getenv("password")
	db_name := os.Getenv("db_name")
	db, err := sql.Open("mysql", username+":"+password+"@tcp(localhost:3306)/"+db_name+"_test")

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(10)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(10 * time.Minute)

	if err != nil {
		panic(err)
	}

	return db

}

func TruncateTodolist(db *sql.DB) {
	db.Exec("TRUNCATE todolist")
}

var todolistRepository = repository.New()

func TestRepositorySave(t *testing.T) {
	db := SetupNewDb()
	defer db.Close()
	TruncateTodolist(db)

	tx, err := db.Begin()
	assert.Nil(t, err)

	ctx := context.Background()
	todolist := model.Todolist{
		Name: "Belajar Golang",
	}

	response := todolistRepository.Save(ctx, tx, todolist)

	err2 := tx.Commit()
	assert.Nil(t, err2)
	assert.Equal(t, 1, response.Id)
	assert.Equal(t, todolist.Name, response.Name)
}

func TestRepositoryUpdate(t *testing.T) {
	db := SetupNewDb()
	defer db.Close()
	TruncateTodolist(db)

	tx, err := db.Begin()
	assert.Nil(t, err)

	ctx := context.Background()
	todolist := model.Todolist{
		Name: "Belajar Golang",
	}

	todolistRepository.Save(ctx, tx, todolist)
	todolist2 := model.Todolist{
		Id:   1,
		Name: "Belajar PHP",
	}
	response := todolistRepository.Update(ctx, tx, todolist2)

	err2 := tx.Commit()
	assert.Nil(t, err2)
	assert.NotEqual(t, todolist.Name, response.Name)
}

func TestRepositoryDelete(t *testing.T) {
	db := SetupNewDb()
	defer db.Close()
	TruncateTodolist(db)

	tx, err := db.Begin()
	assert.Nil(t, err)

	ctx := context.Background()
	todolist := model.Todolist{
		Name: "Belajar Golang",
	}

	todolistRepository.Save(ctx, tx, todolist)
	todolistRepository.Delete(ctx, tx, 1)
	_, err2 := todolistRepository.FindById(ctx, tx, 1)
	assert.NotNil(t, err2)

	tx.Commit()
}

func TestRepositoryFindByIdSuccess(t *testing.T) {
	db := SetupNewDb()
	defer db.Close()
	TruncateTodolist(db)

	tx, err := db.Begin()
	assert.Nil(t, err)

	ctx := context.Background()
	todolist := model.Todolist{
		Name: "Belajar Golang",
	}

	todolistRepository.Save(ctx, tx, todolist)
	result, err2 := todolistRepository.FindById(ctx, tx, 1)
	assert.Nil(t, err2)
	assert.Equal(t, todolist.Name, result.Name)
	tx.Commit()

}

func TestRepositoryFindByIdFailed(t *testing.T) {
	db := SetupNewDb()
	defer db.Close()
	TruncateTodolist(db)

	tx, err := db.Begin()
	assert.Nil(t, err)

	ctx := context.Background()
	todolist := model.Todolist{
		Name: "Belajar Golang",
	}

	todolistRepository.Save(ctx, tx, todolist)
	result, err2 := todolistRepository.FindById(ctx, tx, 4)
	assert.NotNil(t, err2)
	assert.NotEqual(t, todolist.Name, result.Name)
	tx.Commit()

}

func TestRepositoryFindAll(t *testing.T) {
	db := SetupNewDb()
	defer db.Close()
	TruncateTodolist(db)

	tx, err := db.Begin()
	assert.Nil(t, err)

	ctx := context.Background()
	todolist := model.Todolist{
		Name: "Belajar Golang",
	}

	response1 := todolistRepository.Save(ctx, tx, todolist)
	response2 := todolistRepository.Save(ctx, tx, model.Todolist{
		Name: "Belajar MySQL",
	})

	results := todolistRepository.FindAll(ctx, tx)
	assert.Equal(t, response1.Id, results[0].Id)
	assert.Equal(t, response2.Id, results[1].Id)
	assert.Equal(t, response1.Name, results[0].Name)
	assert.Equal(t, response2.Name, results[1].Name)
	tx.Commit()
}
