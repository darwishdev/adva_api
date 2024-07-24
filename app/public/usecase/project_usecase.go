package usecase

import (
	"context"

	"github.com/bufbuild/connect-go"
	rmsv1 "github.com/darwishdev/bzns_pro_api/common/pb/bznspro/v1"
)

func (u *PublicUsecase) ProjectCreate(ctx context.Context, req *rmsv1.ProjectCreateRequest) (*rmsv1.ProjectCreateResponse, error) {
	if err := u.validator.Validate(req); err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}
	params := u.adapter.ProjectCreateSqlFromGrpc(req)
	record, err := u.repo.ProjectCreate(ctx, params)
	if err != nil {
		return nil, err
	}
	return u.adapter.ProjectCreateGrpcFromSql(record), nil

}

func (u *PublicUsecase) ProjectFindForUpdate(ctx context.Context, req *rmsv1.ProjectFindForUpdateRequest) (*rmsv1.ProjectUpdateRequest, error) {
	category, err := u.repo.ProjectFindForUpdate(ctx, &req.ProjectId)

	if err != nil {
		return nil, err
	}
	res := u.adapter.ProjectFindForUpdateGrpcFromSql(category)

	return res, nil
}

func (s *PublicUsecase) ProjectUpdate(ctx context.Context, req *rmsv1.ProjectUpdateRequest) (*rmsv1.ProjectUpdateResponse, error) {
	if err := s.validator.Validate(req); err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}
	params := s.adapter.ProjectUpdateSqlFromGrpc(req)
	record, err := s.repo.ProjectUpdate(ctx, params)
	if err != nil {
		return nil, err
	}
	return s.adapter.ProjectUpdateGrpcFromSql(record), nil

}

func (s *PublicUsecase) ProjectsList(ctx context.Context, req *rmsv1.ProjectsListRequest) (*rmsv1.ProjectsListResponse, error) {
	if err := s.validator.Validate(req); err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}
	record, err := s.repo.ProjectsList(ctx)
	if err != nil {
		return nil, err
	}

	resp := s.adapter.ProjectsListGrpcFromSql(record)

	// resp.Options = authorizedUser.GetAccessableActionsForGroup("categories")

	// time.Sleep(4 * time.Second)
	return resp, nil
}

func (s *PublicUsecase) ProjectDeleteRestore(ctx context.Context, req *rmsv1.ProjectDeleteRestoreRequest) (*rmsv1.ProjectDeleteRestoreResponse, error) {
	err := s.repo.ProjectDeleteRestore(ctx, req.ProjectIds)
	if err != nil {
		return nil, err
	}
	return &rmsv1.ProjectDeleteRestoreResponse{}, nil
}
