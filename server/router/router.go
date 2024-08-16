package router

import (
	"blog/controller"

	"github.com/gofiber/fiber/v2"
)

func SetUpRoutes(app *fiber.App) {

	app.Get("/", controller.BlogList)
	app.Post("/", controller.BlogCreate)
	app.Put("/:id", controller.BlogUpdate)
	app.Delete("/", controller.BlogDelete)
}
