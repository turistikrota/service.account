package res

import (
	"time"

	"api.turistikrota.com/account/src/app/query"
	"github.com/turistikrota/service.shared/helper"
)

type AccountListMyResponse struct {
	UserName      string     `json:"userName"`
	UserCode      string     `json:"userCode"`
	FullName      string     `json:"fullName"`
	AvatarURL     string     `json:"avatarUrl"`
	Description   string     `json:"description"`
	IsActive      bool       `json:"isActive"`
	CompletedRate int        `json:"completedRate"`
	IsVerified    bool       `json:"isVerified"`
	CreatedAt     *time.Time `json:"createdAt"`
}

func (r *response) AccountListMy(res *query.AccountListMyResult) []*AccountListMyResponse {
	list := make([]*AccountListMyResponse, 0)
	for _, account := range res.Entities {
		list = append(list, &AccountListMyResponse{
			UserName:      account.UserName,
			UserCode:      account.UserCode,
			FullName:      account.FullName,
			Description:   account.Description,
			IsActive:      account.IsActive,
			AvatarURL:     helper.CDN.DressUser(account.UserName, account.UserCode),
			CompletedRate: account.CompletedRate,
			IsVerified:    account.IsVerified,
			CreatedAt:     account.CreatedAt,
		})
	}
	return list
}
