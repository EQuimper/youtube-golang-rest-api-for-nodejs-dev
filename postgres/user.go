package postgres

import (
	"errors"

	"github.com/go-pg/pg/v9"

	"todo/domain"
)

type UserRepo struct {
	DB *pg.DB
}

func (u *UserRepo) GetByEmail(email string) (*domain.User, error) {
	user := new(domain.User)

	err := u.DB.Model(user).Where("email = ?", email).First()
	if err != nil {
		if errors.Is(err, pg.ErrNoRows) {
			return nil, domain.ErrNoResult
		}

		return nil, err
	}

	return user, nil
}

func (u *UserRepo) GetByUsername(username string) (*domain.User, error) {
	user := new(domain.User)

	err := u.DB.Model(user).Where("username = ?", username).First()
	if err != nil {
		if errors.Is(err, pg.ErrNoRows) {
			return nil, domain.ErrNoResult
		}

		return nil, err
	}

	return user, nil
}

func (u *UserRepo) GetByID(id int64) (*domain.User, error) {
	user := new(domain.User)

	err := u.DB.Model(user).Where("id = ?", id).First()
	if err != nil {
		if errors.Is(err, pg.ErrNoRows) {
			return nil, domain.ErrNoResult
		}

		return nil, err
	}

	return user, nil
}

func (u *UserRepo) Create(user *domain.User) (*domain.User, error) {
	_, err := u.DB.Model(user).Returning("*").Insert()
	if err != nil {
		return nil, err
	}

	return user, nil
}

func NewUserRepo(DB *pg.DB) *UserRepo {
	return &UserRepo{DB: DB}
}
