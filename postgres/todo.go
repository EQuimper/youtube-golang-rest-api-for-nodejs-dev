package postgres

import (
	"github.com/go-pg/pg/v9"

	"todo/domain"
)

type TodoRepo struct {
	DB *pg.DB
}

func (t *TodoRepo) GetByID(id int64) (*domain.Todo, error) {
	todo := new(domain.Todo)
	err := t.DB.Model(todo).Where("id = ?", id).First()
	if err != nil {
		return nil, err
	}

	return todo, err
}

func (t *TodoRepo) Update(todo *domain.Todo) (*domain.Todo, error) {
	_, err := t.DB.Model(todo).Where("id = ?", todo.ID).Returning("*").Update()
	if err != nil {
		return nil, err
	}

	return todo, nil
}

func (t *TodoRepo) Delete(todo *domain.Todo) error {
	err := t.DB.Delete(todo)
	if err != nil {
		return err
	}

	return nil
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
