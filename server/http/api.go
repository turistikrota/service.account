package http

import (
	"time"

	"github.com/cilloparch/cillop/middlewares/i18n"
	"github.com/cilloparch/cillop/result"
	"github.com/gofiber/fiber/v2"
	"github.com/turistikrota/service.account/app/command"
	"github.com/turistikrota/service.account/app/query"
	"github.com/turistikrota/service.account/domains/account"
	"github.com/turistikrota/service.account/pkg/utils"
	"github.com/turistikrota/service.shared/server/http/auth/current_account"
	"github.com/turistikrota/service.shared/server/http/auth/current_user"
)

func (h srv) AccountDelete(ctx *fiber.Ctx) error {
	cmd := command.AccountDeleteCmd{}
	h.parseParams(ctx, &cmd)
	res, err := h.app.Commands.AccountDelete(ctx.Context(), cmd)
	if err != nil {
		l, a := i18n.GetLanguagesInContext(*h.i18n, ctx)
		return result.Error(h.i18n.TranslateFromError(*err, l, a))

	}
	return result.SuccessDetail(Messages.Success.Ok, res)
}

func (h srv) AccountRestore(ctx *fiber.Ctx) error {
	cmd := command.AccountRestoreCmd{}
	h.parseParams(ctx, &cmd)
	res, err := h.app.Commands.AccountRestore(ctx.Context(), cmd)
	if err != nil {
		l, a := i18n.GetLanguagesInContext(*h.i18n, ctx)
		return result.Error(h.i18n.TranslateFromError(*err, l, a))

	}
	return result.SuccessDetail(Messages.Success.Ok, res)
}

func (h srv) AccountCreate(ctx *fiber.Ctx) error {
	cmd := command.AccountCreateCmd{}
	h.parseBody(ctx, &cmd)
	cmd.UserUUID = current_user.Parse(ctx).UUID
	res, err := h.app.Commands.AccountCreate(ctx.Context(), cmd)
	if err != nil {
		l, a := i18n.GetLanguagesInContext(*h.i18n, ctx)
		return result.Error(h.i18n.TranslateFromError(*err, l, a))

	}
	return result.SuccessDetail(Messages.Success.Ok, res)
}

func (h srv) AccountEnable(ctx *fiber.Ctx) error {
	cmd := command.AccountEnableCmd{}
	cmd.UserName = current_account.Parse(ctx).Name
	cmd.UserUUID = current_user.Parse(ctx).UUID
	res, err := h.app.Commands.AccountEnable(ctx.Context(), cmd)
	if err != nil {
		l, a := i18n.GetLanguagesInContext(*h.i18n, ctx)
		return result.Error(h.i18n.TranslateFromError(*err, l, a))

	}
	return result.SuccessDetail(Messages.Success.Ok, res)
}

func (h srv) AccountDisable(ctx *fiber.Ctx) error {
	cmd := command.AccountDisableCmd{}
	cmd.UserName = current_account.Parse(ctx).Name
	cmd.UserUUID = current_user.Parse(ctx).UUID
	res, err := h.app.Commands.AccountDisable(ctx.Context(), cmd)
	if err != nil {
		l, a := i18n.GetLanguagesInContext(*h.i18n, ctx)
		return result.Error(h.i18n.TranslateFromError(*err, l, a))

	}
	return result.SuccessDetail(Messages.Success.Ok, res)
}

func (h srv) AccountUpdate(ctx *fiber.Ctx) error {
	cmd := command.AccountUpdateCmd{}
	h.parseBody(ctx, &cmd)
	cmd.UserUUID = current_user.Parse(ctx).UUID
	cmd.UserName = current_account.Parse(ctx).Name
	res, err := h.app.Commands.AccountUpdate(ctx.Context(), cmd)
	if err != nil {
		l, a := i18n.GetLanguagesInContext(*h.i18n, ctx)
		return result.Error(h.i18n.TranslateFromError(*err, l, a))

	}
	return result.SuccessDetail(Messages.Success.Ok, res)
}

func (h srv) AccountFilter(ctx *fiber.Ctx) error {
	p := utils.Pagination{}
	h.parseQuery(ctx, &p)
	query := query.AccountFilterQuery{}
	h.parseParams(ctx, &query)
	filter := account.FilterEntity{}
	h.parseQuery(ctx, &filter)
	query.Pagination = &p
	query.FilterEntity = &filter
	res, err := h.app.Queries.AccountFilter(ctx.UserContext(), query)
	if err != nil {
		l, a := i18n.GetLanguagesInContext(*h.i18n, ctx)
		return result.ErrorDetail(h.i18n.TranslateFromError(*err, l, a), res)
	}
	return result.SuccessDetail(Messages.Success.Ok, res.List)
}

