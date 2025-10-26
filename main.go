package main

import (
	"html/template"
	"log"
	"time"

	"github.com/ZdarkBlackShadow/the-gamers-trial/config"
	"github.com/ZdarkBlackShadow/the-gamers-trial/controller"
	"github.com/ZdarkBlackShadow/the-gamers-trial/middleware"
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

	config.InitLogger()

	gormLogger := config.GormLogger{}.LogMode(0)

	db, err := config.InitDatabase(gormLogger)
	if err != nil {
		config.Log.Fatal("error when init database:" + err.Error())
	}

	app.Static("/static", "./assets")

	template, err := template.ParseGlob("./views/*html")
	if err != nil {
		config.Log.Fatal("error when parsing templates:" + err.Error())
	}

	authentificationService := service.InitAuthentificationService(db)
	sessionService := service.InitSessionService()
	rateLimiter := middleware.NewRateLimiter(30, 1*time.Minute)

	homeController := controller.InitHomeController(template, authentificationService)
	authentificationController := controller.InitAuthentificationController(template, sessionService, authentificationService)

	app.Use(middleware.Logger())
	app.Use(rateLimiter.Handle())
	app.Use(middleware.SessionMiddleware(sessionService))

	homeRoutes := app.Group("/home")
	routes.RegisterDocumentTypesRoutes(homeRoutes, homeController)

	authRoutes := app.Group("/authentification")
	routes.RegisterAuthentificationRoutes(authRoutes, authentificationController)

	log.Fatal(app.Listen(":8080"))
}
