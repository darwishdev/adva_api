package usecase

import (
	"context"

	"github.com/bufbuild/connect-go"
	rmsv1 "github.com/darwishdev/bzns_pro_api/common/pb/bznspro/v1"
)

func (u *PublicUsecase) ProgramCreate(ctx context.Context, req *rmsv1.ProgramCreateRequest) (*rmsv1.ProgramCreateResponse, error) {
	if err := u.validator.Validate(req); err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}
	params := u.adapter.ProgramCreateSqlFromGrpc(req)
	record, err := u.repo.ProgramCreate(ctx, params)
	if err != nil {
		return nil, err
	}
	return u.adapter.ProgramCreateGrpcFromSql(record), nil

}

func (u *PublicUsecase) ProgramFindForUpdate(ctx context.Context, req *rmsv1.ProgramFindForUpdateRequest) (*rmsv1.ProgramUpdateRequest, error) {
	category, err := u.repo.ProgramFindForUpdate(ctx, &req.ProgramId)

	if err != nil {
		return nil, err
	}
	res := u.adapter.ProgramFindForUpdateGrpcFromSql(category)

	return res, nil
}

func (s *PublicUsecase) ProgramUpdate(ctx context.Context, req *rmsv1.ProgramUpdateRequest) (*rmsv1.ProgramUpdateResponse, error) {
	if err := s.validator.Validate(req); err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}
	params := s.adapter.ProgramUpdateSqlFromGrpc(req)
	record, err := s.repo.ProgramUpdate(ctx, params)
	if err != nil {
		return nil, err
	}
	return s.adapter.ProgramUpdateGrpcFromSql(record), nil

}

func (s *PublicUsecase) ProgramsList(ctx context.Context, req *rmsv1.ProgramsListRequest) (*rmsv1.ProgramsListResponse, error) {
	if err := s.validator.Validate(req); err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}
	record, err := s.repo.ProgramsList(ctx)
	if err != nil {
		return nil, err
	}

	resp := s.adapter.ProgramsListGrpcFromSql(record)

	// resp.Options = authorizedUser.GetAccessableActionsForGroup("categories")

	// time.Sleep(4 * time.Second)
	return resp, nil
}

func (s *PublicUsecase) ProgramDeleteRestore(ctx context.Context, req *rmsv1.ProgramDeleteRestoreRequest) (*rmsv1.ProgramDeleteRestoreResponse, error) {
	err := s.repo.ProgramDeleteRestore(ctx, req.ProgramIds)
	if err != nil {
		return nil, err
	}
	return &rmsv1.ProgramDeleteRestoreResponse{}, nil
}
