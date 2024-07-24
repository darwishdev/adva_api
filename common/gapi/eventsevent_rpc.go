package gapi

import (
	"context"

	"github.com/bufbuild/connect-go"
	rmsv1 "github.com/darwishdev/bzns_pro_api/common/pb/bznspro/v1"
)

func (api *Api) EventCreate(ctx context.Context, req *connect.Request[rmsv1.EventCreateRequest]) (*connect.Response[rmsv1.EventCreateResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	resp, err := api.eventsUsecase.EventCreate(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

func (api *Api) EventFindForUpdate(ctx context.Context, req *connect.Request[rmsv1.EventFindForUpdateRequest]) (*connect.Response[rmsv1.EventUpdateRequest], error) {
	if err := ctx.Err(); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	resp, err := api.eventsUsecase.EventFindForUpdate(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}
func (api *Api) EventUpdate(ctx context.Context, req *connect.Request[rmsv1.EventUpdateRequest]) (*connect.Response[rmsv1.EventUpdateResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	resp, err := api.eventsUsecase.EventUpdate(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

func (api *Api) EventsList(ctx context.Context, req *connect.Request[rmsv1.EventsListRequest]) (*connect.Response[rmsv1.EventsListResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	resp, err := api.eventsUsecase.EventsList(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	opts, err := api.GetAccessableActionsForGroup(req.Header(), "events")
	if err != nil {
		return nil, err
	}
	resp.Options = opts

	return connect.NewResponse(resp), nil
}

func (api *Api) EventDeleteRestore(ctx context.Context, req *connect.Request[rmsv1.EventDeleteRestoreRequest]) (*connect.Response[rmsv1.EventDeleteRestoreResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	resp, err := api.eventsUsecase.EventDeleteRestore(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}
