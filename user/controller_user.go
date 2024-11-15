package user

import (
	"strconv"

	"github.com/MessiasFilho/backend_figter_player.git/database"
	"github.com/MessiasFilho/backend_figter_player.git/schemas"
	"github.com/gofiber/fiber/v2"
)

func CreateUser(c *fiber.Ctx) error {
	r := DtoUser{}

	if err := c.BodyParser(&r); err != nil {
		return c.Status(400).SendString("erro requisição")
	}

	if err := r.UserValidate(); err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	user := schemas.User{
		Name:  r.Name,
		Email: r.Email,
	}
	if err := database.DB.Create(&user).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"Error": "error ao criar user",
		})
	}
	return nil

}

func GetUserByID(c *fiber.Ctx) error {

	var user schemas.User

	id, err := strconv.Atoi(c.Params("Id"))

	if err != nil {
		return c.Status(400).SendString("construct not fund")
	}

	if err := database.DB.First(&user, id).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"Error": "error ao buscar usurario"})
	}
	return nil
}
