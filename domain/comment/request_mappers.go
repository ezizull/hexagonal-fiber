package comment

func (n *NewComment) ToDomainMapper() *Comment {
	return &Comment{
		UserID:  n.UserID,
		PhotoID: n.PhotoID,
		Message: n.Message,
	}
}

func (n *UpdateComment) ToDomainMapper() Comment {
	updateDomain := Comment{}

	if n.Message != nil {
		updateDomain.Message = *n.Message
	}

	return updateDomain
}
