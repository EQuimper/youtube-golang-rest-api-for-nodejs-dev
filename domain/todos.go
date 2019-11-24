package domain

import "time"

type Todo struct {
	ID        int64  `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
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
