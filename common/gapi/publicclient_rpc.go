package gapi

import (
	"context"

	"github.com/bufbuild/connect-go"
	rmsv1 "github.com/darwishdev/bzns_pro_api/common/pb/bznspro/v1"
)

func (api *Api) LoadInitialData(ctx context.Context, req *connect.Request[rmsv1.ClientInitializeRequest]) (*connect.Response[rmsv1.ClientInitializeResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	resp, err := api.publicUsecase.ClientInitialize(ctx, req.Msg)
	if err != nil {
		return nil, err
	}

	events, err := api.eventsUsecase.EventsList(ctx, &rmsv1.EventsListRequest{})
	if err != nil {
		return nil, err
	}
	resp.Events = events.Records

	blogs, err := api.blogUsecase.BlogsList(ctx, &rmsv1.BlogsListRequest{})
	if err != nil {
		return nil, err
	}
	resp.Blogs = blogs.Records

	return connect.NewResponse(resp), nil
}
