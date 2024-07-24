package repo

import (
	"context"

	db "github.com/darwishdev/bzns_pro_api/common/db/gen"
)

func (repo *PlacesRepo) DistrictsInputList(ctx context.Context, req int32) (*[]db.DistrictsInputListRow, error) {
	resp, err := repo.store.DistrictsInputList(context.Background(), req)
	if err != nil {
		return nil, repo.store.DbErrorParser(err, repo.errorHandler)
	}

	return &resp, nil
}

func (repo *PlacesRepo) DistrictsList(ctx context.Context) (*[]db.DistrictsListRow, error) {
	resp, err := repo.store.DistrictsList(context.Background())

	if err != nil {
		return nil, repo.store.DbErrorParser(err, repo.errorHandler)
	}

	return &resp, nil
}

func (repo *PlacesRepo) DistrictCreate(ctx context.Context, req *db.DistrictCreateParams) (*db.PlacesSchemaDistrict, error) {
	resp, err := repo.store.DistrictCreate(context.Background(), *req)

	if err != nil {

		return nil, repo.store.DbErrorParser(err, repo.errorHandler)
	}

	return &resp, nil
}

func (repo *PlacesRepo) DistrictFindForUpdate(ctx context.Context, req *int32) (*db.DistrictFindForUpdateRow, error) {
	resp, err := repo.store.DistrictFindForUpdate(context.Background(), *req)
	if err != nil {
		return nil, repo.store.DbErrorParser(err, repo.errorHandler)
	}

	return &resp, nil
}
func (repo *PlacesRepo) DistrictUpdate(ctx context.Context, req *db.DistrictUpdateParams) (*db.PlacesSchemaDistrict, error) {
	resp, err := repo.store.DistrictUpdate(context.Background(), *req)
	if err != nil {
		return nil, repo.store.DbErrorParser(err, repo.errorHandler)
	}

	return &resp, nil
}

func (repo *PlacesRepo) DistrictDeleteRestore(ctx context.Context, req []int32) error {
	err := repo.store.DistrictDeleteRestore(context.Background(), req)
	if err != nil {
		return repo.store.DbErrorParser(err, repo.errorHandler)
	}

	return nil
}
