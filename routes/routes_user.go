package routes

import (
	"github.com/MessiasFilho/backend_figter_player.git/character"
	"github.com/MessiasFilho/backend_figter_player.git/user"
	"github.com/gofiber/fiber/v2"
)

func RunRoutes(app *fiber.App) {
	setuserRoutes(app)
	setCharacterRoutes(app)
}

func setuserRoutes(app *fiber.App) {
	apipath := "/api"
	v1 := app.Group(apipath)
	v1.Post("/user", user.CreateUser)
}

func setCharacterRoutes(app *fiber.App) {
	apipath := "/api"
	v1 := app.Group(apipath)
	v1.Post("/character", character.CreateCharacter)
}
