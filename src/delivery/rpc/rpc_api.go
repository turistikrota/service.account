package rpc

import (
	"context"

	"github.com/turistikrota/service.account/protos/account"
	"github.com/turistikrota/service.account/src/app/query"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s Server) GetAccountListByIds(ctx context.Context, req *account.GetAccountListByIdsRequest) (*account.AccountListByIdsResult, error) {
	res, err := s.app.Queries.AccountListByIds.Handle(ctx, query.AccountListByIdsQuery{
		UUIDs: req.Uuids,
	})
	if err != nil {
		return nil, err
	}
	list := make([]*account.Entity, len(res.Entities))
	for i, entity := range res.Entities {
		social := make([]*account.EntitySocial, len(entity.Social))
		for j, socialEntity := range entity.Social {
			social[j] = &account.EntitySocial{
				Platform:   socialEntity.Platform,
				Value:      socialEntity.Value,
				FixedValue: socialEntity.FixedValue,
			}
		}
		e := &account.Entity{
			Uuid:          entity.UUID,
			UserUuid:      entity.UserUUID,
			UserName:      entity.UserName,
			FullName:      entity.FullName,
			Description:   entity.Description,
			Social:        social,
			IsActive:      entity.IsActive,
			CompletedRate: int32(entity.CompletedRate),
			IsDeleted:     entity.IsDeleted,
			IsVerified:    entity.IsVerified,
			BirthDate:     nil,
			CreatedAt:     timestamppb.New(*entity.CreatedAt),
			UpdatedAt:     timestamppb.New(*entity.UpdatedAt),
		}
		if entity.BirthDate != nil {
			e.BirthDate = timestamppb.New(*entity.BirthDate)
		}
		list[i] = e
	}
	return &account.AccountListByIdsResult{
		Entities: list,
	}, nil
}
