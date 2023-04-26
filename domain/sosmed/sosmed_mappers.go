package sosmed

func ArrayToDomainMapper(sosmeds *[]SocialMedia) *[]SocialMedia {
	booksDomain := make([]SocialMedia, len(*sosmeds))
	for i, sosmed := range *sosmeds {
		booksDomain[i] = sosmed
	}

	return &booksDomain
}
