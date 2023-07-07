package req

import (
	"github.com/turistikrota/service.account/src/app/command"
	"github.com/turistikrota/service.account/src/domain/platform"
)

type PlatformTranslationRequest struct {
	Slug   string `param:"slug" validate:"required"`
	Locale string `param:"locale" validate:"required,locale,oneof=tr en"`
}

func (r *PlatformTranslationRequest) ToRemoveCommand(adminUUID string) command.PlatformTranslationRemoveCommand {
	return command.PlatformTranslationRemoveCommand{
		AdminUUID: adminUUID,
		Slug:      r.Slug,
		Locale:    platform.Locale(r.Locale),
	}
}
