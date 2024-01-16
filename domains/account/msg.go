package account

type messages struct {
	AccountUserNameRequired string
	AccountAlreadyExist     string
	AccountInvalidDate      string
	AccountMinAge           string
	AccountMaxAge           string
	AccountFailed           string
	AccountNotFound         string
}

var I18nMessages = messages{
	AccountUserNameRequired: "error_account_user_name_required",
	AccountAlreadyExist:     "error_account_already_exist",
	AccountInvalidDate:      "error_account_invalid_date",
	AccountMinAge:           "error_account_min_age",
	AccountMaxAge:           "error_account_max_age",
	AccountFailed:           "error_account_failed",
	AccountNotFound:         "error_account_not_found",
}
