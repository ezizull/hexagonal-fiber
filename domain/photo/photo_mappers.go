package photo

func ArrayToDomainMapper(photos *[]Photo) *[]Photo {
	booksDomain := make([]Photo, len(*photos))
	for i, photo := range *photos {
		booksDomain[i] = photo
	}

	return &booksDomain
}
