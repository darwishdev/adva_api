package gapi

import (
	"context"

	"github.com/bufbuild/connect-go"
	rmsv1 "github.com/darwishdev/bzns_pro_api/common/pb/bznspro/v1"
)

func (api *Api) TestemonialCreate(ctx context.Context, req *connect.Request[rmsv1.TestemonialCreateRequest]) (*connect.Response[rmsv1.TestemonialCreateResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	resp, err := api.publicUsecase.TestemonialCreate(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

func (api *Api) TestemonialFindForUpdate(ctx context.Context, req *connect.Request[rmsv1.TestemonialFindForUpdateRequest]) (*connect.Response[rmsv1.TestemonialUpdateRequest], error) {
	if err := ctx.Err(); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	resp, err := api.publicUsecase.TestemonialFindForUpdate(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}
func (api *Api) TestemonialUpdate(ctx context.Context, req *connect.Request[rmsv1.TestemonialUpdateRequest]) (*connect.Response[rmsv1.TestemonialUpdateResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	resp, err := api.publicUsecase.TestemonialUpdate(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

func (api *Api) TestemonialsList(ctx context.Context, req *connect.Request[rmsv1.TestemonialsListRequest]) (*connect.Response[rmsv1.TestemonialsListResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	resp, err := api.publicUsecase.TestemonialsList(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	opts, err := api.GetAccessableActionsForGroup(req.Header(), "testemonials")
	if err != nil {
		return nil, err
	}
	resp.Options = opts
	return connect.NewResponse(resp), nil
}

func (api *Api) TestemonialDeleteRestore(ctx context.Context, req *connect.Request[rmsv1.TestemonialDeleteRestoreRequest]) (*connect.Response[rmsv1.TestemonialDeleteRestoreResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	resp, err := api.publicUsecase.TestemonialDeleteRestore(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}
