package sosmed

// NewSocialMedia is a struct that contains the data for new social media
type NewSocialMedia struct {
	Name           string `json:"name" example:"name" validate:"required"`
	SocialMediaUrl string `json:"social_media_url" example:"www.sosmed.com" validate:"required"`
	UserID         string `json:"user_id" gorm:"index" validate:"-"`
}

// UpdateSocialMedia is a struct that contains the data for update social media
type UpdateSocialMedia struct {
	Name           *string `json:"name,omitempty" example:"name" validate:"-"`
	SocialMediaUrl *string `json:"social_media_url,omitempty" example:"www.sosmed.com" validate:"-"`
}
