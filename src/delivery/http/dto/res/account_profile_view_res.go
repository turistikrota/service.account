package res

import (
	"time"

	"api.turistikrota.com/account/src/app/query"
	"api.turistikrota.com/account/src/domain/account"
	"github.com/turistikrota/service.shared/helper"
)

type AccountProfileViewResponse struct {
	UserName    string                         `json:"userName"`
	UserCode    string                         `json:"userCode"`
	FullName    string                         `json:"fullName"`
	AvatarURL   string                         `json:"avatarUrl"`
	Description string                         `json:"description"`
	Social      []AccountProfileSocialResponse `json:"social"`
	IsVerified  bool                           `json:"isVerified"`
	CreatedAt   *time.Time                     `json:"createdAt"`
}

type AccountProfileSocialResponse struct {
	Platform string `json:"platform"`
	Url      string `json:"url"`
}

func (r *response) AccountProfileView(res *query.AccountProfileViewResult) *AccountProfileViewResponse {
	return &AccountProfileViewResponse{
		UserName:    res.Entity.UserName,
		UserCode:    res.Entity.UserCode,
		FullName:    res.Entity.FullName,
		Description: res.Entity.Description,
		Social:      r.AccountProfileSocialResponse(res.Entity.Social),
		AvatarURL:   helper.CDN.DressUser(res.Entity.UserName, res.Entity.UserCode),
		IsVerified:  res.Entity.IsVerified,
		CreatedAt:   res.Entity.CreatedAt,
	}
}

func (r *response) AccountProfileSocialResponse(social []account.EntitySocial) []AccountProfileSocialResponse {
	res := make([]AccountProfileSocialResponse, 0)
	for _, item := range social {
		res = append(res, AccountProfileSocialResponse{
			Platform: item.Platform,
			Url:      item.FixedValue,
		})
	}
	return res
}
