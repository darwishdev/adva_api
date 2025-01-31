package usecase

import (
	"context"

	"github.com/bufbuild/protovalidate-go"
	"github.com/darwishdev/bzns_pro_api/app/events/adapter"
	"github.com/darwishdev/bzns_pro_api/app/events/repo"
	"github.com/darwishdev/bzns_pro_api/common/auth"
	db "github.com/darwishdev/bzns_pro_api/common/db/gen"
	rmsv1 "github.com/darwishdev/bzns_pro_api/common/pb/bznspro/v1"
)

type EventsUsecaseInterface interface {
	EventCreate(ctx context.Context, req *rmsv1.EventCreateRequest) (*rmsv1.EventCreateResponse, error)
	EventUpdate(ctx context.Context, req *rmsv1.EventUpdateRequest) (*rmsv1.EventUpdateResponse, error)
	EventsList(ctx context.Context, req *rmsv1.EventsListRequest) (*rmsv1.EventsListResponse, error)
	EventDeleteRestore(ctx context.Context, req *rmsv1.EventDeleteRestoreRequest) (*rmsv1.EventDeleteRestoreResponse, error)
	EventFindForUpdate(ctx context.Context, req *rmsv1.EventFindForUpdateRequest) (*rmsv1.EventUpdateRequest, error)
	EventRequestCreate(ctx context.Context, req *rmsv1.EventRequestCreateRequest) (*rmsv1.EventRequestCreateResponse, error)
	EventRequestUpdate(ctx context.Context, req *rmsv1.EventRequestUpdateRequest) (*rmsv1.EventRequestUpdateResponse, error)
	EventRequestsList(ctx context.Context, req *rmsv1.EventRequestsListRequest, authorizedUser *auth.Payload) (*rmsv1.EventRequestsListResponse, error)
	EventRequestDeleteRestore(ctx context.Context, req *rmsv1.EventRequestDeleteRestoreRequest) (*rmsv1.EventRequestDeleteRestoreResponse, error)
}

type EventsUsecase struct {
	repo      repo.EventsRepoInterface
	validator *protovalidate.Validator
	adapter   adapter.EventsAdapterInterface
}

func NewEventsUsecase(store db.Store, validator *protovalidate.Validator) EventsUsecaseInterface {
	repo := repo.NewEventsRepo(store)
	adapter := adapter.NewEventsAdapter()

	return &EventsUsecase{
		repo:      repo,
		validator: validator,
		adapter:   adapter,
	}
}
