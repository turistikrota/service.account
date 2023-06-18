package account

import (
	"strconv"
	"time"

	"github.com/mixarchitecture/i18np"
)

type Factory struct {
	Errors Errors
	minAge int
	maxAge int
}

func NewFactory() Factory {
	return Factory{
		Errors: newAccountErrors(),
		minAge: 13,
		maxAge: 100,
	}
}

func (f Factory) IsZero() bool {
	return f.Errors == nil
}

func (f Factory) NewAccount(userUUID string, username string, usercode string) *Entity {
	t := time.Now()
	e := &Entity{
		UserUUID:      userUUID,
		UserName:      username,
		UserCode:      usercode,
		IsActive:      false,
		CompletedRate: 0,
		IsDeleted:     false,
		IsVerified:    false,
		CreatedAt:     &t,
		UpdatedAt:     &t,
	}
	e.CompletedRate = f.CalcCompletedRate(e)
	return e
}

func (f Factory) FixCode(code string) string {
	for len(code) < 4 {
		code = "0" + code
	}
	return code
}

func (f Factory) CalcCompletedRate(e *Entity) int {
	var rate int
	denominatorCount := 7 // 7 field
	list := []string{e.UserName, e.FullName, e.Description}
	if e.BirthDate != nil && e.BirthDate.Year() > 0 {
		rate += 1
	}
	for _, v := range list {
		if v != "" {
			rate++
		}
	}
	if len(e.Social) > 0 {
		rate++
	}
	return rate * 100 / denominatorCount
}

func (f Factory) Validate(e *Entity) *i18np.Error {
	if err := f.validateUserCode(e.UserCode); err != nil {
		return err
	}
	if err := f.validateUserName(e.UserName); err != nil {
		return err
	}
	return nil
}

func (f Factory) validateUserName(username string) *i18np.Error {
	if username == "" {
		return f.Errors.UserNameRequired()
	}
	return nil
}

func (f Factory) validateUserCode(usercode string) *i18np.Error {
	code, err := strconv.Atoi(usercode)
	if err != nil {
		return f.Errors.UserCodeInvalid()
	}
	if code == 0 {
		return f.Errors.UserCodeRequired()
	}
	if code < 0 || code > 9999 {
		return f.Errors.UserCodeInvalid()
	}
	return nil
}

func (f Factory) ValidateMinAge(birthDate *time.Time) *i18np.Error {
	if birthDate == nil {
		return nil
	}
	userAge := time.Now().Year() - birthDate.Year()
	if userAge < f.minAge {
		return f.Errors.MinAge(f.minAge)
	}
	if userAge == f.minAge && time.Now().Month() < birthDate.Month() {
		return f.Errors.MinAge(f.minAge)
	}
	if userAge == f.minAge && time.Now().Month() == birthDate.Month() && time.Now().Day() < birthDate.Day() {
		return f.Errors.MinAge(f.minAge)
	}
	if userAge > 150 {
		return f.Errors.MaxAge(f.maxAge)
	}
	return nil
}
