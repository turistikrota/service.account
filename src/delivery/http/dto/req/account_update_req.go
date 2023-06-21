package req

import (
	"time"

	"api.turistikrota.com/account/src/app/command"
	"github.com/turistikrota/service.shared/formats"
)

type AccountUpdateRequest struct {
	UserName    string
	UserCode    string
	AccountName string `json:"accountName,omitempty" validate:"omitempty,username,max=20,min=3"`
	FullName    string `json:"fullName" validate:"required,max=70,min=3"`
	Description string `json:"description" validate:"required,max=1000"`
	BirthDate   string `json:"birthDate" validate:"required,datetime=2006-01-02"`
}

func (r *AccountUpdateRequest) LoadDetail(detail *AccountDetailRequest) {
	r.UserName = detail.UserName
}

func (r *AccountUpdateRequest) ToCommand(userUUID string) command.AccountUpdateCommand {
	birth, _ := time.Parse(formats.DateYYYYMMDD, r.BirthDate)
	return command.AccountUpdateCommand{
		UserUUID:    userUUID,
		CurrentName: r.UserName,
		UserName:    r.AccountName,
		FullName:    r.FullName,
		Description: r.Description,
		BirthDate:   &birth,
	}
}
