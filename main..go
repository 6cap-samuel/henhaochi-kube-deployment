package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"golang-clean-arch/configurations"
	"golang-clean-arch/controllers"
	"golang-clean-arch/exceptions"
	"golang-clean-arch/repositories/mongo"
	"golang-clean-arch/usecases"
)

func main() {
	configuration := configurations.New()
	database := configurations.NewMongoDatabase(configuration)

	postRepository := mongo.NewPostRepository(database)

	retrievePost := usecases.NewRetrievePostInteractor(&postRepository)
	createPost := usecases.NewCreatePostInput(&postRepository)

	postController := controllers.NewPostController(
		&retrievePost,
		&createPost,
	)

	app := fiber.New(configurations.NewFiberConfig())
	app.Use(recover.New())

	// Setup Routing
	postController.Route(app)

	err := app.Listen(":3000")
	exceptions.PanicIfNeeded(err)
}
