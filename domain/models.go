package domain

type user struct {
	ID       int
	Email    string
	Username string
	Password string
}

func (u *user) fromCreateUserRequest(dto *CreateUserRequest) error {
	u.Username = dto.Username
	u.Email = dto.Email

	pwd, err := hashPassword(dto.Password)
	if err != nil {
		return err
	}

	u.Password = pwd

	return nil
}

func (u user) toCreateUserResponse() *CreateUserResponse {
	return &CreateUserResponse{
		ID:       u.ID,
		Username: u.Username,
		Email:    u.Email,
	}
}
