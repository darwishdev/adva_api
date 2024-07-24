package gapi

import (
	"context"

	"github.com/bufbuild/connect-go"
	rmsv1 "github.com/darwishdev/bzns_pro_api/common/pb/bznspro/v1"
)

func (api *Api) CategoryCreate(ctx context.Context, req *connect.Request[rmsv1.CategoryCreateRequest]) (*connect.Response[rmsv1.CategoryCreateResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	resp, err := api.publicUsecase.CategoryCreate(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

func (api *Api) CategoryFindForUpdate(ctx context.Context, req *connect.Request[rmsv1.CategoryFindForUpdateRequest]) (*connect.Response[rmsv1.CategoryUpdateRequest], error) {
	if err := ctx.Err(); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	resp, err := api.publicUsecase.CategoryFindForUpdate(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}
func (api *Api) CategoryUpdate(ctx context.Context, req *connect.Request[rmsv1.CategoryUpdateRequest]) (*connect.Response[rmsv1.CategoryUpdateResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	resp, err := api.publicUsecase.CategoryUpdate(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

func (api *Api) CategoriesList(ctx context.Context, req *connect.Request[rmsv1.CategoriesListRequest]) (*connect.Response[rmsv1.CategoriesListResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	resp, err := api.publicUsecase.CategoriesList(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	opts, err := api.GetAccessableActionsForGroup(req.Header(), "categories")
	if err != nil {
		return nil, err
	}
	resp.Options = opts
	return connect.NewResponse(resp), nil
}

func (api *Api) CategoriesInputList(ctx context.Context, req *connect.Request[rmsv1.CategoriesInputListRequest]) (*connect.Response[rmsv1.CategoriesInputListResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	resp, err := api.publicUsecase.CategoriesInputList(ctx, req.Msg)
	if err != nil {
		return nil, err
	}

	return connect.NewResponse(resp), nil
}

func (api *Api) CategoryDeleteRestore(ctx context.Context, req *connect.Request[rmsv1.CategoryDeleteRestoreRequest]) (*connect.Response[rmsv1.CategoryDeleteRestoreResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	resp, err := api.publicUsecase.CategoryDeleteRestore(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}
