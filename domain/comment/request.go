package comment

// GetComment is a struct that contains the data for new social media
type GetComment struct {
	PhotoID int `json:"photo_id" gorm:"index" binding:"required"`
}

// NewComment is a struct that contains the data for new social media
type NewComment struct {
	UserID  int    `json:"user_id" gorm:"index" binding:"-"`
	PhotoID int    `json:"photo_id" gorm:"index" binding:"required"`
	Message string `json:"message" example:"message" binding:"required"`
}

// UpdateComment is a struct that contains the data for update social media
type UpdateComment struct {
	Message *string `json:"message,omitempty" example:"message" binding:"-"`
}
