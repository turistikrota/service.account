package account

type DetailDto struct{}

type ListDto struct{}

type ProfileDto struct{}

func (e *Entity) ToDetail() DetailDto {
	return DetailDto{}
}

func (e *Entity) ToList() ListDto {
	return ListDto{}
}

func (e *Entity) ToProfile() ProfileDto {
	return ProfileDto{}
}
