package account

import (
	"time"

	"github.com/cilloparch/cillop/events"
	"github.com/turistikrota/service.account/config"
)

type Events interface {
	Deleted(user UserUnique)
	Created(user UserUnique)
	Disabled(user UserUnique)
	Enabled(user UserUnique)
	Updated(user UserUnique, entity Entity)
}

type (
	EventDeleted struct {
		UserUUID    string `json:"user_uuid"`
		AccountName string `json:"name"`
	}
	EventCreated struct {
		UserUUID    string     `json:"user_uuid"`
		AccountName string     `json:"name"`
		CreatedAt   *time.Time `json:"created_at"`
	}
	EventDisabled struct {
		UserUUID    string `json:"user_uuid"`
		AccountName string `json:"name"`
	}
	EventEnabled struct {
		UserUUID    string `json:"user_uuid"`
		AccountName string `json:"name"`
	}
	EventUpdated struct {
		UserUUID    string `json:"user_uuid"`
		AccountName string `json:"name"`
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
	_ = e.publisher.Publish(e.topics.Account.Deleted, &EventDeleted{
		UserUUID:    user.UUID,
		AccountName: user.Name,
	})
}

func (e *accountEvents) Created(user UserUnique) {
	_ = e.publisher.Publish(e.topics.Account.Created, &EventCreated{
		UserUUID:    user.UUID,
		AccountName: user.Name,
	})
}

func (e *accountEvents) Disabled(user UserUnique) {
	_ = e.publisher.Publish(e.topics.Account.Disabled, &EventDisabled{
		UserUUID:    user.UUID,
		AccountName: user.Name,
	})
}

func (e *accountEvents) Enabled(user UserUnique) {
	_ = e.publisher.Publish(e.topics.Account.Enabled, &EventEnabled{
		UserUUID:    user.UUID,
		AccountName: user.Name,
	})
}

func (e *accountEvents) Updated(user UserUnique, entity Entity) {
	_ = e.publisher.Publish(e.topics.Account.Updated, &EventUpdated{
		UserUUID:    user.UUID,
		AccountName: user.Name,
		Entity:      entity,
	})
}
