package usecases

import "github.com/willian-lemann/poc-golang-api/internal/repositories"

type Login struct {
	usersRepo repositories.UsersRepository
}

func NewLogin(usersRepo repositories.UsersRepository) *Login {
	return &Login{
		usersRepo,
	}
}

func (l *Login) Execute() {}
