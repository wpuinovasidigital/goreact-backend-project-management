package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/triadynata/project-management/config"
	"github.com/triadynata/project-management/controllers"
	"github.com/triadynata/project-management/database/seed"
	"github.com/triadynata/project-management/repositories"
	"github.com/triadynata/project-management/routes"
	"github.com/triadynata/project-management/services"
)

func main() {
	config.LoadEnv()
	config.ConnectDB()

	seed.SeedAdmin()
	app := fiber.New()

	userRepo := repositories.NewUserRepository()
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)

	routes.Setup(app, userController)

	port := config.AppConfig.AppPort
	log.Println("Server is running on port :", port)
	log.Fatal(app.Listen(":" + port))

}
