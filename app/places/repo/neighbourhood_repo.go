package repo

import (
	"context"

	db "github.com/darwishdev/bzns_pro_api/common/db/gen"
)

func (repo *PlacesRepo) NeighbourhoodsInputList(ctx context.Context, req int32) (*[]db.NeighbourhoodsInputListRow, error) {
	resp, err := repo.store.NeighbourhoodsInputList(context.Background(), req)
	if err != nil {
		return nil, repo.store.DbErrorParser(err, repo.errorHandler)
	}

	return &resp, nil
}

func (repo *PlacesRepo) NeighbourhoodsList(ctx context.Context) (*[]db.NeighbourhoodsListRow, error) {
	resp, err := repo.store.NeighbourhoodsList(context.Background())

	if err != nil {
		return nil, repo.store.DbErrorParser(err, repo.errorHandler)
	}

	return &resp, nil
}

func (repo *PlacesRepo) NeighbourhoodCreate(ctx context.Context, req *db.NeighbourhoodCreateParams) (*db.PlacesSchemaNeighbourhood, error) {
	resp, err := repo.store.NeighbourhoodCreate(context.Background(), *req)

	if err != nil {

		return nil, repo.store.DbErrorParser(err, repo.errorHandler)
	}

	return &resp, nil
}

func (repo *PlacesRepo) NeighbourhoodFindForUpdate(ctx context.Context, req *int32) (*db.NeighbourhoodFindForUpdateRow, error) {
	resp, err := repo.store.NeighbourhoodFindForUpdate(context.Background(), *req)
	if err != nil {
		return nil, repo.store.DbErrorParser(err, repo.errorHandler)
	}

	return &resp, nil
}
func (repo *PlacesRepo) NeighbourhoodUpdate(ctx context.Context, req *db.NeighbourhoodUpdateParams) (*db.PlacesSchemaNeighbourhood, error) {
	resp, err := repo.store.NeighbourhoodUpdate(context.Background(), *req)
	if err != nil {
		return nil, repo.store.DbErrorParser(err, repo.errorHandler)
	}

	return &resp, nil
}

func (repo *PlacesRepo) NeighbourhoodDeleteRestore(ctx context.Context, req []int32) error {
	err := repo.store.NeighbourhoodDeleteRestore(context.Background(), req)
	if err != nil {
		return repo.store.DbErrorParser(err, repo.errorHandler)
	}

	return nil
}
