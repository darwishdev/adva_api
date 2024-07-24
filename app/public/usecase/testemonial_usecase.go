package usecase

import (
	"context"

	"github.com/bufbuild/connect-go"
	rmsv1 "github.com/darwishdev/bzns_pro_api/common/pb/bznspro/v1"
)

func (u *PublicUsecase) TestemonialCreate(ctx context.Context, req *rmsv1.TestemonialCreateRequest) (*rmsv1.TestemonialCreateResponse, error) {
	if err := u.validator.Validate(req); err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}
	params := u.adapter.TestemonialCreateSqlFromGrpc(req)
	record, err := u.repo.TestemonialCreate(ctx, params)
	if err != nil {
		return nil, err
	}
	return u.adapter.TestemonialCreateGrpcFromSql(record), nil

}

func (u *PublicUsecase) TestemonialFindForUpdate(ctx context.Context, req *rmsv1.TestemonialFindForUpdateRequest) (*rmsv1.TestemonialUpdateRequest, error) {
	category, err := u.repo.TestemonialFindForUpdate(ctx, &req.TestemonialId)

	if err != nil {
		return nil, err
	}
	res := u.adapter.TestemonialFindForUpdateGrpcFromSql(category)

	return res, nil
}

func (s *PublicUsecase) TestemonialUpdate(ctx context.Context, req *rmsv1.TestemonialUpdateRequest) (*rmsv1.TestemonialUpdateResponse, error) {
	if err := s.validator.Validate(req); err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}
	params := s.adapter.TestemonialUpdateSqlFromGrpc(req)
	record, err := s.repo.TestemonialUpdate(ctx, params)
	if err != nil {
		return nil, err
	}
	return s.adapter.TestemonialUpdateGrpcFromSql(record), nil

}

func (s *PublicUsecase) TestemonialsList(ctx context.Context, req *rmsv1.TestemonialsListRequest) (*rmsv1.TestemonialsListResponse, error) {
	if err := s.validator.Validate(req); err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}
	record, err := s.repo.TestemonialsList(ctx)
	if err != nil {
		return nil, err
	}

	resp := s.adapter.TestemonialsListGrpcFromSql(record)

	// resp.Options = authorizedUser.GetAccessableActionsForGroup("categories")

	// time.Sleep(4 * time.Second)
	return resp, nil
}

func (s *PublicUsecase) TestemonialDeleteRestore(ctx context.Context, req *rmsv1.TestemonialDeleteRestoreRequest) (*rmsv1.TestemonialDeleteRestoreResponse, error) {
	err := s.repo.TestemonialDeleteRestore(ctx, req.TestemonialIds)
	if err != nil {
		return nil, err
	}
	return &rmsv1.TestemonialDeleteRestoreResponse{}, nil
}
