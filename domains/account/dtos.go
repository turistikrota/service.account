package account

type DetailDto struct{}

func (e *Entity) ToDetail() DetailDto {
	return DetailDto{}
}
