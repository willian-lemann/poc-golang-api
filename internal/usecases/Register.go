package usecases

import "github.com/willian-lemann/poc-golang-api/internal/repositories"

type Register struct {
	usersRepo repositories.UsersRepository
}

func NewRegister(usersRepo repositories.UsersRepository) *Register {
	return &Register{
		usersRepo,
	}
}

func (usecase *Register) Execute() {}
