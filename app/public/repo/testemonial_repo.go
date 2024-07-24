package repo

import (
	"context"

	db "github.com/darwishdev/bzns_pro_api/common/db/gen"
)

func (repo *PublicRepo) TestemonialCreate(ctx context.Context, req *db.TestemonialCreateParams) (*db.Testemonial, error) {
	category, err := repo.store.TestemonialCreate(context.Background(), *req)

	if err != nil {
		return nil, repo.store.DbErrorParser(err, repo.errorHandler)
	}

	return &category, nil
}

func (repo *PublicRepo) TestemonialUpdate(ctx context.Context, req *db.TestemonialUpdateParams) (*db.Testemonial, error) {
	category, err := repo.store.TestemonialUpdate(context.Background(), *req)

	if err != nil {
		return nil, repo.store.DbErrorParser(err, repo.errorHandler)
	}

	return &category, nil
}

func (repo *PublicRepo) TestemonialsList(ctx context.Context) (*[]db.TestemonialsListRow, error) {
	resp, err := repo.store.TestemonialsList(context.Background())

	if err != nil {
		return nil, repo.store.DbErrorParser(err, repo.errorHandler)
	}

	return &resp, nil
}
func (repo *PublicRepo) TestemonialDeleteRestore(ctx context.Context, req []int32) error {
	err := repo.store.TestemonialDeleteRestore(context.Background(), req)
	if err != nil {
		return repo.store.DbErrorParser(err, repo.errorHandler)
	}
	return nil
}

func (repo *PublicRepo) TestemonialFindForUpdate(ctx context.Context, req *int32) (*db.TestemonialFindForUpdateRow, error) {
	resp, err := repo.store.TestemonialFindForUpdate(context.Background(), *req)
	if err != nil {
		return nil, repo.store.DbErrorParser(err, repo.errorHandler)
	}

	return &resp, nil
}
