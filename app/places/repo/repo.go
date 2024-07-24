package repo

import (
	"context"

	db "github.com/darwishdev/bzns_pro_api/common/db/gen"
)

type PlacesRepoInterface interface {
	//global
	PlacesList(ctx context.Context) (*[][]byte, error)
	//citeis
	CitiesInputList(ctx context.Context) (*[]db.CitiesInputListRow, error)
	CitiesList(ctx context.Context) (*[]db.CitiesListRow, error)
	CityCreate(ctx context.Context, req *db.CityCreateParams) (*db.PlacesSchemaCity, error)
	CityFindForUpdate(ctx context.Context, req *int32) (*db.CityFindForUpdateRow, error)
	CityUpdate(ctx context.Context, req *db.CityUpdateParams) (*db.PlacesSchemaCity, error)
	CityDeleteRestore(ctx context.Context, req []int32) error
	// ditricts
	DistrictsInputList(ctx context.Context, req int32) (*[]db.DistrictsInputListRow, error)
	DistrictsList(ctx context.Context) (*[]db.DistrictsListRow, error)
	DistrictCreate(ctx context.Context, req *db.DistrictCreateParams) (*db.PlacesSchemaDistrict, error)
	DistrictFindForUpdate(ctx context.Context, req *int32) (*db.DistrictFindForUpdateRow, error)
	DistrictUpdate(ctx context.Context, req *db.DistrictUpdateParams) (*db.PlacesSchemaDistrict, error)
	DistrictDeleteRestore(ctx context.Context, req []int32) error
	// Neighbourhoods
	NeighbourhoodsInputList(ctx context.Context, req int32) (*[]db.NeighbourhoodsInputListRow, error)
	NeighbourhoodsList(ctx context.Context) (*[]db.NeighbourhoodsListRow, error)
	NeighbourhoodCreate(ctx context.Context, req *db.NeighbourhoodCreateParams) (*db.PlacesSchemaNeighbourhood, error)
	NeighbourhoodFindForUpdate(ctx context.Context, req *int32) (*db.NeighbourhoodFindForUpdateRow, error)
	NeighbourhoodUpdate(ctx context.Context, req *db.NeighbourhoodUpdateParams) (*db.PlacesSchemaNeighbourhood, error)
	NeighbourhoodDeleteRestore(ctx context.Context, req []int32) error
}

type PlacesRepo struct {
	store        db.Store
	errorHandler map[string]string
}

func NewPlacesRepo(store db.Store) *PlacesRepo {
	errorHandler := map[string]string{
		"cities_city_name_key":                  "cityName",
		"cities_city_code_key":                  "cityCode",
		"districts_district_name_key":           "districtName",
		"districts_district_code_key":           "districtCode",
		"neighbourhoods_neighbourhood_name_key": "neighbourhoodName",
		"neighbourhoods_neighbourhood_code_key": "neighbourhoodCode",
	}
	return &PlacesRepo{
		store:        store,
		errorHandler: errorHandler,
	}
}
