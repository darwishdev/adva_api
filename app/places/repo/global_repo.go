package repo

import (
	"context"
)

func (repo *PlacesRepo) PlacesList(ctx context.Context) (*[][]byte, error) {
	resp, err := repo.store.PlacesList(context.Background())
	if err != nil {
		return nil, repo.store.DbErrorParser(err, repo.errorHandler)
	}

	return &resp, nil
}
