package gapi

import (
	"context"

	"github.com/bufbuild/connect-go"
	rmsv1 "github.com/darwishdev/bzns_pro_api/common/pb/bznspro/v1"
	"github.com/rs/zerolog/log"
)

func (api *Api) EventRequestCreate(ctx context.Context, req *connect.Request[rmsv1.EventRequestCreateRequest]) (*connect.Response[rmsv1.EventRequestCreateResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	resp, err := api.eventsUsecase.EventRequestCreate(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

func (api *Api) EventRequestUpdate(ctx context.Context, req *connect.Request[rmsv1.EventRequestUpdateRequest]) (*connect.Response[rmsv1.EventRequestUpdateResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	resp, err := api.eventsUsecase.EventRequestUpdate(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

func (api *Api) EventRequestsList(ctx context.Context, req *connect.Request[rmsv1.EventRequestsListRequest]) (*connect.Response[rmsv1.EventRequestsListResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	authorizedUser, _, err := api.authorizeUser(req.Header())
	if err != nil {
		return nil, err
	}

	resp, err := api.eventsUsecase.EventRequestsList(ctx, req.Msg, authorizedUser)
	if err != nil {
		return nil, err
	}

	options, err := api.GetAccessableActionsForGroup(req.Header(), "event_requests")
	if err != nil {
		return nil, err
	}

	resp.Options = options

	log.Debug().Interface("optios", options).Msg("hooasd")
	return connect.NewResponse(resp), nil
}

func (api *Api) EventRequestDeleteRestore(ctx context.Context, req *connect.Request[rmsv1.EventRequestDeleteRestoreRequest]) (*connect.Response[rmsv1.EventRequestDeleteRestoreResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	resp, err := api.eventsUsecase.EventRequestDeleteRestore(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}
