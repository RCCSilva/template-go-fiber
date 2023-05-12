package domain

import (
	"database/sql"
	"errors"
	"fmt"
)

func (app App) CreateUser(dto *CreateUserRequest) (*CreateUserResponse, error) {
	err := validate(dto)
	if err != nil {
		return nil, err
	}

	user := new(user)
	user.fromCreateUserRequest(dto)

	createdUser, err := app.memberRepository.Create(user)
	if err != nil {
		return nil, err
	}

	return createdUser.toCreateUserResponse(), nil
}

func (app App) GetUser(userId int) (*CreateUserResponse, error) {
	user, err := app.memberRepository.Get(userId)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, &ResultError{StatusCode: 404, Message: fmt.Sprintf("user (id: %d) not found", userId)}
		}
		return nil, err
	}

	return user.toCreateUserResponse(), nil
}
