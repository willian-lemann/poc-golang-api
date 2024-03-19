package usecases

import "github.com/willian-lemann/poc-golang-api/internal/repositories"

type GetUserUsecase struct {
	usersRepo repositories.UsersRepository
}

func NewGetUser(usersRepo repositories.UsersRepository) *GetUserUsecase {
	return &GetUserUsecase{
		usersRepo,
	}
}

func (u *GetUserUsecase) Execute() {
	u.usersRepo.Get()
}
