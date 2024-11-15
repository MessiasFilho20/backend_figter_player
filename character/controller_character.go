package character

import (
	"github.com/MessiasFilho/backend_figter_player.git/database"
	"github.com/MessiasFilho/backend_figter_player.git/schemas"
	"github.com/gofiber/fiber/v2"
)

func CreateCharacter(c *fiber.Ctx) error {

	r := DtoCharacter{}

	if err := c.BodyParser(&r); err != nil {
		return c.Status(400).SendString("erro requisição")
	}

	if err := r.ValidCharacter(); err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	user, _ := GetUserID(int(r.UserID))

	character := schemas.Character{
		UserID: user.ID,
		Name:   r.Name,
		Sexo:   r.Sexo,
		Classe: r.Classe,
	}

	if err := database.DB.Create(&character).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"Error": "error ao criar Character",
		})
	}

	return nil
}
