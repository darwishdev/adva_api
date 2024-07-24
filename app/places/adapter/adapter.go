package adapter

import (
	db "github.com/darwishdev/bzns_pro_api/common/db/gen"
	rmsv1 "github.com/darwishdev/bzns_pro_api/common/pb/bznspro/v1"
)

type PlacesAdapterInterface interface {
	//global
	PlacesListGrpcFromSql(resp *[][]byte) (*rmsv1.PlacesListResponse, error)
	PlacesInputListGrpcFromSql(resp []*rmsv1.Place) (*rmsv1.PlacesInputListResponse, error)

	//cities
	CitiesListGrpcFromSql(resp *[]db.CitiesListRow) *rmsv1.CitiesListResponse
	CitiesInputListGrpcFromSql(resp *[]db.CitiesInputListRow) *rmsv1.CitiesInputListResponse
	CityCreateSqlFromGrpc(req *rmsv1.CityCreateRequest) *db.CityCreateParams
	CityEntityGrpcFromSql(resp *db.PlacesSchemaCity) *rmsv1.City
	CityCreateGrpcFromSql(resp *db.PlacesSchemaCity) *rmsv1.CityCreateResponse
	CityFindForUpdateGrpcFromSql(resp *db.CityFindForUpdateRow) (*rmsv1.CityUpdateRequest, error)
	CityUpdateSqlFromGrpc(req *rmsv1.CityUpdateRequest) *db.CityUpdateParams
	CityUpdateGrpcFromSql(resp *db.PlacesSchemaCity) *rmsv1.CityUpdateResponse

	// districts
	DistrictsListGrpcFromSql(resp *[]db.DistrictsListRow) *rmsv1.DistrictsListResponse
	DistrictsInputListGrpcFromSql(resp *[]db.DistrictsInputListRow) *rmsv1.DistrictsInputListResponse
	DistrictCreateSqlFromGrpc(req *rmsv1.DistrictCreateRequest) *db.DistrictCreateParams
	DistrictEntityGrpcFromSql(resp *db.PlacesSchemaDistrict) *rmsv1.District
	DistrictCreateGrpcFromSql(resp *db.PlacesSchemaDistrict) *rmsv1.DistrictCreateResponse
	DistrictFindForUpdateGrpcFromSql(resp *db.DistrictFindForUpdateRow) (*rmsv1.DistrictUpdateRequest, error)
	DistrictUpdateSqlFromGrpc(req *rmsv1.DistrictUpdateRequest) *db.DistrictUpdateParams
	DistrictUpdateGrpcFromSql(resp *db.PlacesSchemaDistrict) *rmsv1.DistrictUpdateResponse
	//neighbourhood
	NeighbourhoodsListGrpcFromSql(resp *[]db.NeighbourhoodsListRow) *rmsv1.NeighbourhoodsListResponse
	NeighbourhoodsInputListGrpcFromSql(resp *[]db.NeighbourhoodsInputListRow) *rmsv1.NeighbourhoodsInputListResponse
	NeighbourhoodCreateSqlFromGrpc(req *rmsv1.NeighbourhoodCreateRequest) *db.NeighbourhoodCreateParams
	NeighbourhoodEntityGrpcFromSql(resp *db.PlacesSchemaNeighbourhood) *rmsv1.Neighbourhood
	NeighbourhoodCreateGrpcFromSql(resp *db.PlacesSchemaNeighbourhood) *rmsv1.NeighbourhoodCreateResponse
	NeighbourhoodFindForUpdateGrpcFromSql(resp *db.NeighbourhoodFindForUpdateRow) (*rmsv1.NeighbourhoodUpdateRequest, error)
	NeighbourhoodUpdateSqlFromGrpc(req *rmsv1.NeighbourhoodUpdateRequest) *db.NeighbourhoodUpdateParams
	NeighbourhoodUpdateGrpcFromSql(resp *db.PlacesSchemaNeighbourhood) *rmsv1.NeighbourhoodUpdateResponse
}

type PlacesAdapter struct {
	dateFormat string
}

func NewPlacesAdapter() PlacesAdapterInterface {
	return &PlacesAdapter{
		dateFormat: "2006-01-02 15:04:05",
	}
}
