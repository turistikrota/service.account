package account

import (
	"time"

	"github.com/turistikrota/service.shared/helper"
)

type DetailDto struct {
	UUID          string     `json:"uuid"`
	UserName      string     `json:"userName"`
	AvatarURL     string     `json:"avatarUrl"`
	FullName      string     `json:"fullName"`
	Description   string     `json:"description"`
	IsActive      bool       `json:"isActive"`
	CompletedRate int        `json:"completedRate"`
	IsVerified    bool       `json:"isVerified"`
	BirthDate     *time.Time `json:"birthDate"`
	CreatedAt     *time.Time `json:"createdAt"`
	UpdatedAt     *time.Time `json:"updatedAt"`
}

type ListDto struct {
	UserName      string     `json:"userName"`
	FullName      string     `json:"fullName"`
	AvatarURL     string     `json:"avatarUrl"`
	Description   string     `json:"description"`
	IsActive      bool       `json:"isActive"`
	CompletedRate int        `json:"completedRate"`
	IsVerified    bool       `json:"isVerified"`
	CreatedAt     *time.Time `json:"createdAt"`
}

type ProfileDto struct {
	UserName    string     `json:"userName"`
	FullName    string     `json:"fullName"`
	AvatarURL   string     `json:"avatarUrl"`
	Description string     `json:"description"`
	IsVerified  bool       `json:"isVerified"`
	BirthDate   *time.Time `json:"birthDate"`
	CreatedAt   *time.Time `json:"createdAt"`
}

type AdminListDto struct {
	UserName      string     `json:"userName"`
	FullName      string     `json:"fullName"`
	AvatarURL     string     `json:"avatarUrl"`
	Description   string     `json:"description"`
	IsActive      bool       `json:"isActive"`
	CompletedRate int        `json:"completedRate"`
	IsVerified    bool       `json:"isVerified"`
	CreatedAt     *time.Time `json:"createdAt"`
}

func (e *Entity) ToDetail() DetailDto {
	return DetailDto{
		UUID:          e.UUID,
		UserName:      e.UserName,
		AvatarURL:     helper.CDN.DressAvatar(e.UserName),
		FullName:      e.FullName,
		Description:   e.Description,
		IsActive:      e.IsActive,
		CompletedRate: e.CompletedRate,
		IsVerified:    e.IsVerified,
		BirthDate:     e.BirthDate,
		CreatedAt:     e.CreatedAt,
		UpdatedAt:     e.UpdatedAt,
	}
}

func (e *Entity) ToList() ListDto {
	return ListDto{
		UserName:      e.UserName,
		FullName:      e.FullName,
		AvatarURL:     helper.CDN.DressAvatar(e.UserName),
		Description:   e.Description,
		IsActive:      e.IsActive,
		CompletedRate: e.CompletedRate,
		IsVerified:    e.IsVerified,
		CreatedAt:     e.CreatedAt,
	}
}

func (e *Entity) ToProfile() ProfileDto {
	return ProfileDto{
		UserName:    e.UserName,
		FullName:    e.FullName,
		AvatarURL:   helper.CDN.DressAvatar(e.UserName),
		Description: e.Description,
		IsVerified:  e.IsVerified,
		BirthDate:   e.BirthDate,
		CreatedAt:   e.CreatedAt,
	}
}

func (e *Entity) ToAdminList() AdminListDto {
	return AdminListDto{
		UserName:      e.UserName,
		FullName:      e.FullName,
		AvatarURL:     helper.CDN.DressAvatar(e.UserName),
		Description:   e.Description,
		IsActive:      e.IsActive,
		CompletedRate: e.CompletedRate,
		IsVerified:    e.IsVerified,
		CreatedAt:     e.CreatedAt,
	}
}
