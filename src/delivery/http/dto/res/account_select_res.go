package res

import (
	"github.com/turistikrota/service.account/src/domain/account"
	"github.com/turistikrota/service.shared/helper"
)

type AccountSelectNotSelectedResponse struct {
	MustSelect bool `json:"mustSelect"`
}

func (r *response) AccountGetSelectedOk(account account.Entity) *AccountListMyResponse {
	return &AccountListMyResponse{
		UserName:      account.UserName,
		FullName:      account.FullName,
		Description:   account.Description,
		IsActive:      account.IsActive,
		AvatarURL:     helper.CDN.DressAvatar(account.UserName),
		CompletedRate: account.CompletedRate,
		IsVerified:    account.IsVerified,
		CreatedAt:     account.CreatedAt,
	}
}

func (r *response) AccountGetSelectedNotFound() *AccountSelectNotSelectedResponse {
	return &AccountSelectNotSelectedResponse{
		MustSelect: true,
	}
}
