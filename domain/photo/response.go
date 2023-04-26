package photo

import commentDomain "hacktiv/final-project/domain/comment"

// ResponsePhotoComments is a struct that contains the response body for the photo comments
type ResponsePhotoComments struct {
	Photo
	Comments commentDomain.PaginationResultComment
}
