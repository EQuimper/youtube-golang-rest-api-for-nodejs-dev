package domain

type UserRepo interface {
	GetByID(id int64) (*User, error)
	GetByEmail(email string) (*User, error)
	GetByUsername(username string) (*User, error)
	Create(user *User) (*User, error)
}

type TodoRepo interface {
	Create(todo *Todo) (*Todo, error)
}

type DB struct {
	UserRepo UserRepo
	TodoRepo TodoRepo
}

type Domain struct {
	DB DB
}
