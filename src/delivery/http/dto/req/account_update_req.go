package req

import (
	"time"

	"api.turistikrota.com/account/src/app/command"
	"github.com/turistikrota/service.shared/formats"
)

type AccountUpdateRequest struct {
	UserName    string
	UserCode    string
	AccountName string `json:"accountName,omitempty" validate:"omitempty,max=20,min=3"`
	AccountCode string `json:"accountCode,omitempty" validate:"omitempty,numeric,gt=0,lte=9999"`
	FullName    string `json:"fullName" validate:"required,max=70,min=3"`
	Description string `json:"description" validate:"required,max=1000"`
	BirthDate   string `json:"birthDate" validate:"required,datetime=2006-01-02"`
}

func (r *AccountUpdateRequest) LoadDetail(detail *AccountDetailRequest) {
	detail.Parse()
	r.UserName = detail.UserName
	r.UserCode = detail.UserCode
}

func (r *AccountUpdateRequest) ToCommand(userUUID string) command.AccountUpdateCommand {
	birth, _ := time.Parse(formats.DateYYYYMMDD, r.BirthDate)
	return command.AccountUpdateCommand{
		UserUUID:    userUUID,
		CurrentName: r.UserName,
		CurrentCode: r.UserCode,
		UserName:    r.AccountName,
		UserCode:    r.AccountCode,
		FullName:    r.FullName,
		Description: r.Description,
		BirthDate:   &birth,
	}
}
