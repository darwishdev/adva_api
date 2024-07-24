package usecase

import (
	"context"

	"github.com/bufbuild/connect-go"
	"github.com/darwishdev/bzns_pro_api/common/auth"
	rmsv1 "github.com/darwishdev/bzns_pro_api/common/pb/bznspro/v1"
)

func (u *EventsUsecase) EventRequestCreate(ctx context.Context, req *rmsv1.EventRequestCreateRequest) (*rmsv1.EventRequestCreateResponse, error) {
	if err := u.validator.Validate(req); err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}
	params := u.adapter.EventRequestCreateSqlFromGrpc(req)
	_, err := u.repo.EventRequestCreate(ctx, params)
	if err != nil {
		return nil, err
	}
	return &rmsv1.EventRequestCreateResponse{}, nil

}

func (s *EventsUsecase) EventRequestUpdate(ctx context.Context, req *rmsv1.EventRequestUpdateRequest) (*rmsv1.EventRequestUpdateResponse, error) {
	if err := s.validator.Validate(req); err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}
	params := s.adapter.EventRequestUpdateSqlFromGrpc(req)
	_, err := s.repo.EventRequestUpdate(ctx, params)
	if err != nil {
		return nil, err
	}
	return &rmsv1.EventRequestUpdateResponse{}, nil

}

func (s *EventsUsecase) EventRequestsList(ctx context.Context, req *rmsv1.EventRequestsListRequest, authorizedUser *auth.Payload) (*rmsv1.EventRequestsListResponse, error) {
	if err := s.validator.Validate(req); err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}
	record, err := s.repo.EventRequestsList(ctx)
	if err != nil {
		return nil, err
	}
	resp := s.adapter.EventRequestsListGrpcFromSql(record)
	return resp, nil
}

func (s *EventsUsecase) EventRequestDeleteRestore(ctx context.Context, req *rmsv1.EventRequestDeleteRestoreRequest) (*rmsv1.EventRequestDeleteRestoreResponse, error) {
	err := s.repo.EventRequestDeleteRestore(ctx, req.EventRequestIds)
	if err != nil {
		return nil, err
	}
	return &rmsv1.EventRequestDeleteRestoreResponse{}, nil
}
