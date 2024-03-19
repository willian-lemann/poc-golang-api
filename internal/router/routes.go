package router

import (
	"database/sql"

	"github.com/gofiber/fiber/v3"
	"github.com/willian-lemann/poc-golang-api/internal/handlers"
	"github.com/willian-lemann/poc-golang-api/internal/repositories"
	"github.com/willian-lemann/poc-golang-api/internal/usecases"
)

func InitializeRoutes(app *fiber.App, dbConnection *sql.DB) {
	repo := repositories.NewUsersDatabaseRepository(dbConnection)
	getUserUseCase := usecases.NewGetUser(repo)
	getUsersUseCase := usecases.NewGetUsers(repo)
	loginUsecase := usecases.NewLogin(repo)
	registerUsecase := usecases.NewRegister(repo)

	usersHandlers := handlers.NewUsersHandlers(getUsersUseCase, getUserUseCase)
	authHandlers := handlers.NewAuthHandlers(loginUsecase, registerUsecase)

	v1 := app.Group("/api")
	auth := v1.Group("/auth")
	auth.Post("/auth/login", authHandlers.CreateLoginHandler)
	auth.Post("/auth/register", authHandlers.CreateRegisterHandler)

	v1.Get("/users", usersHandlers.GetUsersHandler)
	v1.Get("/users/{id}", usersHandlers.GetUserHandler)
}
