package photo

import (
	"time"
)

// Photo is a struct that contains the photo information
type Photo struct {
	ID        int        `json:"id" example:"1099" gorm:"primaryKey"`
	Title     string     `json:"title" example:"title"`
	Caption   string     `json:"caption" example:"caption"`
	PhotoUrl  string     `json:"photo_url" example:"www.photo.com"`
	UserID    int        `json:"user_id" gorm:"index"`
	CreatedAt time.Time  `json:"created_at,omitempty" example:"2021-02-24 20:19:39" gorm:"autoCreateTime:mili"`
	UpdatedAt time.Time  `json:"updated_at,omitempty" example:"2021-02-24 20:19:39" gorm:"autoUpdateTime:mili"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" example:"null"`
}

// TableName overrides the table name used by Photo to `photos`
func (*Photo) TableName() string {
	return "photos"
}

// PaginationPhoto is a struct that contains the pagination result for photo
type PaginationPhoto struct {
	Data       *[]Photo
	Total      int64
	Limit      int64
	Current    int64
	NextCursor uint
	PrevCursor uint
	NumPages   int64
}
