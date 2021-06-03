package Users

import "InShip/models"

type UserFormatter struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Initial string `json:"initial"`
	Email      string `json:"email"`
	Token      string `json:"token"`
	Profile_photo string `json:"image_url"`
}

func FormatUser(user models.Users, token string) UserFormatter {
	formatter := UserFormatter{
		ID:         int(user.ID),
		Name:       user.Name,
		Initial: user.Initial,
		Email:      user.Email,
		Token:      token,
		Profile_photo:  user.Profile_photo,
	}
	return formatter
}