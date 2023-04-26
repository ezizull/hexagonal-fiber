package sosmed

func (n *NewSocialMedia) ToDomainMapper() *SocialMedia {
	return &SocialMedia{
		Name:           n.Name,
		UserID:         n.UserID,
		SocialMediaUrl: n.SocialMediaUrl,
	}
}

func (n *UpdateSocialMedia) ToDomainMapper() SocialMedia {
	updateDomain := SocialMedia{}

	if n.Name != nil {
		updateDomain.Name = *n.Name
	}

	if n.SocialMediaUrl != nil {
		updateDomain.SocialMediaUrl = *n.SocialMediaUrl
	}

	return updateDomain
}
