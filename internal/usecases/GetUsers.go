package usecases

import (
	"github.com/willian-lemann/poc-golang-api/internal/entities"
	"github.com/willian-lemann/poc-golang-api/internal/repositories"
)

type GetUsersUsecase struct {
	usersRepo repositories.UsersRepository
}

func NewGetUsers(usersRepo repositories.UsersRepository) *GetUsersUsecase {
	return &GetUsersUsecase{
		usersRepo,
	}
}

func (u *GetUsersUsecase) Execute() ([]entities.User, error) {
	users, err := u.usersRepo.ListAll()
	if err != nil {
		return nil, err
	}
	return users, nil
}
