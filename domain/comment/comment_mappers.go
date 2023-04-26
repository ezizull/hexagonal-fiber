package comment

func ArrayToDomainMapper(comments *[]Comment) *[]Comment {
	booksDomain := make([]Comment, len(*comments))
	for i, comment := range *comments {
		booksDomain[i] = comment
	}

	return &booksDomain
}
