package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v3"
	"github.com/willian-lemann/poc-golang-api/internal/usecases"
)

type UsersHandlers struct {
	GetUsersUseCase *usecases.GetUsersUsecase
	GetUserUseCase  *usecases.GetUserUsecase
}

func NewUsersHandlers(getUsersUseCase *usecases.GetUsersUsecase, getUserUseCase *usecases.GetUserUsecase) *UsersHandlers {
	return &UsersHandlers{
		GetUsersUseCase: getUsersUseCase,
		GetUserUseCase:  getUserUseCase,
	}
}

func (u *UsersHandlers) GetUsersHandler(c fiber.Ctx) error {
	output, err := u.GetUsersUseCase.Execute()
	if err != nil {
		c.Status(http.StatusBadRequest)
	}
	return c.JSON(output)
}

func (u *UsersHandlers) GetUserHandler(c fiber.Ctx) error {
	return c.Send(nil)
}
