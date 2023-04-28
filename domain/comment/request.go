package comment

// GetComment is a struct that contains the data for new social media
type GetComment struct {
	PhotoID string `json:"photo_id" gorm:"index" validate:"required"`
}

// NewComment is a struct that contains the data for new social media
type NewComment struct {
	UserID  string `json:"user_id" gorm:"index" validate:"-"`
	PhotoID string `json:"photo_id" gorm:"index" validate:"required"`
	Message string `json:"message" example:"message" validate:"required"`
}

// UpdateComment is a struct that contains the data for update social media
type UpdateComment struct {
	Message *string `json:"message,omitempty" example:"message" validate:"-"`
}
