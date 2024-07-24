package usecase

import (
	"context"

	"github.com/bufbuild/connect-go"
	rmsv1 "github.com/darwishdev/bzns_pro_api/common/pb/bznspro/v1"
)

func (u *PublicUsecase) TeamMemberCreate(ctx context.Context, req *rmsv1.TeamMemberCreateRequest) (*rmsv1.TeamMemberCreateResponse, error) {
	if err := u.validator.Validate(req); err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}
	params := u.adapter.TeamMemberCreateSqlFromGrpc(req)
	record, err := u.repo.TeamMemberCreate(ctx, params)
	if err != nil {
		return nil, err
	}
	return u.adapter.TeamMemberCreateGrpcFromSql(record), nil

}

func (u *PublicUsecase) TeamMemberFindForUpdate(ctx context.Context, req *rmsv1.TeamMemberFindForUpdateRequest) (*rmsv1.TeamMemberUpdateRequest, error) {
	category, err := u.repo.TeamMemberFindForUpdate(ctx, &req.TeamMemberId)

	if err != nil {
		return nil, err
	}
	res := u.adapter.TeamMemberFindForUpdateGrpcFromSql(category)

	return res, nil
}

func (s *PublicUsecase) TeamMemberUpdate(ctx context.Context, req *rmsv1.TeamMemberUpdateRequest) (*rmsv1.TeamMemberUpdateResponse, error) {
	if err := s.validator.Validate(req); err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}
	params := s.adapter.TeamMemberUpdateSqlFromGrpc(req)
	record, err := s.repo.TeamMemberUpdate(ctx, params)
	if err != nil {
		return nil, err
	}
	return s.adapter.TeamMemberUpdateGrpcFromSql(record), nil

}

func (s *PublicUsecase) TeamMembersList(ctx context.Context, req *rmsv1.TeamMembersListRequest) (*rmsv1.TeamMembersListResponse, error) {
	if err := s.validator.Validate(req); err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}
	record, err := s.repo.TeamMembersList(ctx)
	if err != nil {
		return nil, err
	}

	resp := s.adapter.TeamMembersListGrpcFromSql(record)

	// resp.Options = authorizedUser.GetAccessableActionsForGroup("categories")

	// time.Sleep(4 * time.Second)
	return resp, nil
}

func (s *PublicUsecase) TeamMemberDeleteRestore(ctx context.Context, req *rmsv1.TeamMemberDeleteRestoreRequest) (*rmsv1.TeamMemberDeleteRestoreResponse, error) {
	err := s.repo.TeamMemberDeleteRestore(ctx, req.TeamMemberIds)
	if err != nil {
		return nil, err
	}
	return &rmsv1.TeamMemberDeleteRestoreResponse{}, nil
}
