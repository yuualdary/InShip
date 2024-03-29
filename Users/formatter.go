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

type DetailUserFormatter struct{
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Initial string `json:"initial"`
	Email      string `json:"email"`
	Profile_photo string `json:"image_url"`
	IsVerif bool `json:"isverif"`

}

func DetailUserFunc(user models.Users)DetailUserFormatter{

	DetailUserFormatter := DetailUserFormatter{}
	DetailUserFormatter.ID = int(user.ID)
	DetailUserFormatter.Name = user.Name
	DetailUserFormatter.Initial = user.Initial
	DetailUserFormatter.Email = user.Email
	DetailUserFormatter.Profile_photo = user.Profile_photo
	DetailUserFormatter.IsVerif = user.IsVerif

	return DetailUserFormatter
}