package req

import (
	"time"

	"github.com/mixarchitecture/microp/formats"
	"github.com/turistikrota/service.account/src/app/command"
)

type AccountUpdateRequest struct {
	UserName    string
	FullName    string `json:"fullName" validate:"required,max=70,min=3"`
	Description string `json:"description" validate:"required,max=1000"`
	BirthDate   string `json:"birthDate" validate:"omitempty,datetime=2006-01-02"`
}

func (r *AccountUpdateRequest) LoadDetail(detail *AccountDetailRequest) {
	r.UserName = detail.UserName
}

func (r *AccountUpdateRequest) ToCommand(userUUID string) command.AccountUpdateCommand {
	date := time.Time{}
	if r.BirthDate != "" {
		date, _ = time.Parse(formats.DateYYYYMMDD, r.BirthDate)
	}
	return command.AccountUpdateCommand{
		UserUUID:    userUUID,
		UserName:    r.UserName,
		FullName:    r.FullName,
		Description: r.Description,
		BirthDate:   &date,
	}
}
