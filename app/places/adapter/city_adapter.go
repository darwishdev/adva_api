package adapter

import (
	"github.com/darwishdev/bzns_pro_api/common/convertor"
	db "github.com/darwishdev/bzns_pro_api/common/db/gen"
	rmsv1 "github.com/darwishdev/bzns_pro_api/common/pb/bznspro/v1"
)

//list

func (a *PlacesAdapter) citiesListRowGrpcFromSql(resp *db.CitiesListRow) *rmsv1.CitiesListRow {
	formattedTime := resp.CreatedAt.Time.Format(a.dateFormat)
	var count int32

	if resp.DistrictsCount.Valid {
		count = int32(resp.DistrictsCount.Int64)
	}
	return &rmsv1.CitiesListRow{
		CityId:         resp.CityID,
		CityName:       resp.CityName,
		CityCode:       resp.CityCode.String,
		DistrictsCount: count,
		CreatedAt:      formattedTime,
		DeletedAt:      resp.DeletedAt.Time.String(),
	}
}

func (a *PlacesAdapter) CitiesListGrpcFromSql(resp *[]db.CitiesListRow) *rmsv1.CitiesListResponse {
	records := make([]*rmsv1.CitiesListRow, 0)
	deletedRecords := make([]*rmsv1.CitiesListRow, 0)
	for _, v := range *resp {
		record := a.citiesListRowGrpcFromSql(&v)
		if v.DeletedAt.Valid {
			deletedRecords = append(deletedRecords, record)
		} else {
			records = append(records, record)
		}
	}
	return &rmsv1.CitiesListResponse{
		Records:        records,
		DeletedRecords: deletedRecords,
	}
}

func (a *PlacesAdapter) CitiesInputListGrpcFromSql(resp *[]db.CitiesInputListRow) *rmsv1.CitiesInputListResponse {
	// CitiesInputListGrpcFromSql
	records := make([]*rmsv1.SelectInputOption, 0)
	for _, v := range *resp {
		record := convertor.ToSelectInput(v.CityID, v.CityName)
		records = append(records, record)
	}
	return &rmsv1.CitiesInputListResponse{
		Options: records,
	}
}

func (a *PlacesAdapter) CityCreateSqlFromGrpc(req *rmsv1.CityCreateRequest) *db.CityCreateParams {
	return &db.CityCreateParams{
		CityName: req.CityName,
		CityCode: convertor.ToPgType(req.CityCode),
	}
}
func (a *PlacesAdapter) CityEntityGrpcFromSql(resp *db.PlacesSchemaCity) *rmsv1.City {
	return &rmsv1.City{
		CityId:    int32(resp.CityID),
		CityName:  resp.CityName,
		CityCode:  resp.CityCode.String,
		CreatedAt: resp.CreatedAt.Time.Format(a.dateFormat),
		DeletedAt: resp.DeletedAt.Time.Format(a.dateFormat),
	}

}
func (a *PlacesAdapter) CityCreateGrpcFromSql(resp *db.PlacesSchemaCity) *rmsv1.CityCreateResponse {
	return &rmsv1.CityCreateResponse{
		City: a.CityEntityGrpcFromSql(resp),
	}
}

func (a *PlacesAdapter) CityFindForUpdateGrpcFromSql(resp *db.CityFindForUpdateRow) (*rmsv1.CityUpdateRequest, error) {
	response := &rmsv1.CityUpdateRequest{
		CityId:   resp.CityID,
		CityName: resp.CityName,
		CityCode: resp.CityCode.String,
	}
	return response, nil
}

func (a *PlacesAdapter) CityUpdateSqlFromGrpc(req *rmsv1.CityUpdateRequest) *db.CityUpdateParams {
	return &db.CityUpdateParams{
		CityID:   req.CityId,
		CityName: req.CityName,
		CityCode: convertor.ToPgType(req.CityCode),
	}
}
func (a *PlacesAdapter) CityUpdateGrpcFromSql(resp *db.PlacesSchemaCity) *rmsv1.CityUpdateResponse {
	return &rmsv1.CityUpdateResponse{
		City: a.CityEntityGrpcFromSql(resp),
	}
}
