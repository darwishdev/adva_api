package repo

import (
	"context"

	db "github.com/darwishdev/bzns_pro_api/common/db/gen"
)

func (repo *PlacesRepo) CitiesInputList(ctx context.Context) (*[]db.CitiesInputListRow, error) {
	resp, err := repo.store.CitiesInputList(context.Background())
	if err != nil {
		return nil, repo.store.DbErrorParser(err, repo.errorHandler)
	}

	return &resp, nil
}

func (repo *PlacesRepo) CitiesList(ctx context.Context) (*[]db.CitiesListRow, error) {
	resp, err := repo.store.CitiesList(context.Background())

	if err != nil {
		return nil, repo.store.DbErrorParser(err, repo.errorHandler)
	}

	return &resp, nil
}

func (repo *PlacesRepo) CityCreate(ctx context.Context, req *db.CityCreateParams) (*db.PlacesSchemaCity, error) {
	resp, err := repo.store.CityCreate(context.Background(), *req)

	if err != nil {

		return nil, repo.store.DbErrorParser(err, repo.errorHandler)
	}

	return &resp, nil
}

func (repo *PlacesRepo) CityFindForUpdate(ctx context.Context, req *int32) (*db.CityFindForUpdateRow, error) {
	resp, err := repo.store.CityFindForUpdate(context.Background(), *req)
	if err != nil {
		return nil, repo.store.DbErrorParser(err, repo.errorHandler)
	}

	return &resp, nil
}
func (repo *PlacesRepo) CityUpdate(ctx context.Context, req *db.CityUpdateParams) (*db.PlacesSchemaCity, error) {
	resp, err := repo.store.CityUpdate(context.Background(), *req)
	if err != nil {
		return nil, repo.store.DbErrorParser(err, repo.errorHandler)
	}

	return &resp, nil
}

func (repo *PlacesRepo) CityDeleteRestore(ctx context.Context, req []int32) error {
	err := repo.store.CityDeleteRestore(context.Background(), req)
	if err != nil {
		return repo.store.DbErrorParser(err, repo.errorHandler)
	}

	return nil
}
