package usecase

import (
	"context"

	"github.com/bufbuild/connect-go"
	rmsv1 "github.com/darwishdev/bzns_pro_api/common/pb/bznspro/v1"
)

func (u *EventsUsecase) EventCreate(ctx context.Context, req *rmsv1.EventCreateRequest) (*rmsv1.EventCreateResponse, error) {
	if err := u.validator.Validate(req); err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}
	params := u.adapter.EventCreateSqlFromGrpc(req)
	_, err := u.repo.EventCreate(ctx, params)
	if err != nil {
		return nil, err
	}
	return &rmsv1.EventCreateResponse{}, nil

}

func (u *EventsUsecase) EventFindForUpdate(ctx context.Context, req *rmsv1.EventFindForUpdateRequest) (*rmsv1.EventUpdateRequest, error) {
	category, err := u.repo.EventFindForUpdate(ctx, &req.EventId)

	if err != nil {
		return nil, err
	}
	res := u.adapter.EventFindForUpdateGrpcFromSql(category)

	return res, nil
}

func (s *EventsUsecase) EventUpdate(ctx context.Context, req *rmsv1.EventUpdateRequest) (*rmsv1.EventUpdateResponse, error) {
	if err := s.validator.Validate(req); err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}
	params := s.adapter.EventUpdateSqlFromGrpc(req)
	_, err := s.repo.EventUpdate(ctx, params)
	if err != nil {
		return nil, err
	}
	return &rmsv1.EventUpdateResponse{}, nil

}

func (s *EventsUsecase) EventsList(ctx context.Context, req *rmsv1.EventsListRequest) (*rmsv1.EventsListResponse, error) {
	if err := s.validator.Validate(req); err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}
	record, err := s.repo.EventsList(ctx)
	if err != nil {
		return nil, err
	}

	resp := s.adapter.EventsListGrpcFromSql(record)
	return resp, nil
}

func (s *EventsUsecase) EventDeleteRestore(ctx context.Context, req *rmsv1.EventDeleteRestoreRequest) (*rmsv1.EventDeleteRestoreResponse, error) {
	err := s.repo.EventDeleteRestore(ctx, req.EventIds)
	if err != nil {
		return nil, err
	}
	return &rmsv1.EventDeleteRestoreResponse{}, nil
}
