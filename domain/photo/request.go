package photo

// NewPhoto is a struct that contains the data for new photo
type NewPhoto struct {
	Title    string `json:"title" example:"title" binding:"required"`
	Caption  string `json:"caption,omitempty" example:"caption" binding:"-"`
	PhotoUrl string `json:"photo_url" example:"www.photo.com" binding:"required"`
	UserID   int    `json:"user_id" gorm:"index" binding:"-"`
}

// UpdatePhoto is a struct that contains the data for update photo
type UpdatePhoto struct {
	Title    *string `json:"title,omitempty" example:"title" binding:"-"`
	Caption  *string `json:"caption,omitempty,omitempty" example:"caption" binding:"-"`
	PhotoUrl *string `json:"photo_url,omitempty" example:"www.photo.com" binding:"-"`
}
