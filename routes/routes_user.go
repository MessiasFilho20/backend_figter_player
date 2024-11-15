package routes

import (
	"github.com/MessiasFilho/backend_figter_player.git/user"
	"github.com/gofiber/fiber/v2"
)

func SetuserRoutes(app *fiber.App) {
	apipath := "/api"
	v1 := app.Group(apipath)

	v1.Post("/user", user.CreateUser)
}
