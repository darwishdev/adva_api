package usecase

import (
	"context"

	"github.com/bufbuild/protovalidate-go"
	"github.com/darwishdev/bzns_pro_api/app/blog/adapter"
	"github.com/darwishdev/bzns_pro_api/app/blog/repo"
	db "github.com/darwishdev/bzns_pro_api/common/db/gen"
	rmsv1 "github.com/darwishdev/bzns_pro_api/common/pb/bznspro/v1"
)

type BlogsUsecaseInterface interface {
	BlogCreate(ctx context.Context, req *rmsv1.BlogCreateRequest) (*rmsv1.BlogCreateResponse, error)
	BlogUpdate(ctx context.Context, req *rmsv1.BlogUpdateRequest) (*rmsv1.BlogUpdateResponse, error)
	BlogsList(ctx context.Context, req *rmsv1.BlogsListRequest) (*rmsv1.BlogsListResponse, error)
	BlogDeleteRestore(ctx context.Context, req *rmsv1.BlogDeleteRestoreRequest) (*rmsv1.BlogDeleteRestoreResponse, error)
	BlogFindForUpdate(ctx context.Context, req *rmsv1.BlogFindForUpdateRequest) (*rmsv1.BlogUpdateRequest, error)
}

type BlogsUsecase struct {
	repo      repo.BlogsRepoInterface
	validator *protovalidate.Validator
	adapter   adapter.BlogsAdapterInterface
}

func NewBlogsUsecase(store db.Store, validator *protovalidate.Validator) BlogsUsecaseInterface {
	repo := repo.NewBlogsRepo(store)
	adapter := adapter.NewBlogsAdapter()

	return &BlogsUsecase{
		repo:      repo,
		validator: validator,
		adapter:   adapter,
	}
}
