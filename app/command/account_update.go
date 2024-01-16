package command

import (
	"context"
	"time"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.account/domains/account"
)

type AccountUpdateCmd struct {
	UserUUID    string `json:"-"`
	UserName    string `json:"-"`
	FullName    string `json:"fullName" validate:"required,max=70,min=3"`
	Description string `json:"description" validate:"required,max=1000"`
	BirthDate   string `json:"birthDate" validate:"omitempty,datetime=2006-01-02"`
}

type AccountUpdateRes struct{}

type AccountUpdateHandler cqrs.HandlerFunc[AccountUpdateCmd, *AccountUpdateRes]

func NewAccountUpdateHandler(factory account.Factory, repo account.Repo, events account.Events) AccountUpdateHandler {
	return func(ctx context.Context, cmd AccountUpdateCmd) (*AccountUpdateRes, *i18np.Error) {
		u := account.UserUnique{
			UUID: cmd.UserUUID,
			Name: cmd.UserName,
		}
		acc, err := repo.Get(ctx, u)
		if err != nil {
			return nil, err
		}
		var date *time.Time
		if cmd.BirthDate != "" {
			d, err := time.Parse("2006-01-02", cmd.BirthDate)
			if err != nil {
				return nil, factory.Errors.ErrInvalidDate()
			}
			date = &d
		}
		if acc.BirthDate != date && date != nil {
			ageErr := factory.ValidateMinAge(date)
			if ageErr != nil {
				return nil, ageErr
			}
			acc.BirthDate = date
		} else {
			acc.BirthDate = nil
		}
		acc.FullName = cmd.FullName
		acc.Description = cmd.Description
		acc.CompletedRate = factory.CalcCompletedRate(acc)
		t := time.Now()
		acc.UpdatedAt = &t
		err = repo.Update(ctx, u, acc)
		if err != nil {
			return nil, err
		}
		events.Updated(u, *acc)
		return &AccountUpdateRes{}, nil
	}
}
