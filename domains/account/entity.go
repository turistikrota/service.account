package account

import "time"

type Entity struct {
	UUID          string         `json:"uuid" bson:"_id,omitempty"`
	UserUUID      string         `json:"userUuid" bson:"user_uuid"`
	UserName      string         `json:"userName" bson:"user_name"`
	FullName      string         `json:"fullName" bson:"full_name"`
	Description   string         `json:"description" bson:"description"`
	Social        []EntitySocial `json:"social" bson:"social"`
	IsActive      bool           `json:"isActive" bson:"is_active"`
	CompletedRate int            `json:"completedRate" bson:"completed_rate"`
	IsDeleted     bool           `json:"isDeleted" bson:"is_deleted"`
	IsVerified    bool           `json:"isVerified" bson:"is_verified"`
	BirthDate     *time.Time     `json:"birthDate" bson:"birth_date"`
	CreatedAt     *time.Time     `json:"createdAt" bson:"created_at"`
	UpdatedAt     *time.Time     `json:"updatedAt" bson:"updated_at"`
}

type EntitySocial struct {
	Platform   string `json:"platform"`
	Value      string `json:"value"`
	FixedValue string `json:"fixed_value"`
}
