package domain

import (
	"time"
)

type Todo struct {
	ID        int64  `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed" pg:",use_zero"`
	UserID    int64  `json:"userId"`

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type CreateTodoPayload struct {
	Title string `json:"title"`
}

func (c *CreateTodoPayload) IsValid() (bool, map[string]string) {
	v := NewValidator()

	v.MustBeNotEmpty("title", c.Title)
	v.MustBeLongerThan("title", c.Title, 3)

	return v.IsValid(), v.errors
}

func (d *Domain) CreateTodo(payload CreateTodoPayload, user *User) (*Todo, error) {
	data := &Todo{
		Title:     payload.Title,
		Completed: false,
		UserID:    user.ID,
	}

	todo, err := d.DB.TodoRepo.Create(data)
	if err != nil {
		return nil, err
	}

	return todo, nil
}

func (d *Domain) GetTodoByID(id int64) (*Todo, error) {
	todo, err := d.DB.TodoRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	return todo, nil
}

func (d *Domain) DeleteTodo(todo *Todo) error {
	err := d.DB.TodoRepo.Delete(todo)
	if err != nil {
		return err
	}

	return nil
}

type UpdateTodoPayload struct {
	Title     *string `json:"title"`
	Completed *bool   `json:"completed"`
}

func (u *UpdateTodoPayload) IsValid() (bool, map[string]string) {
	v := NewValidator()

	if *u.Title != "" {
		v.MustBeLongerThan("title", *u.Title, 3)
	}

	return v.IsValid(), v.errors
}

func (d *Domain) UpdateTodo(todo *Todo, payload UpdateTodoPayload) (*Todo, error) {
	didUpdate := false

	if *payload.Title != "" {
		todo.Title = *payload.Title
		didUpdate = true
	}

	if payload.Completed != nil {
		todo.Completed = *payload.Completed
		didUpdate = true
	}

	if didUpdate {
		todo.UpdatedAt = time.Now()
	}

	todo, err := d.DB.TodoRepo.Update(todo)
	if err != nil {
		return nil, err
	}

	return todo, nil
}

func (t *Todo) IsOwner(user *User) bool {
	return t.UserID == user.ID
}
