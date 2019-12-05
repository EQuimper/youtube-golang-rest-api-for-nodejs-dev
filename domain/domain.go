package domain

type UserRepo interface {
	GetByID(id int64) (*User, error)
	GetByEmail(email string) (*User, error)
	GetByUsername(username string) (*User, error)
	Create(user *User) (*User, error)
}

type TodoRepo interface {
	GetByID(id int64) (*Todo, error)
	Create(todo *Todo) (*Todo, error)
	Update(todo *Todo) (*Todo, error)
	Delete(todo *Todo) error
}

type HaveOwner interface {
	IsOwner(user *User) bool
}

type DB struct {
	UserRepo UserRepo
	TodoRepo TodoRepo
}

type Domain struct {
	DB DB
}
