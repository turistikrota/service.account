package req

import "github.com/turistikrota/service.account/src/app/query"

type PlatformListRequest struct{}

func (r *PlatformListRequest) ToQuery() query.PlatformListAllQuery {
	return query.PlatformListAllQuery{}
}
