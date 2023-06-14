package account

import (
	"time"

	"api.turistikrota.com/account/src/config"
	"github.com/turistikrota/service.shared/events"
)

type Events interface {
	Deleted(user UserUnique)
	Created(user UserUnique)
	Disabled(user UserUnique)
	Enabled(user UserUnique)
	SocialAdded(user UserUnique, social EntitySocial)
	SocialUpdated(user UserUnique, social EntitySocial)
	SocialRemoved(user UserUnique, platform string)
	Updated(user UserUnique, entity Entity)
}

type (
	EventAccountDeleted struct {
		UserUUID    string `json:"user_uuid"`
		AccountName string `json:"name"`
		AccountCode string `json:"code"`
	}
	EventAccountCreated struct {
		UserUUID    string     `json:"user_uuid"`
		AccountName string     `json:"name"`
		AccountCode string     `json:"code"`
		CreatedAt   *time.Time `json:"created_at"`
	}
	EventAccountDisabled struct {
		UserUUID    string `json:"user_uuid"`
		AccountName string `json:"name"`
		AccountCode string `json:"code"`
	}
	EventAccountEnabled struct {
		UserUUID    string `json:"user_uuid"`
		AccountName string `json:"name"`
		AccountCode string `json:"code"`
	}
	EventAccountSocialAdded struct {
		UserUUID      string `json:"user_uuid"`
		AccountName   string `json:"name"`
		AccountCode   string `json:"code"`
		PlatformName  string `json:"platform_name"`
		PlatformValue string `json:"platform_value"`
	}
	EventAccountSocialUpdated struct {
		UserUUID      string `json:"user_uuid"`
		AccountName   string `json:"name"`
		AccountCode   string `json:"code"`
		PlatformName  string `json:"platform_name"`
		PlatformValue string `json:"platform_value"`
	}
	EventAccountSocialRemoved struct {
		UserUUID     string `json:"user_uuid"`
		AccountName  string `json:"name"`
		AccountCode  string `json:"code"`
		PlatformName string `json:"platform_name"`
	}
	EventAccountUpdated struct {
		UserUUID    string `json:"user_uuid"`
		AccountName string `json:"name"`
		AccountCode string `json:"code"`
		Entity      Entity `json:"entity"`
	}
)

type accountEvents struct {
	publisher events.Publisher
	topics    config.Topics
}

type EventConfig struct {
	Topics    config.Topics
	Publisher events.Publisher
}

func NewEvents(config EventConfig) Events {
	return &accountEvents{
		publisher: config.Publisher,
		topics:    config.Topics,
	}
}

func (e *accountEvents) Deleted(user UserUnique) {
	_ = e.publisher.Publish(e.topics.Account.Deleted, &EventAccountDeleted{
		UserUUID:    user.UUID,
		AccountName: user.Name,
		AccountCode: user.Code,
	})
}

func (e *accountEvents) Created(user UserUnique) {
	_ = e.publisher.Publish(e.topics.Account.Created, &EventAccountCreated{
		UserUUID:    user.UUID,
		AccountName: user.Name,
		AccountCode: user.Code,
	})
}

func (e *accountEvents) Disabled(user UserUnique) {
	_ = e.publisher.Publish(e.topics.Account.Disabled, &EventAccountDisabled{
		UserUUID:    user.UUID,
		AccountName: user.Name,
		AccountCode: user.Code,
	})
}

func (e *accountEvents) Enabled(user UserUnique) {
	_ = e.publisher.Publish(e.topics.Account.Enabled, &EventAccountEnabled{
		UserUUID:    user.UUID,
		AccountName: user.Name,
		AccountCode: user.Code,
	})
}

func (e *accountEvents) SocialAdded(user UserUnique, social EntitySocial) {
	_ = e.publisher.Publish(e.topics.Account.SocialAdded, &EventAccountSocialAdded{
		UserUUID:      user.UUID,
		AccountName:   user.Name,
		AccountCode:   user.Code,
		PlatformName:  social.Platform,
		PlatformValue: social.Value,
	})
}

func (e *accountEvents) SocialUpdated(user UserUnique, social EntitySocial) {
	_ = e.publisher.Publish(e.topics.Account.SocialUpdated, &EventAccountSocialUpdated{
		UserUUID:      user.UUID,
		AccountName:   user.Name,
		AccountCode:   user.Code,
		PlatformName:  social.Platform,
		PlatformValue: social.Value,
	})
}

func (e *accountEvents) SocialRemoved(user UserUnique, platform string) {
	_ = e.publisher.Publish(e.topics.Account.SocialRemoved, &EventAccountSocialRemoved{
		UserUUID:     user.UUID,
		AccountName:  user.Name,
		AccountCode:  user.Code,
		PlatformName: platform,
	})
}

func (e *accountEvents) Updated(user UserUnique, entity Entity) {
	_ = e.publisher.Publish(e.topics.Account.Updated, &EventAccountUpdated{
		UserUUID:    user.UUID,
		AccountName: user.Name,
		AccountCode: user.Code,
		Entity:      entity,
	})
}
