package usecase

import (
	"context"

	"github.com/bufbuild/connect-go"
	rmsv1 "github.com/darwishdev/bzns_pro_api/common/pb/bznspro/v1"
)

func (u *PlacesUsecase) PlacesList(ctx context.Context, req *rmsv1.PlacesListRequest) (*rmsv1.PlacesListResponse, error) {
	if err := u.validator.Validate(req); err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}
	// params := u.adapter.PlacesListSqlFromGrpc(req)
	records, err := u.repo.PlacesList(ctx)
	if err != nil {
		return nil, err
	}

	resp, err := u.adapter.PlacesListGrpcFromSql(records)
	if err != nil {
		return nil, err
	}
	return resp, nil

}

func (u *PlacesUsecase) PlacesInputList(ctx context.Context, req *rmsv1.PlacesInputListRequest) (*rmsv1.PlacesInputListResponse, error) {
	if err := u.validator.Validate(req); err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}
	// params := u.adapter.PlacesListSqlFromGrpc(req)
	records, err := u.repo.PlacesList(ctx)
	if err != nil {
		return nil, err
	}

	resp, err := u.adapter.PlacesListGrpcFromSql(records)
	if err != nil {
		return nil, err
	}
	inputResp, err := u.adapter.PlacesInputListGrpcFromSql(resp.Places)
	if err != nil {
		return nil, err
	}

	return inputResp, nil

}
