package photo

import commentDomain "hacktiv/final-project/domain/comment"

// PhotoComment is a struct that contains role of user
type PhotoComment struct {
	Photo
	Comment []commentDomain.Comment `gorm:"foreignKey:PhotoID"`
}

// TableName overrides the table name used by User to `users`
func (*PhotoComment) TableName() string {
	return "photos"
}
