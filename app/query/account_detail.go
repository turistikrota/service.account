package query

type AccountDetailQuery struct {
	UserName string `params:"userName" validate:"required,username"`
}
