package res

import (
	"github.com/turistikrota/service.account/src/app/command"
	"github.com/turistikrota/service.account/src/app/query"
	"github.com/turistikrota/service.account/src/domain/account"
)

type Response interface {
	AccountCreate(res *command.AccountCreateResult) *AccountCreateResponse
	AccountGet(res *query.AccountGetResult) *AccountGetResponse
	AccountListMy(res *query.AccountListMyResult) []*AccountListMyResponse
	AccountProfileView(res *query.AccountProfileViewResult) *AccountProfileViewResponse
	PlatformGet(res *query.PlatformGetBySlugResult) *PlatformGetResponse
	PlatformList(res *query.PlatformListAllResult) []PlatformListResponse
	AccountGetSelectedOk(account account.Entity) *AccountListMyResponse
	AccountGetSelectedNotFound() *AccountSelectNotSelectedResponse
}

type response struct{}

func New() Response {
	return &response{}
}
