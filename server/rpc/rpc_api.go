package rpc

import (
	"context"

	"github.com/turistikrota/service.account/app/query"
	"github.com/turistikrota/service.account/domains/account"
	protos "github.com/turistikrota/service.account/protos/account"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s srv) GetAccountListByIds(ctx context.Context, req *protos.GetAccountListByIdsRequest) (*protos.AccountListByIdsResult, error) {
	users := make([]account.UserUnique, len(req.Users))
	for i, user := range req.Users {
		users[i] = account.UserUnique{
			UUID: user.Uuid,
			Name: user.Name,
		}
	}
	res, err := s.app.Queries.AccountListByIds(ctx, query.AccountListByIdsQuery{
		Users: users,
	})
	if err != nil {
		return nil, err
	}
	list := make([]*protos.Entity, len(res.Entities))
	for i, entity := range res.Entities {
		e := &protos.Entity{
			Uuid:          entity.UUID,
			UserUuid:      entity.UserUUID,
			UserName:      entity.UserName,
			FullName:      entity.FullName,
			Description:   entity.Description,
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
	return &protos.AccountListByIdsResult{
		Entities: list,
	}, nil
}

func (s srv) GetAccountListAsClaim(ctx context.Context, req *protos.AccountListAsClaimRequest) (*protos.AccountListAsClaimResult, error) {
	accounts, err := s.app.Queries.AccountListAsClaim(ctx, query.AccountListAsClaimQuery{
		UserUUID: req.UserId,
	})
	if err != nil {
		return nil, err
	}
	list := make([]*protos.Account, len(accounts.Entities))
	for i, dto := range accounts.Entities {
		list[i] = &protos.Account{
			Id:   dto.UUID,
			Name: dto.UserName,
		}
	}
	return &protos.AccountListAsClaimResult{
		Accounts: list,
	}, nil
}
