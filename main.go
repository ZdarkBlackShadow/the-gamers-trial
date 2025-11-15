package main

import (
	"html/template"
	"log"

	"github.com/ZdarkBlackShadow/the-gamers-trial/config"
	"github.com/ZdarkBlackShadow/the-gamers-trial/controller"
	"github.com/ZdarkBlackShadow/the-gamers-trial/routes"
	"github.com/ZdarkBlackShadow/the-gamers-trial/service"
	"github.com/gofiber/fiber/v2"
	"github.com/subosito/gotenv"
)

func main() {
	app := fiber.New()

	err := gotenv.Load()
	if err != nil {
		log.Fatal("error when load .env:" + err.Error())
	}

	db, err := config.InitDatabase()
	if err != nil {
		log.Fatal("error when init database:" + err.Error())
	}

	app.Static("/static", "./assets")

	template, err := template.ParseGlob("./views/*html")
	if err != nil {
		log.Fatal(err)
	}
	
	authentificationService := service.InitAuthentificationService(db)

	homeController := controller.InitHomeController(template, authentificationService)
	questionsController := controller.InitQuestionsController(template, authentificationService)
	scoreController := controller.InitScoreController(template, authentificationService)

	homeRoutes := app.Group("/home")
	routes.RegisterDocumentTypesRoutes(homeRoutes, homeController)

	questionsRoutes := app.Group("/questions")
	routes.RegisterQuestionsRoutes(questionsRoutes, questionsController)

	scoreRoutes := app.Group("/score")
	routes.RegisterScoreRoutes(scoreRoutes, scoreController)

	log.Fatal(app.Listen(":8080"))
}
