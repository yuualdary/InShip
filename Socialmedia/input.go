package Socialmedia

type CreateSocialInput struct {
	SocType string `json:"socialmedia_type" binding:"required"`
	SocLink string `json:"socialmedia_link" binding:"required"`
}
