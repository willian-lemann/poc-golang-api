package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v3"
	"github.com/willian-lemann/poc-golang-api/internal/usecases"
)

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

type AuthHandlers struct {
	LoginUsecase    *usecases.Login
	RegisterUsecase *usecases.Register
}

func NewAuthHandlers(loginUsecase *usecases.Login, registerUsecase *usecases.Register) *AuthHandlers {
	return &AuthHandlers{
		LoginUsecase:    loginUsecase,
		RegisterUsecase: registerUsecase,
	}
}

func (a *AuthHandlers) CreateLoginHandler(c fiber.Ctx) error {
	input := new(LoginInput)

	if err := c.Bind().Body(input); err != nil {
		return c.Status(http.StatusBadRequest).JSON(err)
	}

	return c.Status(http.StatusOK).JSON(input)
}

func (a *AuthHandlers) CreateRegisterHandler(c fiber.Ctx) error {
	input := new(RegisterInput)

	if err := c.Bind().Body(input); err != nil {
		return c.Status(http.StatusBadRequest).JSON(err)
	}

	return c.Status(http.StatusOK).JSON(input)
}
