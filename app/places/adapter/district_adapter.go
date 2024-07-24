package adapter

import (
	"github.com/darwishdev/bzns_pro_api/common/convertor"
	db "github.com/darwishdev/bzns_pro_api/common/db/gen"
	rmsv1 "github.com/darwishdev/bzns_pro_api/common/pb/bznspro/v1"
)

//list

func (a *PlacesAdapter) districtsListRowGrpcFromSql(resp *db.DistrictsListRow) *rmsv1.DistrictsListRow {
	formattedTime := resp.CreatedAt.Time.Format(a.dateFormat)
	var count int32

	if resp.NeighbourhoodsCount.Valid {
		count = int32(resp.NeighbourhoodsCount.Int64)
	}
	return &rmsv1.DistrictsListRow{
		DistrictId:          resp.DistrictID,
		DistrictName:        resp.DistrictName,
		CityId:              resp.CityID,
		CityName:            resp.CityName,
		CityCode:            resp.CityCode.String,
		DistrictCode:        resp.DistrictCode.String,
		NeighbourhoodsCount: count,
		CreatedAt:           formattedTime,
		DeletedAt:           resp.DeletedAt.Time.String(),
	}
}

func (a *PlacesAdapter) DistrictsListGrpcFromSql(resp *[]db.DistrictsListRow) *rmsv1.DistrictsListResponse {
	records := make([]*rmsv1.DistrictsListRow, 0)
	deletedRecords := make([]*rmsv1.DistrictsListRow, 0)
	for _, v := range *resp {
		record := a.districtsListRowGrpcFromSql(&v)
		if v.DeletedAt.Valid {
			deletedRecords = append(deletedRecords, record)
		} else {
			records = append(records, record)
		}
	}
	return &rmsv1.DistrictsListResponse{
		Records:        records,
		DeletedRecords: deletedRecords,
	}
}

func (a *PlacesAdapter) DistrictsInputListGrpcFromSql(resp *[]db.DistrictsInputListRow) *rmsv1.DistrictsInputListResponse {
	// DistrictsInputListGrpcFromSql
	records := make([]*rmsv1.SelectInputOption, 0)
	for _, v := range *resp {
		record := convertor.ToSelectInput(v.DistrictID, v.DistrictName)
		records = append(records, record)
	}
	return &rmsv1.DistrictsInputListResponse{
		Options: records,
	}
}

func (a *PlacesAdapter) DistrictCreateSqlFromGrpc(req *rmsv1.DistrictCreateRequest) *db.DistrictCreateParams {
	return &db.DistrictCreateParams{
		CityID:       req.CityId,
		DistrictName: req.DistrictName,
		DistrictCode: convertor.ToPgType(req.DistrictCode),
	}
}
func (a *PlacesAdapter) DistrictEntityGrpcFromSql(resp *db.PlacesSchemaDistrict) *rmsv1.District {
	return &rmsv1.District{
		DistrictId:   int32(resp.DistrictID),
		DistrictName: resp.DistrictName,
		CityId:       resp.CityID,
		DistrictCode: resp.DistrictCode.String,
		CreatedAt:    resp.CreatedAt.Time.Format(a.dateFormat),
		DeletedAt:    resp.DeletedAt.Time.Format(a.dateFormat),
	}

}
func (a *PlacesAdapter) DistrictCreateGrpcFromSql(resp *db.PlacesSchemaDistrict) *rmsv1.DistrictCreateResponse {
	return &rmsv1.DistrictCreateResponse{
		District: a.DistrictEntityGrpcFromSql(resp),
	}
}

func (a *PlacesAdapter) DistrictFindForUpdateGrpcFromSql(resp *db.DistrictFindForUpdateRow) (*rmsv1.DistrictUpdateRequest, error) {
	response := &rmsv1.DistrictUpdateRequest{
		CityId:       resp.CityID,
		DistrictId:   resp.DistrictID,
		DistrictName: resp.DistrictName,
		DistrictCode: resp.DistrictCode.String,
	}
	return response, nil
}

func (a *PlacesAdapter) DistrictUpdateSqlFromGrpc(req *rmsv1.DistrictUpdateRequest) *db.DistrictUpdateParams {
	return &db.DistrictUpdateParams{
		CityID:       req.CityId,
		DistrictID:   req.DistrictId,
		DistrictName: req.DistrictName,
		DistrictCode: convertor.ToPgType(req.DistrictCode),
	}
}
func (a *PlacesAdapter) DistrictUpdateGrpcFromSql(resp *db.PlacesSchemaDistrict) *rmsv1.DistrictUpdateResponse {
	return &rmsv1.DistrictUpdateResponse{
		District: a.DistrictEntityGrpcFromSql(resp),
	}
}
