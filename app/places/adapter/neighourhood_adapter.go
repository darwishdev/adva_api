package adapter

import (
	"github.com/darwishdev/bzns_pro_api/common/convertor"
	db "github.com/darwishdev/bzns_pro_api/common/db/gen"
	rmsv1 "github.com/darwishdev/bzns_pro_api/common/pb/bznspro/v1"
)

//list

func (a *PlacesAdapter) neighbourhoodListRowGrpcFromSql(resp *db.NeighbourhoodsListRow) *rmsv1.NeighbourhoodsListRow {
	formattedTime := resp.CreatedAt.Time.Format(a.dateFormat)
	return &rmsv1.NeighbourhoodsListRow{
		NeighbourhoodId:   resp.NeighbourhoodID,
		NeighbourhoodName: resp.NeighbourhoodName,
		DistrictId:        resp.DistrictID,
		DistrictName:      resp.DistrictName,
		CityId:            resp.CityID,
		CityName:          resp.CityName,
		CityCode:          resp.CityCode.String,
		NeighbourhoodCode: resp.NeighbourhoodCode.String,
		CreatedAt:         formattedTime,
		DeletedAt:         resp.DeletedAt.Time.String(),
	}
}

func (a *PlacesAdapter) NeighbourhoodsListGrpcFromSql(resp *[]db.NeighbourhoodsListRow) *rmsv1.NeighbourhoodsListResponse {
	records := make([]*rmsv1.NeighbourhoodsListRow, 0)
	deletedRecords := make([]*rmsv1.NeighbourhoodsListRow, 0)
	for _, v := range *resp {
		record := a.neighbourhoodListRowGrpcFromSql(&v)
		if v.DeletedAt.Valid {
			deletedRecords = append(deletedRecords, record)
		} else {
			records = append(records, record)
		}
	}
	return &rmsv1.NeighbourhoodsListResponse{
		Records:        records,
		DeletedRecords: deletedRecords,
	}
}

func (a *PlacesAdapter) NeighbourhoodsInputListGrpcFromSql(resp *[]db.NeighbourhoodsInputListRow) *rmsv1.NeighbourhoodsInputListResponse {
	// NeighbourhoodsInputListGrpcFromSql
	records := make([]*rmsv1.SelectInputOption, 0)
	for _, v := range *resp {
		record := convertor.ToSelectInput(v.NeighbourhoodID, v.NeighbourhoodName)
		records = append(records, record)
	}
	return &rmsv1.NeighbourhoodsInputListResponse{
		Options: records,
	}
}

func (a *PlacesAdapter) NeighbourhoodCreateSqlFromGrpc(req *rmsv1.NeighbourhoodCreateRequest) *db.NeighbourhoodCreateParams {
	return &db.NeighbourhoodCreateParams{
		DistrictID:        req.DistrictId,
		NeighbourhoodName: req.NeighbourhoodName,
		NeighbourhoodCode: convertor.ToPgType(req.NeighbourhoodCode),
	}
}
func (a *PlacesAdapter) NeighbourhoodEntityGrpcFromSql(resp *db.PlacesSchemaNeighbourhood) *rmsv1.Neighbourhood {
	return &rmsv1.Neighbourhood{
		NeighbourhoodId:   int32(resp.NeighbourhoodID),
		NeighbourhoodName: resp.NeighbourhoodName,
		DistrictId:        resp.DistrictID,
		NeighbourhoodCode: resp.NeighbourhoodCode.String,
		CreatedAt:         resp.CreatedAt.Time.Format(a.dateFormat),
		DeletedAt:         resp.DeletedAt.Time.Format(a.dateFormat),
	}

}
func (a *PlacesAdapter) NeighbourhoodCreateGrpcFromSql(resp *db.PlacesSchemaNeighbourhood) *rmsv1.NeighbourhoodCreateResponse {
	return &rmsv1.NeighbourhoodCreateResponse{
		Neighbourhood: a.NeighbourhoodEntityGrpcFromSql(resp),
	}
}

func (a *PlacesAdapter) NeighbourhoodFindForUpdateGrpcFromSql(resp *db.NeighbourhoodFindForUpdateRow) (*rmsv1.NeighbourhoodUpdateRequest, error) {
	response := &rmsv1.NeighbourhoodUpdateRequest{
		DistrictId:        resp.DistrictID,
		NeighbourhoodId:   resp.NeighbourhoodID,
		NeighbourhoodName: resp.NeighbourhoodName,
		NeighbourhoodCode: resp.NeighbourhoodCode.String,
	}
	return response, nil
}

func (a *PlacesAdapter) NeighbourhoodUpdateSqlFromGrpc(req *rmsv1.NeighbourhoodUpdateRequest) *db.NeighbourhoodUpdateParams {
	return &db.NeighbourhoodUpdateParams{
		DistrictID:        req.DistrictId,
		NeighbourhoodID:   req.NeighbourhoodId,
		NeighbourhoodName: req.NeighbourhoodName,
		NeighbourhoodCode: convertor.ToPgType(req.NeighbourhoodCode),
	}
}
func (a *PlacesAdapter) NeighbourhoodUpdateGrpcFromSql(resp *db.PlacesSchemaNeighbourhood) *rmsv1.NeighbourhoodUpdateResponse {
	return &rmsv1.NeighbourhoodUpdateResponse{
		Neighbourhood: a.NeighbourhoodEntityGrpcFromSql(resp),
	}
}
