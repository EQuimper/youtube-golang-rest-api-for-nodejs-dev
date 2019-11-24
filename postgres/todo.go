package postgres

import (
	"github.com/go-pg/pg/v9"

	"todo/domain"
)

type TodoRepo struct {
	DB *pg.DB
}

func (t *TodoRepo) Create(todo *domain.Todo) (*domain.Todo, error) {
	_, err := t.DB.Model(todo).Returning("*").Insert()
	if err != nil {
		return nil, err
	}

	return todo, nil
}

func NewTodoRepo(DB *pg.DB) *TodoRepo {
	return &TodoRepo{DB: DB}
}