func (h srv) AccountGetByName(ctx *fiber.Ctx) error {
	query := query.AccountGetByNameQuery{}
	h.parseParams(ctx, &query)
	res, err := h.app.Queries.AccountGetByName(ctx.UserContext(), query)
	if err != nil {
		l, a := i18n.GetLanguagesInContext(*h.i18n, ctx)
		return result.ErrorDetail(h.i18n.TranslateFromError(*err, l, a), res)
	}
	return result.SuccessDetail(Messages.Success.Ok, res.Dto)
}

func (h srv) AccountListByUser(ctx *fiber.Ctx) error {
	query := query.AccountListByUserQuery{}
	h.parseParams(ctx, &query)
	res, err := h.app.Queries.AccountListByUser(ctx.UserContext(), query)
	if err != nil {
		l, a := i18n.GetLanguagesInContext(*h.i18n, ctx)
		return result.ErrorDetail(h.i18n.TranslateFromError(*err, l, a), res)
	}
	return result.SuccessDetail(Messages.Success.Ok, res.Dtos)
}

func (h srv) AccountListMy(ctx *fiber.Ctx) error {
	query := query.AccountListMyQuery{}
	query.UserUUID = current_user.Parse(ctx).UUID
	res, err := h.app.Queries.AccountListMy(ctx.UserContext(), query)
	if err != nil {
		l, a := i18n.GetLanguagesInContext(*h.i18n, ctx)
		return result.ErrorDetail(h.i18n.TranslateFromError(*err, l, a), res)
	}
	return result.SuccessDetail(Messages.Success.Ok, res.Dtos)
}

func (h srv) AccountProfileView(ctx *fiber.Ctx) error {
	query := query.AccountProfileViewQuery{}
	h.parseParams(ctx, &query)
	res, err := h.app.Queries.AccountProfileView(ctx.UserContext(), query)
	if err != nil {
		l, a := i18n.GetLanguagesInContext(*h.i18n, ctx)
		return result.ErrorDetail(h.i18n.TranslateFromError(*err, l, a), err.Params)
	}
	return result.SuccessDetail(Messages.Success.Ok, res.Dto)
}

func (h srv) AccountSelect(ctx *fiber.Ctx) error {
	q := query.AccountDetailQuery{}
	h.parseParams(ctx, &q)
	query := query.AccountGetQuery{}
	query.UserName = q.UserName
	query.UserUUID = current_user.Parse(ctx).UUID
	res, err := h.app.Queries.AccountGet(ctx.UserContext(), query)
	if err != nil {
		l, a := i18n.GetLanguagesInContext(*h.i18n, ctx)
		return result.ErrorDetail(h.i18n.TranslateFromError(*err, l, a), res)
	}
	ctx.Cookie(h.CreateServerSideCookie(".s.a.u", res.Dto.UserName))
	return result.Success(Messages.Success.Ok)
}

func (h srv) AccountGetSelected(ctx *fiber.Ctx) error {
	userName := ctx.Cookies(".s.a.u")
	if userName == "" {
		return result.ErrorDetail(Messages.Error.RequiredAccountSelect, map[string]interface{}{
			"mustSelect": true,
		})
	}
	query := query.AccountGetQuery{}
	query.UserName = userName
	query.UserUUID = current_user.Parse(ctx).UUID
	res, err := h.app.Queries.AccountGet(ctx.UserContext(), query)
	if err != nil {
		l, a := i18n.GetLanguagesInContext(*h.i18n, ctx)
		return result.ErrorDetail(h.i18n.TranslateFromError(*err, l, a), res)
	}
	return result.SuccessDetail(Messages.Success.Ok, res.Dto)
}

func (h srv) CreateServerSideCookie(key string, value string) *fiber.Cookie {
	return &fiber.Cookie{
		Name:        key,
		Value:       value,
		Path:        "/",
		HTTPOnly:    true,
		Secure:      true,
		SameSite:    "Strict",
		Domain:      h.httpHeaders.Domain,
		MaxAge:      60 * 60 * 24 * 365,
		Expires:     time.Now().Add(60 * 60 * 24 * 365 * time.Second),
		SessionOnly: false,
	}
}
