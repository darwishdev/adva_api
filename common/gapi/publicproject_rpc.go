package gapi

import (
	"context"

	"github.com/bufbuild/connect-go"
	rmsv1 "github.com/darwishdev/bzns_pro_api/common/pb/bznspro/v1"
)

func (api *Api) ProjectCreate(ctx context.Context, req *connect.Request[rmsv1.ProjectCreateRequest]) (*connect.Response[rmsv1.ProjectCreateResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	resp, err := api.publicUsecase.ProjectCreate(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

func (api *Api) ProjectFindForUpdate(ctx context.Context, req *connect.Request[rmsv1.ProjectFindForUpdateRequest]) (*connect.Response[rmsv1.ProjectUpdateRequest], error) {
	if err := ctx.Err(); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	resp, err := api.publicUsecase.ProjectFindForUpdate(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}
func (api *Api) ProjectUpdate(ctx context.Context, req *connect.Request[rmsv1.ProjectUpdateRequest]) (*connect.Response[rmsv1.ProjectUpdateResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	resp, err := api.publicUsecase.ProjectUpdate(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

func (api *Api) ProjectsList(ctx context.Context, req *connect.Request[rmsv1.ProjectsListRequest]) (*connect.Response[rmsv1.ProjectsListResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	resp, err := api.publicUsecase.ProjectsList(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	opts, err := api.GetAccessableActionsForGroup(req.Header(), "projects")
	if err != nil {
		return nil, err
	}
	resp.Options = opts
	return connect.NewResponse(resp), nil
}

func (api *Api) ProjectDeleteRestore(ctx context.Context, req *connect.Request[rmsv1.ProjectDeleteRestoreRequest]) (*connect.Response[rmsv1.ProjectDeleteRestoreResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	resp, err := api.publicUsecase.ProjectDeleteRestore(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}
