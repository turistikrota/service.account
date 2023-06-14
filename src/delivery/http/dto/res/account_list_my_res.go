package res

import (
	"api.turistikrota.com/account/src/app/query"
	"time"
)

type AccountListMyResponse struct {
	UserName      string     `json:"userName"`
	UserCode      string     `json:"userCode"`
	FullName      string     `json:"fullName"`
	AvatarURL     string     `json:"avatarUrl"`
	CoverURL      string     `json:"coverUrl"`
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
			AvatarURL:     account.AvatarURL,
			CoverURL:      account.CoverURL,
			Description:   account.Description,
			IsActive:      account.IsActive,
			CompletedRate: account.CompletedRate,
			IsVerified:    account.IsVerified,
			CreatedAt:     account.CreatedAt,
		})
	}
	return list
}
