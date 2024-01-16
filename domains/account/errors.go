package account

import "github.com/cilloparch/cillop/i18np"

type Errors interface {
	UserNameRequired() *i18np.Error
	ErrAlreadyExist() *i18np.Error
	ErrInvalidDate() *i18np.Error
	MinAge(age int) *i18np.Error
	MaxAge(age int) *i18np.Error
	Failed(action string) *i18np.Error
	NotFound() *i18np.Error
}

type accountErrors struct{}

func newAccountErrors() Errors {
	return &accountErrors{}
}

func (e *accountErrors) UserNameRequired() *i18np.Error {
	return i18np.NewError(I18nMessages.AccountUserNameRequired)
}

func (e *accountErrors) ErrAlreadyExist() *i18np.Error {
	return i18np.NewError(I18nMessages.AccountAlreadyExist)
}

func (e *accountErrors) ErrInvalidDate() *i18np.Error {
	return i18np.NewError(I18nMessages.AccountInvalidDate)
}

func (e *accountErrors) MinAge(age int) *i18np.Error {
	return i18np.NewError(I18nMessages.AccountMinAge, i18np.P{"Age": age})
}

func (e *accountErrors) MaxAge(age int) *i18np.Error {
	return i18np.NewError(I18nMessages.AccountMaxAge, i18np.P{"Age": age})
}

func (e *accountErrors) Failed(action string) *i18np.Error {
	return i18np.NewError(I18nMessages.AccountFailed, i18np.P{"Action": action})
}

func (e *accountErrors) NotFound() *i18np.Error {
	return i18np.NewError(I18nMessages.AccountNotFound)
}
