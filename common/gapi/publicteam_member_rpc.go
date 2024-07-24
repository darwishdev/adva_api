package gapi

import (
	"context"

	"github.com/bufbuild/connect-go"
	rmsv1 "github.com/darwishdev/bzns_pro_api/common/pb/bznspro/v1"
)

func (api *Api) TeamMemberCreate(ctx context.Context, req *connect.Request[rmsv1.TeamMemberCreateRequest]) (*connect.Response[rmsv1.TeamMemberCreateResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	resp, err := api.publicUsecase.TeamMemberCreate(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

func (api *Api) TeamMemberFindForUpdate(ctx context.Context, req *connect.Request[rmsv1.TeamMemberFindForUpdateRequest]) (*connect.Response[rmsv1.TeamMemberUpdateRequest], error) {
	if err := ctx.Err(); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	resp, err := api.publicUsecase.TeamMemberFindForUpdate(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}
func (api *Api) TeamMemberUpdate(ctx context.Context, req *connect.Request[rmsv1.TeamMemberUpdateRequest]) (*connect.Response[rmsv1.TeamMemberUpdateResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	resp, err := api.publicUsecase.TeamMemberUpdate(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

func (api *Api) TeamMembersList(ctx context.Context, req *connect.Request[rmsv1.TeamMembersListRequest]) (*connect.Response[rmsv1.TeamMembersListResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	resp, err := api.publicUsecase.TeamMembersList(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	opts, err := api.GetAccessableActionsForGroup(req.Header(), "team_members")
	if err != nil {
		return nil, err
	}
	resp.Options = opts
	return connect.NewResponse(resp), nil
}

func (api *Api) TeamMemberDeleteRestore(ctx context.Context, req *connect.Request[rmsv1.TeamMemberDeleteRestoreRequest]) (*connect.Response[rmsv1.TeamMemberDeleteRestoreResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	resp, err := api.publicUsecase.TeamMemberDeleteRestore(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}
