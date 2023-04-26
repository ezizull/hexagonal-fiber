package sosmed

// NewSocialMedia is a struct that contains the data for new social media
type NewSocialMedia struct {
	Name           string `json:"name" example:"name" binding:"required"`
	SocialMediaUrl string `json:"social_media_url" example:"www.sosmed.com" binding:"required"`
	UserID         int    `json:"user_id" gorm:"index" binding:"-"`
}

// UpdateSocialMedia is a struct that contains the data for update social media
type UpdateSocialMedia struct {
	Name           *string `json:"name,omitempty" example:"name" binding:"-"`
	SocialMediaUrl *string `json:"social_media_url,omitempty" example:"www.sosmed.com" binding:"-"`
}
