package photo

import commentDomain "hexagonal-fiber/domain/comment"

// ResponsePhotoComments is a struct that contains the response body for the photo comments
type ResponsePhotoComments struct {
	Photo
	Comments commentDomain.PaginationComment
}
