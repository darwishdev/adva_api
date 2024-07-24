package usecase

import (
	"context"

	"github.com/bufbuild/connect-go"
	rmsv1 "github.com/darwishdev/bzns_pro_api/common/pb/bznspro/v1"
)

func (u *PublicUsecase) ServiceCreate(ctx context.Context, req *rmsv1.ServiceCreateRequest) (*rmsv1.ServiceCreateResponse, error) {
	if err := u.validator.Validate(req); err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}
	params := u.adapter.ServiceCreateSqlFromGrpc(req)
	record, err := u.repo.ServiceCreate(ctx, params)
	if err != nil {
		return nil, err
	}
	return u.adapter.ServiceCreateGrpcFromSql(record), nil

}

func (u *PublicUsecase) ServiceFindForUpdate(ctx context.Context, req *rmsv1.ServiceFindForUpdateRequest) (*rmsv1.ServiceUpdateRequest, error) {
	category, err := u.repo.ServiceFindForUpdate(ctx, &req.ServiceId)

	if err != nil {
		return nil, err
	}
	res := u.adapter.ServiceFindForUpdateGrpcFromSql(category)

	return res, nil
}

func (s *PublicUsecase) ServiceUpdate(ctx context.Context, req *rmsv1.ServiceUpdateRequest) (*rmsv1.ServiceUpdateResponse, error) {
	if err := s.validator.Validate(req); err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}
	params := s.adapter.ServiceUpdateSqlFromGrpc(req)
	record, err := s.repo.ServiceUpdate(ctx, params)
	if err != nil {
		return nil, err
	}
	return s.adapter.ServiceUpdateGrpcFromSql(record), nil

}

func (s *PublicUsecase) ServicesList(ctx context.Context, req *rmsv1.ServicesListRequest) (*rmsv1.ServicesListResponse, error) {
	if err := s.validator.Validate(req); err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}
	record, err := s.repo.ServicesList(ctx)
	if err != nil {
		return nil, err
	}

	resp := s.adapter.ServicesListGrpcFromSql(record)

	// resp.Options = authorizedUser.GetAccessableActionsForGroup("categories")

	// time.Sleep(4 * time.Second)
	return resp, nil
}

func (s *PublicUsecase) ServiceDeleteRestore(ctx context.Context, req *rmsv1.ServiceDeleteRestoreRequest) (*rmsv1.ServiceDeleteRestoreResponse, error) {
	err := s.repo.ServiceDeleteRestore(ctx, req.ServiceIds)
	if err != nil {
		return nil, err
	}
	return &rmsv1.ServiceDeleteRestoreResponse{}, nil
}
