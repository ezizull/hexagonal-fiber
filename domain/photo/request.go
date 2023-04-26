package photo

// NewPhoto is a struct that contains the data for new photo
type NewPhoto struct {
	Title    string `json:"title" example:"title" validate:"required"`
	Caption  string `json:"caption,omitempty" example:"caption" validate:"-"`
	PhotoUrl string `json:"photo_url" example:"www.photo.com" validate:"required"`
	UserID   int    `json:"user_id" gorm:"index" validate:"-"`
}

// UpdatePhoto is a struct that contains the data for update photo
type UpdatePhoto struct {
	Title    *string `json:"title,omitempty" example:"title" validate:"-"`
	Caption  *string `json:"caption,omitempty,omitempty" example:"caption" validate:"-"`
	PhotoUrl *string `json:"photo_url,omitempty" example:"www.photo.com" validate:"-"`
}
