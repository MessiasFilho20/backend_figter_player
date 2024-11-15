package main

import (
	"log"

	"github.com/MessiasFilho/backend_figter_player.git/database"
	"github.com/MessiasFilho/backend_figter_player.git/routes"
	"github.com/MessiasFilho/backend_figter_player.git/schemas"
	"github.com/gofiber/fiber/v2"
)

func main() {
	database.ConectDB()

	app := fiber.New()

	if err := database.DB.AutoMigrate(&schemas.User{}); err != nil {
		log.Fatal("Error migrate", err)
	}

	routes.SetuserRoutes(app)
	log.Fatal(app.Listen(":3000"))
}
