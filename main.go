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
	//user
	userRepo := repositories.NewUserRepository()
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)

	//board
	boardRepo := repositories.NewBoardRepository()
	boardMemberRepo := repositories.NewBoardMemberRepository()
	boardService := services.NewBoardService(boardRepo, userRepo, boardMemberRepo)
	boardController := controllers.NewBoardController(boardService)

	//list
	listPosRepo := repositories.NewListPositionRepository()
	listRepo := repositories.NewListRepository()
	listService := services.NewListService(listRepo, boardRepo, listPosRepo)
	listController := controllers.NewListController(listService)
	labelRepo := repositories.NewLabelRepository()
	//card
	cardRepo := repositories.NewCardRepository()
	cardService := services.NewCardService(cardRepo, listRepo, userRepo, labelRepo)
	cardController := controllers.NewCardController(cardService)

	routes.Setup(app, userController, boardController, listController, cardController)

	port := config.AppConfig.AppPort
	log.Println("Server is running on port :", port)
	log.Fatal(app.Listen(":" + port))

}
