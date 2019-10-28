package domain

type RegisterPayload struct {
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmPassword"`
	Username        string `json:"username"`
}

func (d *Domain) Register(payload RegisterPayload) (*User, error) {
	userExist, _ := d.DB.UserRepo.GetByEmail(payload.Email)
	if userExist != nil {
		return nil, ErrUserWithEmailAlreadyExist
	}

	userExist, _ = d.DB.UserRepo.GetByUsername(payload.Username)
	if userExist != nil {
		return nil, ErrUserWithUsernameAlreadyExist
	}

	password, err := d.setPassword(payload.Password)
	if err != nil {
		return nil, err
	}

	data := &User{
		Username: payload.Username,
		Email:    payload.Email,
		Password: *password,
	}

	user, err := d.DB.UserRepo.Create(data)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (d *Domain) setPassword(password string) (*string, error) {
	return nil, nil
}
