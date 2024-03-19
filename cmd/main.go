package main

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/willian-lemann/poc-golang-api/internal/infra"
	"github.com/willian-lemann/poc-golang-api/internal/router"
)

func main() {
	db := infra.CreateDBConnection()

	app := fiber.New()
	app.Use(logger.New())
	router.InitializeRoutes(app, db)
	app.Listen(":3333")
}
