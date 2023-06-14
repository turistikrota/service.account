package res

import (
	"api.turistikrota.com/account/src/app/query"
	"api.turistikrota.com/account/src/domain/account"
	"time"
)

type AccountGetResponse struct {
	UserName      string                     `json:"userName"`
	UserCode      string                     `json:"userCode"`
	FullName      string                     `json:"fullName"`
	AvatarURL     string                     `json:"avatarUrl"`
	CoverURL      string                     `json:"coverUrl"`
	Description   string                     `json:"description"`
	Social        []AccountGetResponseSocial `json:"social"`
	IsActive      bool                       `json:"isActive"`
	CompletedRate int                        `json:"completedRate"`
	IsVerified    bool                       `json:"isVerified"`
	BirthDate     *time.Time                 `json:"birthDate"`
	CreatedAt     *time.Time                 `json:"createdAt"`
	UpdatedAt     *time.Time                 `json:"updatedAt"`
}

type AccountGetResponseSocial struct {
	Platform string `json:"platform"`
	Url      string `json:"url"`
}

func (r *response) AccountGet(res *query.AccountGetResult) *AccountGetResponse {
	return &AccountGetResponse{
		UserName:      res.Entity.UserName,
		UserCode:      res.Entity.UserCode,
		FullName:      res.Entity.FullName,
		AvatarURL:     res.Entity.AvatarURL,
		CoverURL:      res.Entity.CoverURL,
		Description:   res.Entity.Description,
		Social:        r.AccountGetResponseSocial(res.Entity.Social),
		IsActive:      res.Entity.IsActive,
		CompletedRate: res.Entity.CompletedRate,
		IsVerified:    res.Entity.IsVerified,
		BirthDate:     res.Entity.BirthDate,
		CreatedAt:     res.Entity.CreatedAt,
		UpdatedAt:     res.Entity.UpdatedAt,
	}
}

func (r *response) AccountGetResponseSocial(social []account.EntitySocial) []AccountGetResponseSocial {
	res := make([]AccountGetResponseSocial, 0)
	for _, item := range social {
		res = append(res, AccountGetResponseSocial{
			Platform: item.Platform,
			Url:      item.FixedValue,
		})
	}
	return res
}
