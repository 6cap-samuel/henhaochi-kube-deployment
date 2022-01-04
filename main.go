package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"golang-clean-arch/configurations"
	"golang-clean-arch/controllers"
	"golang-clean-arch/exceptions"
	"golang-clean-arch/repositories/mongo"
	"golang-clean-arch/usecases"
	"os"
)

func main() {
	configuration := configurations.New()
	database := configurations.NewMongoDatabase(configuration)

	postRepository := mongo.NewPostRepository(database)
	hashtagRepository := mongo.NewHashtagRepository(database)

	retrievePost := usecases.NewRetrievePostInteractor(&postRepository)
	createPost := usecases.NewCreatePostInput(
		&postRepository,
		&hashtagRepository,
	)
	updatePost := usecases.NewUpdatePostIntractor(&postRepository)
	retrieveHashtag := usecases.NewRetrieveHashtagInteractor(
		&hashtagRepository,
	)
	detailsPost := usecases.NewRetrievePostDetailsInteractor(
		&postRepository,
	)

	postController := controllers.NewPostController(
		&retrievePost,
		&createPost,
		&updatePost,
		&detailsPost,
	)
	hashtagController := controllers.NewHashtagController(
		&retrieveHashtag,
	)

	app := fiber.New(configurations.NewFiberConfig())
	app.Use(recover.New())

	app.Use(cors.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: os.Getenv("CORS"),
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	app.Use(logger.New())
	app.Use(logger.New(logger.Config{
		Format:     "[${ip}]:${port} ${status} - ${method} ${path}\n",
		TimeFormat: "02-Jan-2006",
		TimeZone:   "Singapore",
	}))

	// Setup Routing
	postController.Route(app)
	hashtagController.Route(app)

	port := os.Getenv("PORT")
	err := app.Listen("0.0.0.0:" + port)

	exceptions.PanicIfNeeded(err)
}
