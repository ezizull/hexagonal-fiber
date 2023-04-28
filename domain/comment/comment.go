package comment

import (
	"time"

	"github.com/google/uuid"
)

// Comment is a struct that contains the comment information
type Comment struct {
	ID        uuid.UUID  `json:"id" example:"cef47ee2-7211-452a-a087-79ce4b8ec3a3" gorm:"gorm:"type:uuid;default:uuid_generate_v4()"`
	UserID    string     `json:"user_id" gorm:"index"`
	PhotoID   string     `json:"photo_id" gorm:"index"`
	Message   string     `json:"message" example:"caption"`
	CreatedAt time.Time  `json:"created_at,omitempty" example:"2021-02-24 20:19:39" gorm:"autoCreateTime:mili"`
	UpdatedAt time.Time  `json:"updated_at,omitempty" example:"2021-02-24 20:19:39" gorm:"autoUpdateTime:mili"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" example:"null"`
}

// TableName overrides the table name used by Comment to `comments`
func (*Comment) TableName() string {
	return "comments"
}

// PaginationComment is a struct that contains the pagination result for comment
type PaginationComment struct {
	Data       *[]Comment
	Total      int64
	Limit      int64
	Current    int64
	NextCursor uint
	PrevCursor uint
	NumPages   int64
}
