package photo

func (n *NewPhoto) ToDomainMapper() *Photo {
	return &Photo{
		Title:    n.Title,
		UserID:   n.UserID,
		Caption:  n.Caption,
		PhotoUrl: n.PhotoUrl,
	}
}

func (n *UpdatePhoto) ToDomainMapper() Photo {
	updateDomain := Photo{}

	if n.Title != nil {
		updateDomain.Title = *n.Title
	}

	if n.Caption != nil {
		updateDomain.Caption = *n.Caption
	}

	if n.PhotoUrl != nil {
		updateDomain.PhotoUrl = *n.PhotoUrl
	}

	return updateDomain
}
