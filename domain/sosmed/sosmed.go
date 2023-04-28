package sosmed

import (
	"time"

	"github.com/google/uuid"
)

// SocialMedia is a struct that contains the social media information
type SocialMedia struct {
	ID             uuid.UUID `json:"id" example:"cef47ee2-7211-452a-a087-79ce4b8ec3a3" gorm:"gorm:"type:uuid;default:uuid_generate_v4()"`
	Name           string    `json:"name" example:"caption"`
	SocialMediaUrl string    `json:"social_media_url" example:"www.sosmed.com"`
	UserID         string    `json:"user_id" gorm:"index"`
	CreatedAt      time.Time `json:"created_at,omitempty" example:"2021-02-24 20:19:39" gorm:"autoCreateTime:mili"`
	UpdatedAt      time.Time `json:"updated_at,omitempty" example:"2021-02-24 20:19:39" gorm:"autoUpdateTime:mili"`
	DeletedAt      time.Time `json:"deleted_at,omitempty" example:"2021-02-24 20:19:39"`
}

// TableName overrides the table name used by SocialMedia to `social_media`
func (*SocialMedia) TableName() string {
	return "social_media"
}

// PaginationSocialMedia is a struct that contains the pagination result for social media
type PaginationSocialMedia struct {
	Data       *[]SocialMedia
	Total      int64
	Limit      int64
	Current    int64
	NextCursor uint
	PrevCursor uint
	NumPages   int64
}
