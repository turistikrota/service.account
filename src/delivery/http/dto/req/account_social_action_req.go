package req

import "api.turistikrota.com/account/src/app/command"

type AccountSocialActionRequest struct {
	UserName string
	UserCode string
	Platform string
	Value    string `json:"value" validate:"required"`
}

func (r *AccountSocialActionRequest) LoadSocial(social *AccountSocialRequest) {
	social.Parse()
	r.UserName = social.UserName
	r.UserCode = social.UserCode
	r.Platform = social.Platform
}

func (r *AccountSocialActionRequest) ToAddCommand(userUUID string) command.AccountSocialAddCommand {
	return command.AccountSocialAddCommand{
		UserUUID:    userUUID,
		AccountName: r.UserName,
		AccountCode: r.UserCode,
		Platform:    r.Platform,
		Value:       r.Value,
	}
}

func (r *AccountSocialActionRequest) ToUpdateCommand(userUUID string) command.AccountSocialUpdateCommand {
	return command.AccountSocialUpdateCommand{
		UserUUID:    userUUID,
		AccountName: r.UserName,
		AccountCode: r.UserCode,
		Platform:    r.Platform,
		Value:       r.Value,
	}
}
