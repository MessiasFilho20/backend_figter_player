package character

import (
	"errors"

	"github.com/MessiasFilho/backend_figter_player.git/database"
	"github.com/MessiasFilho/backend_figter_player.git/schemas"
)

func GetUserID(id int) (*schemas.User, error) {
	var user schemas.User
	if err := database.DB.First(&user, id).Error; err != nil {
		return nil, errors.New("user not found")
	}
	return &user, nil
}
