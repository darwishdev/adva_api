package adapter

import (
	db "github.com/darwishdev/bzns_pro_api/common/db/gen"
	rmsv1 "github.com/darwishdev/bzns_pro_api/common/pb/bznspro/v1"
)

type BlogsAdapterInterface interface {
	blogsListRowGrpcFromSql(resp *db.BlogsListRow) *rmsv1.BlogsListRow
	BlogsListGrpcFromSql(resp *[]db.BlogsListRow) *rmsv1.BlogsListResponse
	BlogCreateSqlFromGrpc(req *rmsv1.BlogCreateRequest) *db.BlogCreateParams
	BlogUpdateSqlFromGrpc(req *rmsv1.BlogUpdateRequest) *db.BlogUpdateParams
	BlogFindForUpdateGrpcFromSql(resp *db.BlogFindForUpdateRow) *rmsv1.BlogUpdateRequest
}

type BlogsAdapter struct {
	dateFormat string
	round      func(value float32, places int) float32
}

func NewBlogsAdapter() BlogsAdapterInterface {
	return &BlogsAdapter{
		dateFormat: "2006-01-02 15:04:05",
	}
}
