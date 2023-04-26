package comment

import (
	"time"
)

// Comment is a struct that contains the comment information
type Comment struct {
	ID        int        `json:"id" example:"1099" gorm:"primaryKey"`
	UserID    int        `json:"user_id" gorm:"index"`
	PhotoID   int        `json:"photo_id" gorm:"index"`
	Message   string     `json:"message" example:"caption"`
	CreatedAt time.Time  `json:"created_at,omitempty" example:"2021-02-24 20:19:39" gorm:"autoCreateTime:mili"`
	UpdatedAt time.Time  `json:"updated_at,omitempty" example:"2021-02-24 20:19:39" gorm:"autoUpdateTime:mili"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" example:"null"`
}

// TableName overrides the table name used by Comment to `comments`
func (*Comment) TableName() string {
	return "comments"
}

// PaginationResultComment is a struct that contains the pagination result for comment
type PaginationResultComment struct {
	Data       *[]Comment
	Total      int64
	Limit      int64
	Current    int64
	NextCursor uint
	PrevCursor uint
	NumPages   int64
}
