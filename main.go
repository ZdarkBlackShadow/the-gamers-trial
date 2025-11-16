package main

import (
	"encoding/gob"
	"html/template"
	"log"
	"time"

	"github.com/ZdarkBlackShadow/the-gamers-trial/config"
	"github.com/ZdarkBlackShadow/the-gamers-trial/controller"
	"github.com/ZdarkBlackShadow/the-gamers-trial/middleware"
	"github.com/ZdarkBlackShadow/the-gamers-trial/model/entity"
	"github.com/ZdarkBlackShadow/the-gamers-trial/model/session"
	"github.com/ZdarkBlackShadow/the-gamers-trial/model/views"
	"github.com/ZdarkBlackShadow/the-gamers-trial/routes"
	"github.com/ZdarkBlackShadow/the-gamers-trial/service"
	"github.com/gofiber/fiber/v2"
	"github.com/subosito/gotenv"
)

func main() {
	gob.Register(views.Error{})
	gob.Register(entity.User{})
	gob.Register(session.UserAnswer{})
	app := fiber.New()

	err := gotenv.Load()
	if err != nil {
		log.Fatal("error when load .env:" + err.Error())
	}

	config.InitLogger()
	config.InitRandomGenerator()

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
	questionService := service.InitQuestionService(db)
	imageService := service.InitImageService(db, "./public")

	rateLimiter := middleware.NewRateLimiter(30, 1*time.Minute)

	homeController := controller.InitHomeController(template, sessionService)
	authentificationController := controller.InitAuthentificationController(template, sessionService, authentificationService)
	questionsController := controller.InitQuestionController(questionService, sessionService, template)
	scoreController := controller.InitScoreController(template, authentificationService)
	imageController := controller.InitImageController(imageService)

	app.Use(middleware.Logger())
	app.Use(rateLimiter.Handle())
	app.Use(middleware.SessionMiddleware(sessionService))

	homeRoutes := app.Group("")
	routes.RegisterHomeRoutes(homeRoutes, homeController)

	authRoutes := app.Group("/authentification")
	routes.RegisterAuthentificationRoutes(authRoutes, authentificationController)

	questionsRoutes := app.Group("/questions")
	routes.RegisterQuestionRoutes(questionsRoutes, questionsController)

	scoreRoutes := app.Group("/score")
	routes.RegisterScoreRoutes(scoreRoutes, scoreController)

	imageRoutes := app.Group("/image")
	routes.RegisterImageRoutes(imageRoutes, imageController)

	log.Fatal(app.Listen(":8080"))
}
