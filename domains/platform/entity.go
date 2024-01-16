package platform

import (
	"time"
)

type Entity struct {
	UUID         string                  `json:"uuid" bson:"_id,omitempty"`
	Name         string                  `json:"name" bson:"name"`
	Slug         string                  `json:"slug" bson:"slug"`
	Regexp       string                  `json:"regexp" bson:"regexp"`
	Prefix       string                  `json:"prefix" bson:"prefix"`
	Translations map[Locale]Translations `json:"translations" bson:"translations"`
	IsActive     bool                    `json:"isActive" bson:"is_active"`
	IsDeleted    bool                    `json:"isDeleted" bson:"is_deleted"`
	UpdatedAt    time.Time               `json:"updatedAt" bson:"updated_at"`
	CreatedAt    time.Time               `json:"createdAt" bson:"created_at"`
}

type Translations struct {
	Name        string `json:"name"`
	Placeholder string `json:"placeholder"`
	Description string `json:"description"`
}

type Locale string

const (
	LocaleEN Locale = "en"
	LocaleTR Locale = "tr"
)

func (l Locale) String() string {
	return string(l)
}
